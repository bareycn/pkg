package r

func OK() map[string]interface{} {
	return map[string]interface{}{
		"code": 0,
		"msg":  "OK",
	}
}

func Data(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code": 0,
		"msg":  "OK",
		"data": data,
	}
}

func OKWithMsg(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code": 0,
		"msg":  msg,
	}
}

func Error(err error) map[string]interface{} {
	return map[string]interface{}{
		"msg": err.Error(),
	}
}

func ErrorWithMsg(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code": 1,
		"msg":  msg,
	}
}

func ErrorWithCode(code int, err error) map[string]interface{} {
	return map[string]interface{}{
		"code":  code,
		"error": err.Error(),
	}
}
