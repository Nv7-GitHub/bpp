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

	default:
		return nil, b, fmt.Errorf("line %d: unknown type %s", s.Line(), reflect.TypeOf(s))
	}
}
