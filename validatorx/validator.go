package validatorx

import (
	"github.com/go-playground/validator/v10"
)

func RegisterValidation(v *validator.Validate) {
	_ = v.RegisterValidation("mobile", isMobile)

}

func isMobile(fl validator.FieldLevel) bool {
	return mobileRegex.MatchString(fl.Field().String())
}
