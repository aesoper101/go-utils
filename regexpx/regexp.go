package regexpx

import (
	"github.com/asaskevich/govalidator"
	"regexp"
)

// IsEmail judges whether the string is an email.
func IsEmail(email string) bool {
	return govalidator.IsEmail(email)
}

func IsPhoneNumber(phoneNumber string) bool {
	//regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|195|198|199|(147))\\d{8}$"
	regular := "^(1([3-9]))\\d{9}"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phoneNumber)
}

func IsIDCard(card string) bool {
	regRuler := "(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)"
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(card)
}
