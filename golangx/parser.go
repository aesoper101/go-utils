package golangx

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Parser struct {
	fset *token.FileSet
	file *ast.File
}

// NewParser returns a new Parser. If testSrc is not nil,  it parses source code from testSrc and returns the resulting
// Parser. If testSrc is nil, it parses source code from the contents of filename and returns the resulting Parser.
// at least one of testSrc and filename must not be nil. testSrc may be a string, a byte slice, or an io.Reader.
// If an error is returned, parsing fails and the returned Parser is nil.
func NewParser(filename string, src any, mode parser.Mode) (*Parser, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, mode)
	if err != nil {
		return nil, err
	}
	return &Parser{fset: fset, file: file}, nil
}

// MustParser is a convenience wrapper for NewParser that panics if NewParser returns an error.
func MustParser(p *Parser, err error) *Parser {
	if err != nil {
		panic(err)
	}
	return p
}

// FileSet returns the FileSet used to parse the source code.
func (p *Parser) FileSet() *token.FileSet {
	return p.fset
}

// File returns the parsed File.
func (p *Parser) File() *ast.File {
	return p.file
}
