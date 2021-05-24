package compiler

import (
	"fmt"
	"reflect"

	"github.com/Nv7-Github/Bpp/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

func CompileStmt(stm parser.Statement, b *ir.Block) (value.Value, *ir.Block, error) {
	switch s := stm.(type) {
	case *parser.Data:
		return CompileData(s, b)

	case *parser.DefineStmt:
		return CompileDefine(s, b)

	case *parser.VarStmt:
		return CompileVar(s, b)

	case *parser.ConcatStmt:
		return CompileConcat(s, b)

	case *parser.ComparisonStmt:
		return CompileCompare(s, b)

	case *parser.IfStmt:
		return CompileIf(s, b)

	case *parser.IfBlock:
		return CompileIfBlock(s, b)

	case *parser.MathStmt:
		return CompileMath(s, b)

	case *parser.WhileBlock:
		return CompileWhileBlock(s, b)

	default:
		return nil, b, fmt.Errorf("line %d: unknown type %s", s.Line(), reflect.TypeOf(s))
	}
}
