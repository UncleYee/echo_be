package response

// 成功
func Success(data map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    0,
		"message": "成功",
		"data":    data,
	}
}

// 系统错误
func SystemErrpr() map[string]interface{} {
	return map[string]interface{}{
		"code":    -1,
		"message": "系统错误",
	}
}

// 未登录
func NotLogin() map[string]interface{} {
	return map[string]interface{}{
		"code":    1024,
		"message": "未登录",
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
