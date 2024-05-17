package validate

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func init() {
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = zhTrans.RegisterDefaultTranslations(v, trans)
	}
}

func Error(err error) error {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		for _, e := range errs {
			return errors.New(e.Translate(trans))
		}
	}
	return nil
}
