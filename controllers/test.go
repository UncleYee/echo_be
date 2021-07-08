package controllers

import (
	"echo/common/response"
	"log"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

type Query struct {
	Page string `form:"page"`
	Size string `form:"size" binding:"required"`
}

func TestGet(ctx *gin.Context) {
	param := ctx.Param("param")

	var query Query

	page := ctx.DefaultQuery("page", "1")
	size := ctx.DefaultQuery("size", "20")

	err := ctx.ShouldBindQuery(&query)

	if err != nil {
		response.Fail(ctx, err)
		return
	}

	log.Println(page, size)

	data := map[string]interface{}{
		"param": param,
	}

	response.Success(ctx, data)
}

func TestPost(ctx *gin.Context) {

}
