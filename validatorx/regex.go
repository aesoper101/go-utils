package validatorx

import "regexp"

var (
	mobileRegexString = `^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`
)

var (
	mobileRegex = regexp.MustCompile(mobileRegexString)
)
