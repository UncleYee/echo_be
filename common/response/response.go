package response

// 成功
var Success = map[string]interface{}{
	"code":    0,
	"message": "成功",
}

// 系统错误
var SystemError = map[string]interface{}{
	"code":    -1,
	"message": "系统错误",
}

// 未登录
var NotLogin = map[string]interface{}{
	"code":    1024,
	"message": "未登录",
}

// 成功
func SuccessWithData(data map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    0,
		"message": "成功",
		"data":    data,
	}
}

// 参数校验失败
func ValidateParamFailed(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    1001,
		"message": "参数校验失败，" + msg,
	}
}

// 通用型校验错误
func CodeMsgWithDetail(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    1002,
		"message": msg,
	}
}
