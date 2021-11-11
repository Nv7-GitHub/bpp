package builder

import (
	"github.com/Nv7-Github/bpp/old/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Mem struct {
	Val  value.Value
	Type ir.Type
}

type DynamicMem struct {
	Val   DynamicValue
	Type  ir.Type
	Mem   value.Value
	Index int

	Owners map[int]empty
}

func (d *DynamicMem) Own(owner int) {
	d.Owners[owner] = empty{}
}

func (d *DynamicMem) Free(owner int) {
	delete(d.Owners, owner)
}

func (b *builder) addAllocStatic(s *ir.AllocStatic) {
	var mem value.Value
	switch s.Type() {
	case ir.INT:
		mem = b.entry.NewAlloca(types.I64)

	case ir.FLOAT:
		mem = b.entry.NewAlloca(types.Double)
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

func (b *builder) addAllocDynamic(s *ir.AllocDynamic) {
	var mem value.Value
	switch s.Type() {
	case ir.STRING:
		mem = b.entry.NewAlloca(stringType)

	case ir.ARRAY:
		mem = b.entry.NewAlloca(arrayType)
	}
	b.registers[b.index] = &DynamicMem{Val: nil, Type: s.Type(), Mem: mem, Index: b.index, Owners: make(map[int]empty)}
	b.autofreeMem[b.index] = empty{}
}

func (b *builder) addSetMemoryDynamic(s *ir.SetMemoryDynamic) {
	mem := b.registers[s.Mem].(*DynamicMem)
	if mem.Val != nil {
		val := mem.getVal(b)
		val.Free(b, mem.Index)
	}

	val := b.registers[s.Value].(DynamicValue)
	val.Own(b, mem.Index)
	mem.Val = val

	v := val.Value()
	size := val.Size(b)

	ptr1 := b.block.NewBitCast(v, types.I8Ptr)
	ptr2 := b.block.NewBitCast(mem.Mem, types.I8Ptr)
	b.block.NewCall(b.stdFn("memcpy"), ptr2, ptr1, size)

	b.registers[mem.Index] = mem
}

func (m *DynamicMem) getVal(b *builder) DynamicValue {
	var val DynamicValue
	switch m.Type {
	case ir.STRING:
		val = newStringFromStruct(m.Mem, b, false)

	case ir.ARRAY:
		val = newArrayFromStruct(m.Mem, b, m.Val.(*Array).toFree, m.Val.(*Array).ValTyp, false)
	}
	return val
}

func (b *builder) addGetMemoryDynamic(s *ir.GetMemoryDynamic) {
	mem := b.registers[s.Mem].(*DynamicMem)
	b.registers[b.index] = mem.getVal(b)
	b.registers[b.index].(DynamicValue).AddParent(mem)
}
