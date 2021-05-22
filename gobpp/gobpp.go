package gobpp

import (
	"fmt"
	"go/ast"
	"reflect"
)

func Convert(f *ast.File) (string, error) {
	hasReturn = make(map[string]empty)

	out := ""
	for _, fn := range f.Decls {
		dat, err := ConvertDecl(fn)
		if err != nil {
			return "", err
		}
		out += dat + "\n"
	}

	out += "[MAIN]"

	return out, nil
}

func ConvertDecl(decl ast.Decl) (string, error) {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		return ConvertFunc(d)

	default:
		return "", fmt.Errorf("unknown declaration type: %s", reflect.TypeOf(d).String())
	}
}
