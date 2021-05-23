package compiler

import (
	"github.com/Nv7-Github/Bpp/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

var m *ir.Module

func Compile(prog *parser.Program) (string, error) {
	m = ir.NewModule()
	generateBuiltins()

	main := m.NewFunc("main", types.I32)
	block := main.NewBlock("")
	block.NewRet(constant.NewInt(types.I32, 0))

	return m.String(), nil
}
