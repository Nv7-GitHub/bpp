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

func getIndex(str value.Value, ind value.Value, block *ir.Block) value.Value {
	// Convert to localmem
	ptr := getStrPtr(str, block)

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

	return out
}
