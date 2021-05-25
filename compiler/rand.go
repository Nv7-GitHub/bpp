package compiler

import (
	"github.com/Nv7-Github/Bpp/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

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

	out := getIndex(v, ind, block)

	return out, block, nil
}

func getIndex(arr value.Value, ind value.Value, block *ir.Block) value.Value {
	if arr.Type().Equal(types.I8Ptr) {
		// Convert to localmem
		ptr := getStrPtr(arr, block)

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

		return getStrPtr(out, block)
	}

	// It's an array!
	elemType := arr.Type().(*types.PointerType).ElemType

	ptr := block.NewGetElementPtr(elemType, arr, constant.NewInt(types.I64, 0), ind)
	return block.NewLoad(elemType.(*types.ArrayType).ElemType, ptr)
}

func CompileChoose(stm *parser.ChooseStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var v value.Value
	var err error
	v, block, err = CompileStmt(stm.Data, block)
	if err != nil {
		return nil, block, err
	}

	var len value.Value
	if v.Type().Equal(types.I8Ptr) {
		len = block.NewCall(strlen, v)
	} else {
		len = constant.NewInt(types.I64, int64(v.Type().(*types.PointerType).ElemType.(*types.ArrayType).Len))
	}

	randval32 := block.NewCall(rand)
	randval := block.NewZExt(randval32, types.I64)
	ind := block.NewURem(randval, len)

	out := getIndex(v, ind, block)

	return out, block, nil
}

func CompileRandint(stm *parser.RandintStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var lower value.Value
	var upper value.Value
	var err error
	lower, block, err = CompileStmt(stm.Lower, block)
	if err != nil {
		return nil, block, err
	}

	upper, block, err = CompileStmt(stm.Upper, block)
	if err != nil {
		return nil, block, err
	}

	randval32 := block.NewCall(rand)
	randval := block.NewZExt(randval32, types.I64)

	top := block.NewAdd(block.NewSub(upper, lower), constant.NewInt(types.I64, 1))
	out := block.NewAdd(block.NewURem(randval, top), lower)

	return out, block, nil
}

// Using this method: https://stackoverflow.com/a/64286825/11388343
func CompileRandom(stm *parser.RandomStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var lower value.Value
	var upper value.Value
	var err error
	lower, block, err = CompileStmt(stm.Lower, block)
	if err != nil {
		return nil, block, err
	}

	upper, block, err = CompileStmt(stm.Upper, block)
	if err != nil {
		return nil, block, err
	}

	in := block.NewSIToFP(block.NewMul(block.NewCall(rand), block.NewCall(rand)), types.Double)
	randflt := block.NewCall(sin, in) // sin(rand() * rand())
	return block.NewFAdd(lower, block.NewFMul(block.NewFSub(upper, lower), block.NewCall(fabs, randflt))), block, nil
}
