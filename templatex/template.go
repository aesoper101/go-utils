package templatex

import (
	"bytes"
	"embed"
	file1 "github.com/aesoper101/go-utils/filex"
	"io"
	"text/template"
	"text/template/parse"
)

type Template struct {
	*template.Template
	FuncMap template.FuncMap
}

// New returns a new template with the functions in the FuncMap.
func New(name string) *Template {
	t := &Template{Template: template.New(name)}
	return t.Funcs(DefaultTemplateFuncMap)
}

// Parse parses text as a template body for t.
func (t *Template) Parse(text string) (*Template, error) {
	if _, err := t.Template.Parse(text); err != nil {
		return nil, err
	}
	return t, nil
}

// ParseFiles creates a new Template and parses the template definitions from the named files.
// if any file contains an error, the first such error is returned and the rest are ignored.
func (t *Template) ParseFiles(filenames ...string) (*Template, error) {
	if _, err := t.Template.ParseFiles(filenames...); err != nil {
		return nil, err
	}
	return t, nil
}

// ParseGlob parses the template definitions in the files identified by pattern
func (t *Template) ParseGlob(pattern string) (*Template, error) {
	if _, err := t.Template.ParseGlob(pattern); err != nil {
		return nil, err
	}
	return t, nil
}

// ParseFS parses the files in the embedded filesystem. returns an error if any of the files fail to parse.
func (t *Template) ParseFS(fs embed.FS, patterns ...string) (*Template, error) {
	if _, err := t.Template.ParseFS(fs, patterns...); err != nil {
		return nil, err
	}
	return t, nil
}

// ParseReader parses text read from reader and adds the results to the template.
func (t *Template) ParseReader(r io.Reader) (*Template, error) {
	buf := bytes.NewBuffer(nil)
	if _, err := buf.ReadFrom(r); err != nil {
		return nil, err
	}
	if _, err := t.Template.Parse(buf.String()); err != nil {
		return nil, err
	}
	return t, nil
}

// Funcs adds the elements of the argument map to the template's functionx map.
func (t *Template) Funcs(funcMap template.FuncMap) *Template {
	if t.FuncMap == nil {
		t.FuncMap = make(template.FuncMap)
	}
	for k, v := range funcMap {
		t.FuncMap[k] = v
	}
	t.Template.Funcs(funcMap)
	return t
}

// AddParseTree adds the argument parsetree to the template definition
func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error) {
	if _, err := t.Template.AddParseTree(name, tree); err != nil {
		return nil, err
	}
	return t, nil
}

// ExecuteTemplateToFile executes a template with the given name and writes the output to a file.
// if Overwrite is true, the file will be overwritten if it exists.
func (t *Template) ExecuteTemplateToFile(name string, file string, overwrite bool, data any) error {
	return file1.CreateFileFromWriterFunc(file, overwrite, func(w io.Writer) error {
		return t.Template.ExecuteTemplate(w, name, data)
	})
}

// ExecuteToFile executes the template to file,
// if overwrite is true, the file will be overwritten if it exists.
func (t *Template) ExecuteToFile(file string, overwrite bool, data any) error {
	return file1.CreateFileFromWriterFunc(file, overwrite, func(w io.Writer) error {
		return t.Template.Execute(w, data)
	})
}
