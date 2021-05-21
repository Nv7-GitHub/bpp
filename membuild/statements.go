package membuild

import (
	"fmt"
	"reflect"

	"github.com/Nv7-Github/Bpp/parser"
)

func BuildStmt(p *Program, stmt parser.Statement, instructionum ...int) (Instruction, error) {
	switch s := stmt.(type) {
	case *parser.Data:
		d := ParserDataToData(s)
		return func(p *Program) (Data, error) { return d, nil }, nil

	case *parser.DefineStmt:
		return DefineStmt(p, s)

	case *parser.VarStmt:
		return VarStmt(p, s)

	case *parser.IfStmt:
		return IfStmt(p, s)

	case *parser.ComparisonStmt:
		return CompareStmt(p, s)

	case *parser.MathStmt:
		return MathStmt(p, s)

	case *parser.ConcatStmt:
		return ConcatStmt(p, s)

	case *parser.IndexStmt:
		return IndexStmt(p, s)

	case *parser.ArgsStmt:
		return ArgsStmt(p, s)

	case *parser.RandintStmt:
		return RandintStmt(p, s)

	case *parser.ArrayStmt:
		return ArrayStmt(p, s)

	case *parser.ChooseStmt:
		return ChooseStmt(p, s)

	case *parser.RepeatStmt:
		return RepeatStmt(p, s)

	case *parser.RandomStmt:
		return RandomStmt(p, s)

	case *parser.FloorStmt:
		return FloorStmt(p, s)

	case *parser.CeilStmt:
		return CeilStmt(p, s)

	case *parser.RoundStmt:
		return RoundStmt(p, s)

	case *parser.IfBlock:
		return IfBlock(p, s)

	default:
		return nil, fmt.Errorf("line %d: unknown type %s", stmt.Line(), reflect.TypeOf(stmt).String())
	}
}
