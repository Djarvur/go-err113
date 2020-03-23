// Package err113 implements the checks
package err113

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

// NewAnalyzer creates a new analysis.Analyzer instance tuned to run err113 checks
func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "err113",
		Doc:  "checks the error handling rules according to the Go 1.13 new error type",
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		tlds := enumerateFileDecls(file, pass.TypesInfo)
		ast.Inspect(
			file,
			func(n ast.Node) bool {
				return inspectComparision(pass, n) &&
					inspectDefinition(pass, tlds, n)
			},
		)
	}

	return nil, nil
}

// render returns the pretty-print of the given node
func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}

	return buf.String()
}

func enumerateFileDecls(f *ast.File, info *types.Info) map[*ast.CallExpr]struct{} {
	res := make(map[*ast.CallExpr]struct{})

	for _, d := range f.Decls {
		if td, ok := d.(*ast.GenDecl); ok {
			if td.Tok == token.VAR {
				for _, s := range td.Specs {
					if vs, ok := s.(*ast.ValueSpec); ok {
						for _, v := range vs.Values {
							if ce, ok := v.(*ast.CallExpr); ok {
								res[ce] = struct{}{}
							}
						}
					}
				}
			}
		}
	}

	return res
}
