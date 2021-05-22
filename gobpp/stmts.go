package gobpp

import (
	"fmt"
	"go/ast"
	"reflect"
)

func ConvertStmt(stmt ast.Stmt, fn string) (string, error) {
	switch s := stmt.(type) {
	case *ast.ExprStmt:
		return CallExpr(s.X.(*ast.CallExpr))

	case *ast.AssignStmt:
		return AssignStmt(s)

	case *ast.ReturnStmt:
		return ReturnStmt(s, fn)

	case *ast.IfStmt:
		return IfStmt(s, fn)

	case *ast.IncDecStmt:
		return IncDecStmt(s)

	case *ast.ForStmt:
		return ForStmt(s, fn)

	default:
		return "", fmt.Errorf("unknown statement type %s", reflect.TypeOf(s))
	}
}
