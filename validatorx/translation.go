package validatorx

import (
	"github.com/aesoper101/go-utils/validatorx/locales/en"
	"github.com/aesoper101/go-utils/validatorx/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

func RegisterDefaultTranslations(v *validator.Validate, trans ut.Translator) {
	switch trans.Locale() {
	case "zh":
		_ = zh.RegisterDefaultTranslations(v, trans)
		_ = zhTranslations.RegisterDefaultTranslations(v, trans)
		return
	case "en":
		_ = en.RegisterDefaultTranslations(v, trans)
		_ = enTranslations.RegisterDefaultTranslations(v, trans)
		return
	}
}

func TagNameFunc() func(fld reflect.StructField) string {
	return func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	}
}
