package main

import (
	"GoEngine/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	routers.InitRouter(r)
	r.Run(":8001")
}
