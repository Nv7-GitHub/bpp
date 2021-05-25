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

type empty struct{}

var autofree map[value.Value]empty

func CompileData(stm *parser.Data, block *ir.Block) (value.Value, *ir.Block, error) {
	t := stm.Type()
	switch {
	case t.IsEqual(parser.STRING):
		str := getStrPtr(getStr(stm.Data.(string)), block)
		dat := block.NewCall(malloc, constant.NewInt(types.I64, int64(len(stm.Data.(string))+1)))
		block.NewCall(memcpy, dat, str, constant.NewInt(types.I64, int64(len(stm.Data.(string)))))
		last := block.NewGetElementPtr(types.I8, dat, constant.NewInt(types.I64, int64(len(stm.Data.(string)))))
		block.NewStore(constant.NewInt(types.I8, 0), last)
		autofree[dat] = empty{}
		return dat, block, nil

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

	name := stm.Label.(*parser.Data).Data.(string)

	_, exists := variables[name]
	var va value.Value
	if !exists {
		va = block.NewAlloca(v.Type())
	} else {
		va = variables[name].Val
	}

	block.NewStore(v, va)

	if !exists {
		variables[name] = Variable{
			Val:  va,
			Type: va.Type(),
		}
	}

	return nil, block, nil
}

func CompileVar(stm *parser.VarStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	va := variables[stm.Label.(*parser.Data).Data.(string)]
	loaded := block.NewLoad(va.Type.(*types.PointerType).ElemType, va.Val)
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

	out := block.NewCall(malloc, constant.NewInt(types.I64, 0))

	for _, val := range vals {
		size1 := block.NewCall(strlen, out)
		size2 := block.NewCall(strlen, val)

		combined := block.NewAdd(size1, size2)

		res := block.NewCall(malloc, combined)
		block.NewCall(memcpy, res, out, size1)

		ress := block.NewGetElementPtr(types.I8, res, size1)
		block.NewCall(memcpy, ress, val, size2)

		last := block.NewGetElementPtr(types.I8, ress, combined)
		block.NewStore(constant.NewInt(types.I8, 0), last)

		block.NewCall(free, out)
		out = res
	}

	autofree[out] = empty{}

	return out, block, nil
}
