package gobpp

import (
	"fmt"
	"go/ast"
)

func IfStmt(i *ast.IfStmt, fn string) (string, error) {
	cond, err := ConvertExpr(i.Cond)
	if err != nil {
		return "", err
	}
	out := fmt.Sprintf("[IFB %s]\n", cond)

	body, err := ConvertBlock(i.Body.List, fn)
	if err != nil {
		return "", err
	}
	out += body

	if i.Else != nil {
		out += "[ELSE]\n"
		el, err := ConvertBlock(i.Else.(*ast.BlockStmt).List, fn)
		if err != nil {
			return "", err
		}

		out += el
	}

	out += "[ENDIF]"
	return out, nil
}
