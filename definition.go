package err113

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func inspectDefinition(pass *analysis.Pass, tlds map[*ast.CallExpr]struct{}, n ast.Node) bool { //nolint: unparam,gocyclo
	// check whether the call expression matches time.Now().Sub()
	ce, ok := n.(*ast.CallExpr)
	if !ok {
		return true
	}

	if _, ok := tlds[ce]; ok {
		return true
	}

	fn, ok := ce.Fun.(*ast.SelectorExpr)
	if !ok {
		return true
	}

	fx, ok := fn.X.(*ast.Ident)
	if !ok {
		return true
	}

	fp, ok := pass.TypesInfo.ObjectOf(fx).(*types.PkgName)
	if !ok {
		return true
	}

	fxName := fp.Imported().Name()

	if !(fxName == "errors" && fn.Sel.Name == "New") &&
		!(fxName == "fmt" && fn.Sel.Name == "Errorf") {
		return true
	}

	if fxName == "fmt" && fn.Sel.Name == "Errorf" && len(ce.Args) > 0 && strings.Contains(toString(ce.Args[0], pass.TypesInfo), `%w`) {
		return true
	}

	pass.Reportf(
		ce.Pos(),
		"do not define dynamic errors, use wrapped static errors instead: %q",
		render(pass.Fset, ce),
	)

	return true
}

func toString(ex ast.Expr, info *types.Info) string {
	if tv, ok := info.Types[ex]; ok && tv.Value != nil {
		return tv.Value.String()
	}

	return ""
}
