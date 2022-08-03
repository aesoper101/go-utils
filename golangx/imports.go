package golangx

import (
	"fmt"
	"github.com/aesoper101/go-utils/filex"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// ImportSpec represents a single import spec.
// Taken from  "goa.design/goa/v3/codegen/import.go"
type ImportSpec struct {
	// Name of imported package if needed. e.g "utils" or "" if not needed.
	Name string
	// Go import path of package. e.g. "github.com/aesoper101/go-utils"
	Path string
}

// NewImport creates an import spec.
// Taken from  "goa.design/goa/v3/codegen/import.go"
func NewImport(name, path string) *ImportSpec {
	return &ImportSpec{Name: name, Path: path}
}

// Code returns the Go import statement for the ImportSpec.
// Taken from  "goa.design/goa/v3/codegen/import.go"
func (s *ImportSpec) Code() string {
	if len(s.Name) > 0 {
		return fmt.Sprintf(`%s "%s"`, s.Name, s.Path)
	}
	return fmt.Sprintf(`"%s"`, s.Path)
}

// SimpleImport creates an import with no explicit path component.
// Taken from  "goa.design/goa/v3/codegen/import.go"
func SimpleImport(path string) *ImportSpec {
	return &ImportSpec{Path: path}
}

// CollectImports returns all the imports by Parser
func CollectImports(parser *Parser) []*ImportSpec {
	var imports []*ImportSpec

	file := parser.File()
	if file == nil {
		return imports
	}

	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range genDecl.Specs {
			importSpec, ok := spec.(*ast.ImportSpec)
			if !ok {
				continue
			}
			path, err := strconv.Unquote(importSpec.Path.Value)
			if err != nil {
				continue
			}
			if importSpec.Name != nil {
				imports = append(imports, &ImportSpec{
					Name: importSpec.Name.Name,
					Path: path,
				})
			} else {
				imports = append(imports, &ImportSpec{Path: path})
			}
		}
	}
	return imports
}

// RemoveUnusedImports removes unused imports from the file.
func RemoveUnusedImports(parser *Parser) {
	fset := parser.FileSet()
	f := parser.File()

	if fset == nil || f == nil {
		return
	}

	specList := astutil.Imports(fset, f)
	for _, group := range specList {
		for _, spec := range group {
			path := strings.Trim(spec.Path.Value, `"`)
			if !astutil.UsesImport(f, path) {
				if spec.Name != nil {
					astutil.DeleteNamedImport(fset, f, spec.Name.Name, path)
				} else {
					astutil.DeleteImport(fset, f, path)
				}
			}
		}
	}
}

// SortImports sorts the imports in the file.
func SortImports(parser *Parser) {
	fset := parser.FileSet()
	f := parser.File()

	if fset == nil || f == nil {
		return
	}

	ast.SortImports(fset, f)
}

var invalidPackageNameChar = regexp.MustCompile(`\W`)

// SanitizePackageName replaces invalid package name characters with '_'.
// Taken from  "github.com/99designs/gqlgen/internal/codegen/utils.go"
func SanitizePackageName(pkg string) string {
	return invalidPackageNameChar.ReplaceAllLiteralString(filepath.Base(pkg), "_")
}

// IsInvalidPackageName returns true if the package name is invalid.
func IsInvalidPackageName(pkg string) bool {
	return invalidPackageNameChar.MatchString(filepath.Base(pkg))
}

// IsValidPackageName returns true if the package name is valid.
func IsValidPackageName(pkg string) bool {
	return !IsInvalidPackageName(pkg)
}

var modsRegex = regexp.MustCompile(`^(\*|\[\])*`)

// NormalizeVendor takes a qualified package path and turns it into normal one.
// eg .
// github.com/foo/vendor/github.com/aesoper101/go-utils becomes
// github.com/aesoper101/go-utils
// Taken from  "github.com/99designs/gqlgen/internal/codegen/utils.go"
func NormalizeVendor(pkg string) string {
	modifiers := modsRegex.FindAllString(pkg, 1)[0]
	pkg = strings.TrimPrefix(pkg, modifiers)
	parts := strings.Split(pkg, "/vendor/")
	return modifiers + parts[len(parts)-1]
}

