/**
 * 中间件 - 提取当前用户的登录信息
 *
 * @param addTime 2020-09-18
 * @param author  ChengBo
 */
package middleware

import (
	"GoEngine/entity"
	"GoEngine/libs/helper"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

/**
 * 游客登录模式
 * 用户ID可以为空
 */
func CheckGuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从Cookie中过滤请求的用户标识
		userId := getUserIdFromRequest(c)
		if userId != "" {
			c.Set("userId", userId)
			c.Next()
			return
		}

		//从token中提取用户请求标识

		//接续处理HTTP请求
		c.Next()
	}
}

//从请求中提取用户id
func getUserIdFromRequest(c *gin.Context) string {
	//提取请求的签名信息
	signKeyData, err := c.Cookie("SignKey")

	//提取当前用户id
	if err == nil && signKeyData != "" {
		//解析签名参数
		return decodeRequestSign(signKeyData)
	}

	return ""
}

//解析请求签名
func decodeRequestSign(signData string) string {
	//解密签名数据
	decodeStr := helper.AesDecrypt(signData, helper.GetConfig().EncryptionPrivateKey)

	//定义转移后的变量
	decryptionUserInfo := entity.EncryptionUserInfo{}

	//解析签名
	json.Unmarshal([]byte(decodeStr), &decryptionUserInfo)

	//返回用户userId
	return decryptionUserInfo.UserId
}
