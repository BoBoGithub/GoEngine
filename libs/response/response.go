package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//全局错误码定义变量
var MsgCode map[int]string = map[int]string{
	0:    "成功",
	1000: "系统内部异常错误",
	2000: "未登录",
	3000: "Token参数错误",
	3001: "请求参数解析错误",
}

//正常返回 - JSON
func Return(c *gin.Context, code int, data map[string]interface{}) {
	//提取返回消息内容
	msg, ok := MsgCode[code]
	if !ok {
		msg = ""
	}

	//统一返回数据格式
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"code": code,
		"msg":  msg,
	})
}

//正常返回 - HTML
func ReturnHTML(c *gin.Context, templageName string, data map[string]interface{}) {
	//统一返回数据格式
	c.HTML(http.StatusOK, templageName, data)
}

//异常返回
func ReturnException(c *gin.Context, code int) {
	//提取返回消息内容
	msg, ok := MsgCode[code]
	if !ok {
		msg = ""
	}

	//统一返回数据格式
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//抛出异常信息
func ThrowException(a ...interface{}) {
	//自定义异常信息
	panic(fmt.Sprint(a...))
}
