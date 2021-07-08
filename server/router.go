package server

import (
	"echo/common/response"
	"echo/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// 健康检查
	api.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	testRoutes := api.Group("/test")
	{
		testRoutes.GET("/get/:param", controllers.TestGet)
	}

	userRoutes := api.Group("/user")
	{
		userRoutes.GET("/login", func(ctx *gin.Context) {
			response.Success(ctx, nil)
		})
		userRoutes.POST("/register", func(ctx *gin.Context) {
			response.Success(ctx, nil)
		})
	}
}
