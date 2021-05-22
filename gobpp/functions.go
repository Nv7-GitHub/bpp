package gobpp

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

type empty struct{}

var hasReturn map[string]empty

var typeMap = map[string]string{
	"string":  "STRING",
	"float64": "FLOAT",
	"float32": "FLOAT",
	"int32":   "INT",
	"int64":   "INT",
	"int":     "INT",
}

func ConvertFunc(fn *ast.FuncDecl) (string, error) {
	args := ""
	for _, arg := range fn.Type.Params.List {
		var kind string

		switch v := arg.Type.(type) {
		case *ast.Ident:
			kind = typeMap[v.Name]

		case *ast.ArrayType:
			kind = "ARRAY"

		default:
			return "", fmt.Errorf("unknown function parameter type: %s", reflect.TypeOf(arg.Type))
		}

		args += fmt.Sprintf(" [PARAM %s %s]", arg.Names[0].Name, kind)
	}

	name := strings.ToUpper(fn.Name.Name)

	out := fmt.Sprintf("[FUNCTION %s%s]\n", name, args)
	for _, stmt := range fn.Body.List {
		conved, err := ConvertStmt(stmt, name)
		if err != nil {
			return "", err
		}

		out += conved + "\n"
	}

	_, exists := hasReturn[name]
	if !exists {
		out += "[RETURN \"\"]\n"
	}
	return out, nil
}

func ReturnStmt(s *ast.ReturnStmt, fn string) (string, error) {
	res, err := ConvertExpr(s.Results[0])
	if err != nil {
		return "", err
	}

	hasReturn[fn] = empty{}

	return fmt.Sprintf("[RETURN %s]", res), nil
}
