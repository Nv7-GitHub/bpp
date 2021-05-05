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

	case *parser.RandintStmt:
		return CompileRandint(stm)

	case *parser.MathStmt:
		return CompileMath(stm)

	case *parser.IndexStmt:
		return CompileIndex(stm)

	case *parser.FloorStmt:
		return CompileFloor(stm)

	case *parser.CeilStmt:
		return CompileCeil(stm)

	case *parser.RoundStmt:
		return CompileRound(stm)

	case *parser.RandomStmt:
		return CompileRandom(stm)

	case *parser.IfStmt:
		return CompileIf(stm)

	case *parser.ComparisonStmt:
		return CompileComparison(stm)

	case *parser.ConcatStmt:
		return CompileConcat(stm)

	case *parser.GotoStmt:
		return CompileGoto(stm)

	case *parser.SectionStmt:
		return CompileSection(stm)

	case *parser.ArgsStmt:
		return CompileArgs(stm)

	case *parser.RepeatStmt:
		return CompileRepeat(stm)

	case *parser.ChooseStmt:
		return CompileChoose(stm)
	}
	return "", parser.NULL, fmt.Errorf("line %d: unknown type %s", stmt.Line(), reflect.TypeOf(stmt))
}
