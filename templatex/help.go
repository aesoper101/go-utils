package templatex

func Must(t *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return t
}
