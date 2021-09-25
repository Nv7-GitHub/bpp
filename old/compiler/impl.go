package compiler

import (
	"github.com/Nv7-Github/bpp/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func CompileFloor(stm *parser.FloorStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var v value.Value
	var err error
	v, block, err = CompileStmt(stm.Val, block)
	if err != nil {
		return nil, block, err
	}
	return block.NewCall(floor, v), block, nil
}

func CompileCeil(stm *parser.CeilStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var v value.Value
	var err error
	v, block, err = CompileStmt(stm.Val, block)
	if err != nil {
		return nil, block, err
	}
	return block.NewCall(ceil, v), block, nil
}

func CompileRound(stm *parser.RoundStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var v value.Value
	var err error
	v, block, err = CompileStmt(stm.Val, block)
	if err != nil {
		return nil, block, err
	}
	return block.NewCall(round, v), block, nil
}

func CompileConcat(stm *parser.ConcatStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	// Build vals
	vals := make([]value.Value, len(stm.Strings))
	var err error
	for i, s := range stm.Strings {
		vals[i], block, err = CompileStmt(s, block)
		if err != nil {
			return nil, block, err
		}
	}

	var length value.Value = constant.NewInt(types.I64, 0)
	for _, val := range vals {
		length = block.NewAdd(length, block.NewCall(strlen, val))
	}
	length = block.NewAdd(length, constant.NewInt(types.I64, int64(len(vals))))
	out := block.NewCall(malloc, length)
	block.NewCall(memset, out, constant.NewInt(types.I32, 0), length)

	var off value.Value = constant.NewInt(types.I64, 0)
	for _, val := range vals {
		ptr := block.NewGetElementPtr(types.I8, out, off)
		l := block.NewCall(strlen, val)
		block.NewCall(memcpy, ptr, val, l)

		// Is val a variable? if not, free it
		_, ok := val.(*ir.InstLoad)
		if !ok {
			block.NewCall(free, val)
		}

		off = block.NewAdd(off, l)
	}

	return out, block, nil
}
