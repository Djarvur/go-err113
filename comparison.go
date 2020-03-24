package err113

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

func inspectComparision(pass *analysis.Pass, n ast.Node) bool { // nolint: unparam
	// check whether the call expression matches time.Now().Sub()
	be, ok := n.(*ast.BinaryExpr)
	if !ok {
		return true
	}

	// check if it is a comparison operation
	if be.Op != token.EQL && be.Op != token.NEQ {
		return true
	}

	// check that both left and right hand side are not nil
	if pass.TypesInfo.Types[be.X].IsNil() || pass.TypesInfo.Types[be.Y].IsNil() {
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

func isError(v ast.Expr, info *types.Info) bool {
	if intf, ok := info.TypeOf(v).Underlying().(*types.Interface); ok {
		return intf.NumMethods() == 1 && intf.Method(0).FullName() == "(error).Error"
	}

	return false
}
