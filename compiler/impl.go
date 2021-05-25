package compiler

import (
	"github.com/Nv7-Github/Bpp/parser"
	"github.com/llir/llvm/ir"
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
