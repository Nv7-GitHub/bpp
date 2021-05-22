package gobpp

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func CallExpr(stm *ast.CallExpr) (string, error) {
	name := stm.Fun.(*ast.Ident).Name
	converter, exists := converters[name]
	if exists {
		return converter(stm.Args)
	}

	args := ""
	for _, arg := range stm.Args {
		a, err := ConvertExpr(arg)
		if err != nil {
			return "", err
		}

		args += " " + a
	}

	return fmt.Sprintf("[%s%s]", strings.ToUpper(name), args), nil
}

func BasicLit(arg *ast.BasicLit) (string, error) {
	switch arg.Kind {
	case token.STRING, token.INT, token.FLOAT:
		return arg.Value, nil

	default:
		return "", fmt.Errorf("unknown data type: %s", arg.Kind.String())
	}
}

var opMap = map[token.Token]string{
	token.ADD: "+",
	token.SUB: "-",
	token.MUL: "*",
	token.QUO: "/",
	token.REM: "%",
}

func BinaryExpr(expr *ast.BinaryExpr) (string, error) {
	x, err := ConvertExpr(expr.X)
	if err != nil {
		return "", err
	}

	y, err := ConvertExpr(expr.Y)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("[MATH %s %s %s]", x, opMap[expr.Op], y), nil
}

func Ident(expr *ast.Ident) string {
	return fmt.Sprintf("[VAR %s]", expr.Name)
}

func IndexExpr(expr *ast.IndexExpr) (string, error) {
	x, err := ConvertExpr(expr.X)
	if err != nil {
		return "", err
	}

	ind, err := ConvertExpr(expr.Index)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("[INDEX %s %s]", x, ind), nil
}

func CompositeLit(expr *ast.CompositeLit) (string, error) {
	pars := ""
	for _, elt := range expr.Elts {
		e, err := ConvertExpr(elt)
		if err != nil {
			return "", err
		}

		pars += " " + e
	}

	return fmt.Sprintf("[ARRAY%s]", pars), nil
}
