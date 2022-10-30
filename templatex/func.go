package templatex

import (
	"github.com/aesoper101/go-utils/str"
	"strconv"
	"strings"
	"text/template"
)

var DefaultTemplateFuncMap = template.FuncMap{
	"toCamel":              str.ToCamel,
	"toLowerCamel":         str.ToLowerCamel,
	"toSnake":              str.ToSnake,
	"toKebab":              str.ToKebab,
	"toScreamingKebab":     str.ToScreamingKebab,
	"toDelimited":          str.ToDelimited,
	"toScreamingDelimited": str.ToScreamingDelimited,
	"toSnakeWithIgnore":    str.ToSnakeWithIgnore,
	"lower":                strings.ToLower,
	"upper":                strings.ToUpper,
	"replace":              strings.ReplaceAll,
	"hasPrefix":            strings.HasPrefix,
	"hasSuffix":            strings.HasSuffix,
	"split":                strings.Split,
	"toString":             str.ToString,
	"join":                 strings.Join,
	"quote":                strconv.Quote,
	"ucFirst":              str.UcFirst,
	"lcFirst":              str.LcFirst,
	"trim":                 strings.Trim,
	"trimLeft":             strings.TrimLeft,
	"trimRight":            strings.TrimRight,
	"trimSpace":            strings.TrimSpace,
	"trimPrefix":           strings.TrimPrefix,
	"trimSuffix":           strings.TrimSuffix,
	"contains":             strings.Contains,
}
