package gobpp

import (
	"fmt"
	"go/ast"
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
	"[]any":   "ARRAY", // use []any for types
}

func ConvertFunc(fn *ast.FuncDecl) (string, error) {
	args := ""
	for _, arg := range fn.Type.Params.List {
		args += fmt.Sprintf(" [PARAM %s %s]", arg.Names[0].Name, typeMap[arg.Type.(*ast.Ident).Name])
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
