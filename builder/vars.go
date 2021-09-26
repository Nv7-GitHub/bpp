package builder

import (
	"github.com/Nv7-Github/bpp/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Mem struct {
	Val  value.Value
	Type ir.Type
}

func (b *builder) addAllocStatic(s *ir.AllocStatic) {
	var mem value.Value
	switch s.Type() {
	case ir.INT:
		mem = b.block.NewAlloca(types.I64)

	case ir.FLOAT:
		mem = b.block.NewAlloca(types.Double)
	}

	b.registers[b.index] = Mem{Val: mem, Type: s.Type()}
}

func (b *builder) addSetMemory(s *ir.SetMemory) {
	val := b.registers[s.Value].(Value)
	mem := b.registers[s.Mem].(Mem)
	b.block.NewStore(val.Value(), mem.Val)
}

func (b *builder) addGetMemory(s *ir.GetMemory) {
	mem := b.registers[s.Mem].(Mem)
	switch mem.Type {
	case ir.INT:
		b.registers[b.index] = &Int{Val: b.block.NewLoad(types.I64, mem.Val)}

	case ir.FLOAT:
		b.registers[b.index] = &Float{Val: b.block.NewLoad(types.Double, mem.Val)}
	}
}
