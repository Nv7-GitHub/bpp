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

func ForStmt(stm *ast.ForStmt, fn string) (string, error) {
	out := ""
	if stm.Init != nil {
		o, err := ConvertStmt(stm.Init, fn)
		if err != nil {
			return "", err
		}

		out += o + "\n"
	}

	cond, err := ConvertExpr(stm.Cond)
	if err != nil {
		return "", err
	}
	out += fmt.Sprintf("[WHILE %s]\n", cond)

	body, err := ConvertBlock(stm.Body.List, fn)
	if err != nil {
		return "", err
	}
	out += body

	if stm.Post != nil {
		o, err := ConvertStmt(stm.Post, fn)
		if err != nil {
			return "", err
		}

		out += o + "\n"
	}

	out += "[ENDWHILE]\n"
	return out, nil
}
