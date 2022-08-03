package golangx

import (
	"bytes"
	"go/ast"
	"go/format"
	"golang.org/x/tools/imports"
	"io"
)

// FormatCodeFromSource formats the given source code.
func FormatCodeFromSource(src []byte) ([]byte, error) {
	return format.Source(src)
}

// FormatCodeFromParser formats the given source code from a parser. and writes the result to the given writer.
// an error is returned if formatting fails.
func FormatCodeFromParser(p *Parser, dst io.Writer, opts ...FormatOptions) error {
	fset := p.FileSet()
	f := p.File()

	// remove unused imports
	RemoveUnusedImports(p)

	// format code
	ast.SortImports(fset, f)

	body := bytes.NewBuffer(nil)
	if err := format.Node(body, fset, f); err != nil {
		return err
	}

	opt := imports.Options{
		Comments:   true,
		FormatOnly: true,
	}

	for _, o := range opts {
		o(&opt)
	}

	bs, err := imports.Process("", body.Bytes(), &opt)
	if err != nil {
		return err
	}

	_, err = io.Copy(dst, bytes.NewReader(bs))

	return err
}

type FormatOptions func(opt *imports.Options)

func WithOptions(opts imports.Options) FormatOptions {
	return func(opt *imports.Options) {
		*opt = opts
	}
}

func WithComments(comments bool) FormatOptions {
	return func(opt *imports.Options) {
		opt.Comments = comments
	}
}

func WithFormatOnly(formatOnly bool) FormatOptions {
	return func(opt *imports.Options) {
		opt.FormatOnly = formatOnly
	}
}

func WithTabWidth(tabWidth int) FormatOptions {
	return func(opt *imports.Options) {
		opt.TabWidth = tabWidth
	}
}

func WithTabIndent(tabIndent bool) FormatOptions {
	return func(opt *imports.Options) {
		opt.TabIndent = tabIndent
	}
}

func WithFragment(fragment bool) FormatOptions {
	return func(opt *imports.Options) {
		opt.Fragment = fragment
	}
}

func WithAllErrors(allErrors bool) FormatOptions {
	return func(opt *imports.Options) {
		opt.AllErrors = allErrors
	}
}
