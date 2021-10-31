package compiler

import (
	"github.com/Nv7-Github/bpp/old/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func CompileCompare(stm *parser.ComparisonStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var lhs value.Value
	var rhs value.Value
	var err error
	lhs, block, err = CompileStmt(stm.Left, block)
	if err != nil {
		return nil, block, err
	}

	rhs, block, err = CompileStmt(stm.Right, block)
	if err != nil {
		return nil, block, err
	}

	kind := lhs.Type()
	_, ok := kind.(*types.PointerType)
	if ok {
		kind = kind.(*types.PointerType).ElemType
	}
	if kind.Equal(types.I8) {
		kind = types.NewArray(0, types.I8)
	}

	var cmp value.Value

	var ipred enum.IPred
	var fpred enum.FPred

	switch stm.Operation {
	case parser.EQUAL:
		ipred = enum.IPredEQ
		fpred = enum.FPredOEQ

	case parser.NOTEQUAL:
		ipred = enum.IPredNE
		fpred = enum.FPredONE

	case parser.GREATER:
		ipred = enum.IPredSGT
		fpred = enum.FPredOGT

	case parser.GREATEREQUAL:
		ipred = enum.IPredSGE
		fpred = enum.FPredOGE

	case parser.LESS:
		ipred = enum.IPredSLT
		fpred = enum.FPredOLT

	case parser.LESSEQUAL:
		ipred = enum.IPredSLE
		fpred = enum.FPredOLT
	}

	switch kind.(type) {
	case *types.ArrayType:
		// The process is: get element ptrs, allocate new variables, store the element ptrs in those variables, load the variables, compare the loaded vars
		ptr1 := getStrPtr(lhs, block)
		ptr2 := getStrPtr(rhs, block)
		val1 := block.NewAlloca(ptr1.Type())
		val2 := block.NewAlloca(ptr2.Type())
		block.NewStore(ptr1, val1)
		block.NewStore(ptr2, val2)
		val1_l := block.NewLoad(val1.ElemType, val1)
		val2_l := block.NewLoad(val2.ElemType, val2)
		eq := block.NewCall(strcmp, val1_l, val2_l)

		// Compare
		cmp = block.NewICmp(ipred, eq, constant.NewInt(types.I32, 0))

	case *types.FloatType:
		cmp = block.NewFCmp(fpred, lhs, rhs)

	case *types.IntType:
		cmp = block.NewICmp(ipred, lhs, rhs)
	}

	val := block.NewZExt(cmp, types.I64)
	return val, block, nil
}

func CompileMath(stm *parser.MathStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	var lhs value.Value
	var rhs value.Value
	var err error
	lhs, block, err = CompileStmt(stm.Left, block)
	if err != nil {
		return nil, block, err
	}

	rhs, block, err = CompileStmt(stm.Right, block)
	if err != nil {
		return nil, block, err
	}

	kind := lhs.Type()
	_, ok := kind.(*types.PointerType)
	if ok {
		kind = kind.(*types.PointerType).ElemType
	}
	if kind.Equal(types.I8) {
		kind = types.NewArray(0, types.I8)
	}

	var res value.Value

	switch kind.(type) {
	case *types.FloatType:
		switch stm.Operation {
		case parser.ADDITION:
			res = block.NewFAdd(lhs, rhs)

		case parser.SUBTRACTION:
			res = block.NewFSub(lhs, rhs)

		case parser.MULTIPLICATION:
			res = block.NewFMul(lhs, rhs)

		case parser.DIVISION:
			res = block.NewFDiv(lhs, rhs)

		case parser.POWER:
			res = block.NewCall(pow, lhs, rhs)
		}

	case *types.IntType:
		switch stm.Operation {
		case parser.ADDITION:
			res = block.NewAdd(lhs, rhs)

		case parser.SUBTRACTION:
			res = block.NewSub(lhs, rhs)

		case parser.MULTIPLICATION:
			res = block.NewMul(lhs, rhs)

		case parser.DIVISION:
			res = block.NewSDiv(lhs, rhs)

		case parser.POWER:
			dv1 := block.NewSIToFP(lhs, types.Double)
			dv2 := block.NewSIToFP(rhs, types.Double)
			dv3 := block.NewCall(pow, dv1, dv2)
			ival := block.NewFAdd(dv3, constant.NewFloat(types.Double, 0.5))
			res = block.NewFPToSI(ival, types.I64)
		}
	}

	return res, block, nil
}
