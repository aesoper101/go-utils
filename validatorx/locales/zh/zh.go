package zh

import (
	"github.com/aesoper101/go-utils/validatorx/internal/pkg"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func RegisterDefaultTranslations(v *validator.Validate, trans ut.Translator) (err error) {
	translations := []struct {
		tag             string
		translation     string
		override        bool
		customRegisFunc validator.RegisterTranslationsFunc
		customTransFunc validator.TranslationFunc
	}{
		{
			tag:         "mobile",
			translation: "{0}为无效手机号",
			override:    false,
		},
	}

	for _, t := range translations {

		if t.customTransFunc != nil && t.customRegisFunc != nil {

			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, t.customTransFunc)

		} else if t.customTransFunc != nil && t.customRegisFunc == nil {

			err = v.RegisterTranslation(t.tag, trans, pkg.RegistrationFunc(t.tag, t.translation, t.override), t.customTransFunc)

		} else if t.customTransFunc == nil && t.customRegisFunc != nil {

			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, pkg.TranslateFunc)

		} else {
			err = v.RegisterTranslation(t.tag, trans, pkg.RegistrationFunc(t.tag, t.translation, t.override), pkg.TranslateFunc)
		}

		if err != nil {
			return
		}
	}

	return
}
