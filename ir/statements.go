package ir

import (
	"fmt"
	"reflect"

	"github.com/Nv7-Github/bpp/parser"
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

	case *parser.ComparisonStmt:
		return i.addComparison(s)

	case *parser.IfStmt:
		return i.addIf(s)

	case *parser.IfBlock:
		return i.addIfB(s)

	case *parser.WhileBlock:
		return i.addWhile(s)

	case *parser.ConcatStmt:
		return i.addConcat(s)

	case *parser.TypeCastStmt:
		return i.addTypeCast(s)

	case *parser.ArrayStmt:
		return i.addArray(s)

	case *parser.IndexStmt:
		return i.addIndex(s)

	case *parser.LengthStmt:
		return i.addLength(s)

	case *parser.ChooseStmt:
		return i.addChoose(s)

	case *parser.RandintStmt:
		return i.addRandint(s)

	case *parser.RandomStmt:
		return i.addRandom(s)

	case *parser.FloorStmt:
		return i.addMathFunction(FLOOR, s.Val)

	case *parser.CeilStmt:
		return i.addMathFunction(CEIL, s.Val)

	case *parser.RoundStmt:
		return i.addMathFunction(ROUND, s.Val)

	default:
		return 0, fmt.Errorf("%v: unknown statement type: %s", s.Pos(), reflect.TypeOf(s).String())
	}
}
