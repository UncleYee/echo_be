package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseInfo struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Detail  string      `json:"detail,omitempty"`
}

func response(ctx *gin.Context, data interface{}, err error) {
	code := 0
	message := "成功"
	detail := ""

	if err != nil {
		// TODO switch err.(type)

		code = 1001
		message = "请求失败"
	}

	ctx.JSON(http.StatusOK, ResponseInfo{
		Code:    code,
		Message: message,
		Data:    data,
		Detail:  detail,
	})

}

func Success(ctx *gin.Context, data interface{}) {
	response(ctx, data, nil)
}

func Fail(ctx *gin.Context, err error) {
	response(ctx, nil, err)
}
