package compiler

import (
	"fmt"

	"github.com/Nv7-Github/Bpp/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func addLine(block *ir.Block, val value.Value, kind parser.DataType) {
	switch {
	case kind.IsEqual(parser.STRING):
		block.NewCall(printf, getStrPtr(strFmt, block), getStrPtr(val.(*ir.Global), block))

	case kind.IsEqual(parser.FLOAT):
		block.NewCall(printf, getStrPtr(fltFmt, block), val)

	case kind.IsEqual(parser.INT):
		block.NewCall(printf, getStrPtr(intFmt, block), val)
	}
}

func CompileData(stm *parser.Data, block *ir.Block) (value.Value, *ir.Block, error) {
	t := stm.Type()
	switch {
	case t.IsEqual(parser.STRING):
		return getStr(stm.Data.(string)), block, nil

	case t.IsEqual(parser.INT):
		return constant.NewInt(types.I64, int64(stm.Data.(int))), block, nil

	case t.IsEqual(parser.FLOAT):
		return constant.NewFloat(types.Double, stm.Data.(float64)), block, nil

	default:
		return nil, block, fmt.Errorf("line %d: unknown print type", stm.Line())
	}
}
