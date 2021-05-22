package gobpp

import (
	"fmt"
	"go/ast"
)

func AssignStmt(stm *ast.AssignStmt) (string, error) {
	r, err := ConvertExpr(stm.Rhs[0])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("[DEFINE %s %s]", stm.Lhs[0].(*ast.Ident).Name, r), nil
}
