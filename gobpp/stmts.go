package gobpp

import (
	"fmt"
	"go/ast"
	"reflect"
)

func ConvertStmt(stmt ast.Stmt) (string, error) {
	switch s := stmt.(type) {
	case *ast.ExprStmt:
		return CallExpr(s.X.(*ast.CallExpr))

	case *ast.AssignStmt:
		return AssignStmt(s)

	default:
		return "", fmt.Errorf("unknown statement type %s", reflect.TypeOf(s))
	}
}
