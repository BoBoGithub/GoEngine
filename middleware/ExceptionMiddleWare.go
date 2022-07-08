/**
 * 处理异常中间件
 *
 * 主要处理抛出的panic错误异常
 *
 * @param addTime 2020-09-18
 * @param author  ChengBo
 */
package middleware

import (
	"GoEngine/libs/helper"
	"GoEngine/libs/response"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httputil"
	"runtime"
	"strings"
)

//拦截panic错误异常
func ExceptionRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		//提取body数据
		requestBodyBuf, _ := c.GetRawData()

		//重新设置请求Body信息
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBuf))

		defer func() {
			if r := recover(); r != nil {
				//重新设置请求Body信息
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBuf))

				//提取错误码
				errCode := getErrorCode(fmt.Sprintf("%s", r))

				//记录系统异常信息
				if errCode == 1000 {
					//提取请求信息
					httpRequest, _ := httputil.DumpRequest(c.Request, false)
					headers := strings.Split(string(httpRequest), "\\r\\n")

					//调用堆栈信息
					stack := stack(3)

					//记录系统异常日志
					fmt.Println(fmt.Sprintf("[Recovery] panic recovered:\n\n%s\n%s\n%s", strings.Join(headers, "\r\n"), r, stack))
				}

				//处理错误信息
				dealExceptionMsg(c, errCode)

				//相应请求
				response.ReturnException(c, errCode)
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}

//异常错误码信息
func getErrorCode(errMsg string) int {
	//切分异常信息
	strs := strings.Split(errMsg, " ")

	//检查是否是错误码
	if helper.CheckInt(strs[0]) {
		return helper.StringToInt(strs[0])
	}

	//默认返回：系统内部异常错误
	return 1000
}

//处理异常信息
func dealExceptionMsg(c *gin.Context, errCode int) bool {
	//处理下单报警信息, 调用中台接口失败时 会单独带返回结果发送
	if c.FullPath() == "/game/list" {
		//处理异常报警
		dealExceptionNotice(c, errCode)
	}

	return true
}

//处理异常报警
func dealExceptionNotice(c *gin.Context, errCode int) bool {
	//调用服务处理报警信息

	return true
}

//提取运行堆栈信息 -- 从Gin框架中提取的方法
func stack(skip int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	//var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			//data, err := ioutil.ReadFile(file)
			//if err != nil {
			//	continue
			//}
			//lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		//fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}

	return buf.Bytes()
}
