package server

import (
	"echo/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	// 中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")

	// 健康检查
	api.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	userRoutes := api.Group("/user")
	{
		userRoutes.GET("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "login")
		})
		userRoutes.POST("register", func(c *gin.Context) {
			c.JSON(http.StatusOK, response.Success)
		})
	}

	return router
}
