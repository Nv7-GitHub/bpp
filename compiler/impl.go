package compiler

import (
	"github.com/Nv7-Github/Bpp/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
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

func CompileRepeat(stm *parser.RepeatStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var in value.Value
	var cnt value.Value
	var err error
	in, block, err = CompileStmt(stm.Val, block)
	if err != nil {
		return nil, block, err
	}

	cnt, block, err = CompileStmt(stm.Count, block)
	if err != nil {
		return nil, block, err
	}

	ln := block.NewCall(strlen, in)
	out := addMalloc(block.NewMul(cnt, ln), block)
	block.NewCall(memset, out, constant.NewInt(types.I32, 1), block.NewMul(cnt, ln))

	// Make loop
	it := block.NewAlloca(types.I64)
	block.NewStore(constant.NewInt(types.I64, 0), it)

	start := block.Parent.NewBlock(getTmp())
	block.NewBr(start)
	block = start
	it_l := block.NewLoad(types.I64, it)
	cmp := block.NewICmp(enum.IPredULT, it_l, cnt)

	body := block.Parent.NewBlock(getTmp())
	end := block.Parent.NewBlock(getTmp())
	block.NewCondBr(cmp, body, end)

	block = body

	// Memcpy to output
	off := block.NewMul(it_l, ln)
	ptr := block.NewGetElementPtr(types.I8, out, off)
	block.NewCall(memcpy, ptr, in, ln)

	// Increment i by 1
	inced := block.NewAdd(it_l, constant.NewInt(types.I64, 1))
	block.NewStore(inced, it)

	block.NewBr(start)
	block = end
	return getStrPtr(out, block), block, nil
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
	out := addMalloc(length, block)

	var off value.Value = constant.NewInt(types.I64, 0)
	for _, val := range vals {
		ptr := block.NewGetElementPtr(types.I8, out, off)
		l := block.NewCall(strlen, val)
		block.NewCall(memcpy, ptr, val, l)

		off = block.NewAdd(off, l)
	}

	return out, block, nil
}
