/**
 * 项目入口控制器
 *
 * @param addTime 2022-05-06
 * @param author  BoBo
 */
package controller

import (
	"GoEngine/entity"
	"GoEngine/libs/e"
	"GoEngine/libs/response"
	"GoEngine/service"

	"github.com/gin-gonic/gin"
)

//入口页面展示
func ShowIndex(c *gin.Context) {
	//提取首页展示内容数据
	indexShowData := service.GetIndexShowData()

	//返回首页HTML内容
	response.ReturnHTML(c, "index/index", indexShowData)
}

//提取用户游戏列表信息
func GetUserGameList(c *gin.Context) {
	//提取当前登录用户UserId
	userId := c.GetString("userId")

	//查询用户游戏列表信息
	userGameData := service.GetUserGameData(userId)

	//返回用户游戏列表
	response.Return(c, e.SUCCESS, map[string]interface{}{
		"game_list": userGameData,
	})
}

//获取游戏数据
func GetGameInfo(c *gin.Context) {
	//接收token参数
	token := c.DefaultQuery("token", "")
	// token := c.DefaultPostForm("token", "")

	//设置接收请求参数变量
	gameParam := entity.GameRequestParam{}

	//绑定请求参数
	err := c.ShouldBindJSON(&gameParam)
	if err != nil {
		response.ThrowException("3001 GetGameInfo BindJSON error:", err)
	}

	//调用Service查询游戏信息
	// gameInfoData := GameService::GetGameInfo(gameParam)

	//返回游戏信息
	response.Return(c, e.SUCCESS, map[string]interface{}{
		"token":    token,
		"app_id":   gameParam.AppId,
		"app_name": gameParam.AppName,
		"app_desc": gameParam.AppDesc,
	})
}
