package gobpp

import (
	"fmt"
	"go/ast"
)

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
	for i, arg := range fn.Type.Params.List {
		args += fmt.Sprintf("[PARAM %s %s]", arg.Names[0].Name, typeMap[arg.Type.(*ast.Ident).Name])
		if i != len(fn.Type.Params.List)-1 {
			args += " "
		}
	}

	out := fmt.Sprintf("[FUNCTION %s %s]", fn.Name, args)
	for _, stmt := range fn.Body.List {
		conved, err := ConvertStmt(stmt)
		if err != nil {
			return "", err
		}

		out += conved + "\n"
	}
	return out, nil
}
