package routers

import (
	"GoEngine/controller"
	"GoEngine/middleware"

	"github.com/gin-gonic/gin"
)

//初始化路由配置
func InitRouter(router *gin.Engine) {
	//加载HTML模板
	router.LoadHTMLGlob("resources/templates/**/*")

	//设置静态文件加载前缀
	router.Static("/static", "resources/statics")

	//设置路由组
	routerGroup := router.Group("/", middleware.ExceptionRecover(), middleware.CheckGuest())

	//入口页面
	routerGroup.GET("/", controller.ShowIndex)

	//提取游戏信息
	routerGroup.POST("/game/list", controller.GetUserGameList)

	//查询游戏信息
	routerGroup.POST("/game/info", controller.GetGameInfo)

}
