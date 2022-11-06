package pkg

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func RegistrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {

	return func(ut ut.Translator) (err error) {

		if err = ut.Add(tag, translation, override); err != nil {
			return
		}

		return

	}

}

func TranslateFunc(ut ut.Translator, fe validator.FieldError) string {

	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		return fe.(error).Error()
	}

	return t
}
