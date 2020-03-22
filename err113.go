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
		ast.Inspect(file, func(n ast.Node) bool { return inspectComparision(pass, n) })
	}

	return nil, nil
}

func isError(v ast.Expr, info *types.Info) bool {
	if intf, ok := info.TypeOf(v).Underlying().(*types.Interface); ok {
		return intf.NumMethods() == 1 && intf.Method(0).FullName() == "(error).Error"
	}

	return false
}

// render returns the pretty-print of the given node
func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}

	return buf.String()
}
