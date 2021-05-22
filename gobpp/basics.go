package gobpp

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
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
	case token.STRING:
		return arg.Value, nil

	default:
		return "", fmt.Errorf("unknown data type: %s", reflect.TypeOf(arg.Kind))
	}
}
