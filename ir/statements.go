package ir

import (
	"fmt"
	"reflect"

	"github.com/Nv7-Github/Bpp/parser"
)

func (i *IR) AddStmt(stmt parser.Statement) (int, error) {
	switch s := stmt.(type) {
	case *parser.Data:
		cst, err := createConst(s)
		if err != nil {
			return 0, err
		}
		return i.AddInstruction(cst), nil

	case *parser.DefineStmt:
		return i.addDefine(s)

	case *parser.VarStmt:
		return i.addVar(s)

	case *parser.MathStmt:
		return i.addMath(s)

	default:
		return 0, fmt.Errorf("%v: unknown statement type: %s", s.Pos(), reflect.TypeOf(s).String())
	}
}
