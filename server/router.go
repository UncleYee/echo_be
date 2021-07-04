package server

import (
	"echo/common/response"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	// 中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("api")

	// 健康检查
	api.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})

	userApi := api.Group("/user")
	{
		userApi.GET("/login", func(c *gin.Context) {
			c.String(200, "login")
		})
		userApi.POST("register", func(c *gin.Context) {
			c.JSON(200, response.Success(map[string]interface{}{}))
		})
	}

	return router
}
