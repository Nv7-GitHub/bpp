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

var autofree map[string]empty

func CompileData(stm *parser.Data, block *ir.Block) (value.Value, *ir.Block, error) {
	t := stm.Type()
	switch {
	case t.IsEqual(parser.STRING):
		str := getStrPtr(getStr(stm.Data.(string)), block)
		length := int64(len(stm.Data.(string)) + 1)
		dat := block.NewCall(malloc, constant.NewInt(types.I64, length))
		block.NewCall(memcpy, dat, str, constant.NewInt(types.I64, length))
		return dat, block, nil

	case t.IsEqual(parser.INT):
		return constant.NewInt(types.I64, int64(stm.Data.(int))), block, nil

	case t.IsEqual(parser.FLOAT):
		return constant.NewFloat(types.Double, stm.Data.(float64)), block, nil

	case t.IsEqual(parser.NULL):
		return nil, block, nil

	default:
		return nil, block, fmt.Errorf("line %d: unknown data type", stm.Line())
	}
}

func CompileArray(stm *parser.ArrayStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var arr *ir.InstAlloca
	var kind types.Type

	var v value.Value
	var first *ir.InstGetElementPtr
	var err error
	for i, d := range stm.Values {
		v, block, err = CompileStmt(d, block)
		if err != nil {
			return nil, block, err
		}

		if arr == nil {
			kind = v.Type()
			arr = block.NewAlloca(types.NewArray(uint64(len(stm.Values)), kind))
			first = block.NewGetElementPtr(arr.ElemType, arr, constant.NewInt(types.I64, 0), constant.NewInt(types.I64, 0))
		}

		var item *ir.InstGetElementPtr
		if i == 0 {
			item = first
		} else {
			item = block.NewGetElementPtr(kind, first, constant.NewInt(types.I64, int64(i)))
		}
		block.NewStore(v, item)
	}

	return arr, block, nil
}

func CompileDefine(stm *parser.DefineStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var v value.Value
	var err error
	v, block, err = CompileStmt(stm.Value, block)
	if err != nil {
		return nil, block, err
	}

	name := stm.Label.(*parser.Data).Data.(string)

	if v.Type().Equal(types.I8Ptr) {
		leng := block.NewCall(strlen, v)
		length := block.NewAdd(leng, constant.NewInt(types.I64, 1)) // Add 1 for null ptr
		dat := block.NewCall(malloc, length)
		block.NewCall(memcpy, dat, v, length)
		block.NewCall(free, v)
		v = dat
		autofree[name] = empty{}
	}

	_, exists := variables[name]
	var va value.Value
	if !exists {
		va = initBlock.NewAlloca(v.Type())
		if va.Type().Equal(types.NewPointer(types.I8Ptr)) {
			initBlock.NewStore(constant.NewNull(types.I8Ptr), va)
			block.NewCall(free, block.NewLoad(types.I8Ptr, va))
		}
	} else {
		va = variables[name].Val
		if va.Type().Equal(types.NewPointer(types.I8Ptr)) {
			block.NewCall(free, block.NewLoad(types.I8Ptr, va))
		}
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

func CompileArgs(stm *parser.ArgsStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var ind value.Value
	var err error
	ind, block, err = CompileStmt(stm.Index, block)
	if err != nil {
		return nil, block, err
	}
	ind = block.NewAdd(ind, constant.NewInt(types.I64, 1)) // Add 1, because first arg is executable

	argv := block.NewLoad(types.NewPointer(types.I8Ptr), args)
	ptr := block.NewGetElementPtr(types.I8Ptr, argv, ind)
	val := block.NewLoad(types.I8Ptr, ptr)

	return val, block, nil
}
