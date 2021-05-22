package gobpp

import "go/ast"

type Converter func([]ast.Expr) (string, error)

var converters = make(map[string]Converter)

func SetupBasics() {
	converters["print"] = func(args []ast.Expr) (string, error) {
		return ConvertExpr(args[0])
	}
}
