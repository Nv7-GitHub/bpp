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

	empty := block.NewAlloca(types.NewArray(0, types.I8))
	ptr := block.NewBitCast(empty, types.I8Ptr)
	for _, val := range vals {
		block.NewCall(strcat, ptr, getStrPtr(val, block))
	}

	return empty, block, nil
}

func CompileIndex(stm *parser.IndexStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var v value.Value
	var err error
	v, block, err = CompileStmt(stm.Value, block)
	if err != nil {
		return nil, block, err
	}

	var ind value.Value
	ind, block, err = CompileStmt(stm.Index, block)
	if err != nil {
		return nil, block, err
	}

	// Convert to localmem
	ptr := getStrPtr(v, block)

	// Get index, and convert to char
	charptr := block.NewGetElementPtr(types.I8, ptr, ind)
	char := block.NewLoad(types.I8, charptr)

	// Make output (char[2])
	out := block.NewAlloca(types.NewArray(2, types.I8))
	outptr := block.NewGetElementPtr(out.ElemType, out, constant.NewInt(types.I64, 0), constant.NewInt(types.I64, 0)) // convert char[2] to char*

	// Store []char{char, 0} in it
	block.NewStore(char, outptr)
	second := block.NewGetElementPtr(types.I8, outptr, constant.NewInt(types.I64, 1))
	block.NewStore(constant.NewInt(types.I8, 0), second)

	//block.NewLoad(v.Type(), ptr)
	return out, block, nil
}
