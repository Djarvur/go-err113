package err113

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

func inspectComparision(pass *analysis.Pass, n ast.Node) bool {
	// check whether the call expression matches time.Now().Sub()
	be, ok := n.(*ast.BinaryExpr)
	if !ok {
		return true
	}

	if be.Op != token.EQL && be.Op != token.NEQ {
		return true
	}

	// check that both left and right hand side are errors
	if !isError(be.X, pass.TypesInfo) && !isError(be.Y, pass.TypesInfo) {
		return true
	}

	pass.Reportf(
		be.Pos(),
		"do not compare errors directly, use errors.Is() instead: %q",
		render(pass.Fset, be),
	)

	return true
}
