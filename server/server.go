package server

import (
	"echo/config"
	"echo/middlewares"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	// 中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 跨域中间件
	r.Use(middlewares.Cors())

	// 路由注册
	registerRoutes(r)

	// 启动服务
	config := config.GetConfig()
	port := config.GetString("server.port")
	r.Run(port)
}
