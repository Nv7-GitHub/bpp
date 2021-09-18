package compiler

import (
	"github.com/Nv7-Github/bpp/parser"
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
	var res value.Value

	switch {
	case kind.Equal(types.I64):
		switch {
		case stm.NewType.IsEqual(parser.STRING):
			res = block.NewCall(malloc, constant.NewInt(types.I64, 21))
			block.NewCall(sprintf, res, getStrPtr(intFmt, block), v)

		case stm.NewType.IsEqual(parser.FLOAT):
			res = block.NewSIToFP(v, types.Double)

		case stm.NewType.IsEqual(parser.INT):
			res = v
		}

	case kind.Equal(types.Double):
		switch {
		case stm.NewType.IsEqual(parser.STRING):
			res = block.NewCall(malloc, constant.NewInt(types.I64, 16))
			block.NewCall(gcvt, v, constant.NewInt(types.I32, 16), res)

		case stm.NewType.IsEqual(parser.INT):
			res = block.NewFPToSI(v, types.I64)

		case stm.NewType.IsEqual(parser.FLOAT):
			res = v
		}

	case kind.Equal(types.I8Ptr):
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
