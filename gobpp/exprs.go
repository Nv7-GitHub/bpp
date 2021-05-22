package gobpp

import (
	"fmt"
	"go/ast"
	"reflect"
)

func ConvertExpr(expr ast.Expr) (string, error) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		return BasicLit(e)

	default:
		return "", fmt.Errorf("unknown expression type: %s", reflect.TypeOf(e))
	}
}
