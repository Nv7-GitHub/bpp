package compiler

import (
	"fmt"

	"github.com/Nv7-Github/Bpp/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Variable struct {
	Val  value.Value
	Type types.Type
}

var variables map[string]Variable

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

func CompileDefine(stm *parser.DefineStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var v value.Value
	var err error
	v, block, err = CompileStmt(stm.Value, block)
	if err != nil {
		return nil, block, err
	}

	va := block.NewAlloca(v.Type())
	block.NewStore(v, va)
	variables[stm.Label.(*parser.Data).Data.(string)] = Variable{
		Val:  va,
		Type: va.Typ.ElemType,
	}

	return nil, block, nil
}

func CompileVar(stm *parser.VarStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	va := variables[stm.Label.(*parser.Data).Data.(string)]
	loaded := block.NewLoad(va.Type, va.Val)
	return loaded, block, nil
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

	empty := block.NewAlloca(types.NewArray(0, types.I8))
	ptr := block.NewBitCast(empty, types.I8Ptr)
	for _, val := range vals {
		block.NewCall(strcat, ptr, getStrPtr(val, block))
	}

	return empty, block, nil
}
