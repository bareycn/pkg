package r

import (
	"github.com/bareycn/pkg/validate"
)

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
	if errs := validate.Error(err); errs != nil {
		return map[string]interface{}{
			"error": errs.Error(),
		}
	}
	return map[string]interface{}{
		"error": err.Error(),
	}
}

func ErrorWithMsg(msg string) map[string]interface{} {
	return map[string]interface{}{
		"error": msg,
	}
}

func ErrorWithCode(code int, err string) map[string]interface{} {
	return map[string]interface{}{
		"code":  code,
		"error": err,
	}
}
