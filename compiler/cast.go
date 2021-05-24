package compiler

import (
	"github.com/Nv7-Github/Bpp/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func CompileTypeCast(stm *parser.TypeCastStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var err error
	var v value.Value
	v, block, err = CompileStmt(stm.Value, block)
	if err != nil {
		return nil, block, err
	}

	kind := v.Type()
	_, ok := kind.(*types.PointerType)
	if ok {
		kind = kind.(*types.PointerType).ElemType
	}

	var res value.Value

	switch kind.(type) {
	case *types.IntType:
		switch {
		case stm.NewType.IsEqual(parser.STRING):
			res = block.NewAlloca(types.NewArray(21, types.I8))
			ptr := block.NewGetElementPtr(res.(*ir.InstAlloca).ElemType, res, constant.NewInt(types.I64, 0), constant.NewInt(types.I64, 0))
			block.NewCall(sprintf, ptr, intCastFmt, v)

		case stm.NewType.IsEqual(parser.FLOAT):
			res = block.NewSIToFP(v, types.Double)

		case stm.NewType.IsEqual(parser.INT):
			res = v
		}

	case *types.FloatType:
		switch {
		case stm.NewType.IsEqual(parser.STRING):
			res = block.NewAlloca(types.NewArray(16, types.I8))
			ptr := block.NewGetElementPtr(res.(*ir.InstAlloca).ElemType, res, constant.NewInt(types.I64, 0), constant.NewInt(types.I64, 0))
			block.NewCall(gcvt, v, constant.NewInt(types.I32, 16), ptr)

		case stm.NewType.IsEqual(parser.INT):
			res = block.NewFPToSI(v, types.I64)

		case stm.NewType.IsEqual(parser.FLOAT):
			res = v
		}

	case *types.ArrayType:
		switch {
		case stm.NewType.IsEqual(parser.STRING):
			res = v

		case stm.NewType.IsEqual(parser.INT):
			nullstr := constant.NewNull(types.NewPointer(types.I8Ptr))
			res = block.NewCall(strtol, getStrPtr(v, block), nullstr, constant.NewInt(types.I32, 10)) // Base 10

		case stm.NewType.IsEqual(parser.FLOAT):
			nullstr := constant.NewNull(types.NewPointer(types.I8Ptr))
			res = block.NewCall(strtod, getStrPtr(v, block), nullstr) // Base 10
		}
	}

	return res, block, nil
}
