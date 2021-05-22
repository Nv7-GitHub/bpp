package gobpp

import (
	"fmt"
	"go/ast"
	"go/token"
)

func AssignStmt(stm *ast.AssignStmt) (string, error) {
	r, err := ConvertExpr(stm.Rhs[0])
	if err != nil {
		return "", err
	}

	name := stm.Lhs[0].(*ast.Ident).Name

	switch stm.Tok {
	case token.DEFINE, token.ASSIGN:
		return fmt.Sprintf("[DEFINE %s %s]", name, r), nil

	case token.ADD_ASSIGN:
		return fmt.Sprintf("[DEFINE %s [MATH [VAR %s] + %s]]", name, name, r), nil

	case token.SUB_ASSIGN:
		return fmt.Sprintf("[DEFINE %s [MATH [VAR %s] - %s]]", name, name, r), nil

	case token.MUL_ASSIGN:
		return fmt.Sprintf("[DEFINE %s [MATH [VAR %s] * %s]]", name, name, r), nil

	case token.QUO_ASSIGN:
		return fmt.Sprintf("[DEFINE %s [MATH [VAR %s] / %s]]", name, name, r), nil

	default:
		return "", fmt.Errorf("unknown assignment type: %s", stm.Tok.String())
	}
}

var incDecMap = map[token.Token]string{
	token.INC: "+",
	token.DEC: "-",
}

func IncDecStmt(i *ast.IncDecStmt) (string, error) {
	name := i.X.(*ast.Ident).Name

	return fmt.Sprintf("[DEFINE %s [MATH [VAR %s] %s 1]]", name, name, incDecMap[i.Tok]), nil
}
