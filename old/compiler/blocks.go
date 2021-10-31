package compiler

import (
	"github.com/Nv7-Github/bpp/old/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func CompileIf(stm *parser.IfStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var err error
	var v value.Value
	v, block, err = CompileStmt(stm.Condition, block)
	if err != nil {
		return nil, block, err
	}

	cond := block.NewICmp(enum.IPredEQ, v, constant.NewInt(types.I32, 1))

	// TODO: Make IF return value

	iftrue := block.Parent.NewBlock(getTmp())
	iffalse := block.Parent.NewBlock(getTmp())
	end := block.Parent.NewBlock(getTmp())

	block.NewCondBr(cond, iftrue, iffalse)

	var ift value.Value
	ift, block, err = CompileStmt(stm.Body, iftrue)
	if err != nil {
		return nil, block, err
	}
	addLine(iftrue, ift)

	var iff value.Value
	iff, block, err = CompileStmt(stm.Else, iffalse)
	if err != nil {
		return nil, block, err
	}
	addLine(iffalse, iff)

	iftrue.NewBr(end)
	iffalse.NewBr(end)

	return nil, end, nil
}

func CompileIfBlock(stm *parser.IfBlock, block *ir.Block) (value.Value, *ir.Block, error) {
	var err error
	var v value.Value
	v, block, err = CompileStmt(stm.Condition, block)
	if err != nil {
		return nil, block, err
	}

	cond := block.NewICmp(enum.IPredEQ, v, constant.NewInt(types.I32, 1))

	iftrue := block.Parent.NewBlock(getTmp())
	iffalse := block.Parent.NewBlock(getTmp())
	end := block.Parent.NewBlock(getTmp())

	block.NewCondBr(cond, iftrue, iffalse)

	iftrue, err = CompileBlock(stm.Body, iftrue)
	if err != nil {
		return nil, block, err
	}

	if stm.Else != nil {
		iffalse, err = CompileBlock(stm.Else, iffalse)
		if err != nil {
			return nil, block, err
		}
	}

	iftrue.NewBr(end)
	iffalse.NewBr(end)

	return nil, end, nil
}

func CompileWhileBlock(stm *parser.WhileBlock, block *ir.Block) (value.Value, *ir.Block, error) {
	compBlk := block.Parent.NewBlock(getTmp())
	block.NewBr(compBlk)
	block = compBlk

	var err error
	var v value.Value
	v, block, err = CompileStmt(stm.Condition, block)
	if err != nil {
		return nil, block, err
	}

	cond := block.NewICmp(enum.IPredEQ, v, constant.NewInt(types.I32, 1))

	body := block.Parent.NewBlock(getTmp())
	end := block.Parent.NewBlock(getTmp())

	block.NewCondBr(cond, body, end)

	body, err = CompileBlock(stm.Body, body)
	if err != nil {
		return nil, block, err
	}

	body.NewBr(block)

	return nil, end, nil
}
