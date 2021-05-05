package compiler

import (
	"fmt"
	"reflect"

	"github.com/Nv7-Github/Bpp/parser"
)

func compileStmtRaw(stmt parser.Statement) (string, parser.DataType, error) {
	switch stm := stmt.(type) {
	case *parser.Data:
		return CompileData(stm)

	case *parser.DefineStmt:
		return CompileDefine(stm)

	case *parser.VarStmt:
		return CompileVar(stm)
	}
	return "", parser.NULL, fmt.Errorf("line %d: unknown type %s", stmt.Line(), reflect.TypeOf(stmt))
}