// NameForDir returns the name of the package for the given directory.
// Taken from  "github.com/99designs/gqlgen/internal/codegen/imports.go"
func NameForDir(dir string) string {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return SanitizePackageName(filepath.Base(dir))
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return SanitizePackageName(filepath.Base(dir))
	}
	fset := token.NewFileSet()
	for _, file := range files {
		if !strings.HasSuffix(strings.ToLower(file.Name()), ".go") {
			continue
		}

		filename := filepath.Join(dir, file.Name())
		if src, err := parser.ParseFile(fset, filename, nil, parser.PackageClauseOnly); err == nil {
			return src.Name.Name
		}
	}

	return SanitizePackageName(filepath.Base(dir))
}

// ImportPathForDir takes a path and returns a golang import path for the package
// taken from  "github.com/99designs/gqlgen/internal/codegen/imports.go"
func ImportPathForDir(dir string) (res string) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	dir = filepath.ToSlash(dir)

	modDir, ok := goModuleRoot(dir)
	if ok {
		return modDir
	}

	goPaths := filepath.SplitList(GoPath())
	for i, p := range goPaths {
		goPaths[i] = filepath.ToSlash(filepath.Join(p, "testSrc"))
	}

	for _, gopath := range goPaths {
		if len(gopath) < len(dir) && strings.EqualFold(gopath, dir[0:len(gopath)]) {
			return dir[len(gopath)+1:]
		}
	}

	return ""
}

type goModuleSearchResult struct {
	path       string
	goModPath  string
	moduleName string
}

var goModuleRootCache = map[string]goModuleSearchResult{}

// goModuleRoot returns the root of the current go module if there is a go.mod file in the directory tree
// If not, it returns false
// Taken from  "github.com/99designs/gqlgen/internal/codegen/imports.go"
func goModuleRoot(dir string) (string, bool) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	dir = filepath.ToSlash(dir)

	dirs := []string{dir}
	result := goModuleSearchResult{}

	for {
		modDir := dirs[len(dirs)-1]

		if val, ok := goModuleRootCache[dir]; ok {
			result = val
			break
		}

		if content, err := ioutil.ReadFile(filepath.Join(modDir, "go.mod")); err == nil {
			moduleName := string(modRegex.FindSubmatch(content)[1])
			result = goModuleSearchResult{
				path:       moduleName,
				goModPath:  modDir,
				moduleName: moduleName,
			}
			goModuleRootCache[modDir] = result
			break
		}

		if modDir == "" || modDir == "." || modDir == "/" || strings.HasSuffix(modDir, "\\") {
			// Reached the top of the file tree which means go.mod file is not found
			// Set root folder with a sentinel cache value
			goModuleRootCache[modDir] = result
			break
		}

		dirs = append(dirs, filepath.Dir(modDir))
	}

	// create a cache for each path in a tree traversed, except the top one as it is already cached
	for _, d := range dirs[:len(dirs)-1] {
		if result.moduleName == "" {
			// go.mod is not found in the tree, so the same sentinel value fits all the directories in a tree
			goModuleRootCache[d] = result
		} else {
			if relPath, err := filepath.Rel(result.goModPath, d); err != nil {
				panic(err)
			} else {
				path := result.moduleName
				relPath := filepath.ToSlash(relPath)
				if !strings.HasSuffix(relPath, "/") {
					path += "/"
				}
				path += relPath

				goModuleRootCache[d] = goModuleSearchResult{
					path:       path,
					goModPath:  result.goModPath,
					moduleName: result.moduleName,
				}
			}
		}
	}

	res := goModuleRootCache[dir]
	if res.moduleName == "" {
		return "", false
	}
	return res.path, true
}

func GetGoModuleNameFromDir(dir string) string {
	if res, ok := goModuleRoot(dir); ok {
		return res
	}
	return ""
}

// QualifyPackagePath takes an import and fully qualifies it with a vendor dir, if one is required.
// eg .
// github.com/aesoper101/go-utils becomes
// github.com/foo/vendor/github.com/aesoper101/go-utils
//
// x/tools/packages only supports 'qualified package paths' so this will need to be done prior to calling it
// See https://github.com/golang/go/issues/30289
// Taken from  "github.com/99designs/gqlgen/internal/codegen/util.go"
func QualifyPackagePath(importPath string) string {
	wd := filex.Getwd()

	// in go module mode, the import path doesn't need fixing
	if _, ok := goModuleRoot(wd); ok {
		return importPath
	}

	pkg, err := build.Import(importPath, wd, 0)
	if err != nil {
		return importPath
	}

	return pkg.ImportPath
}

var modRegex = regexp.MustCompile(`module ([^\s]*)`)
