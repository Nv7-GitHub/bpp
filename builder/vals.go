package builder

import (
	"fmt"

	"github.com/Nv7-Github/bpp/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (b *builder) addConst(s *ir.Const) error {
	switch s.Type() {
	case ir.INT:
		b.registers[b.index] = &Int{Val: constant.NewInt(types.I64, int64(s.Data.(int)))}
		return nil

	case ir.FLOAT:
		b.registers[b.index] = &Float{Val: constant.NewFloat(types.Double, float64(s.Data.(float64)))}
		return nil

	case ir.STRING:
		str := s.Data.(string)
		name := fmt.Sprintf("str%d", b.tmpCount)
		b.tmpCount++
		globVal := b.mod.NewGlobalDef(name, constant.NewCharArrayFromString(str))
		ptr := b.block.NewGetElementPtr(types.NewArray(uint64(len(str)), types.I8), globVal, constant.NewInt(types.I64, 0), constant.NewInt(types.I64, 0))
		mem := b.block.NewCall(b.stdFn("malloc"), constant.NewInt(types.I64, int64(len(str))))
		b.block.NewCall(b.stdFn("memcpy"), mem, ptr, constant.NewInt(types.I64, int64(len(str))))
		b.registers[b.index] = newString(b.block, constant.NewInt(types.I64, int64(len(str))), mem, b)
		return nil

	default:
		return fmt.Errorf("unknown constant type: %s", s.Type().String())
	}
}

type Value interface {
	Type() ir.Type
	Value() value.Value
}

type DynamicValue interface {
	Value

	Free(b *builder, ownder int)
	Own(b *builder, index int)
}

type Int struct {
	Val value.Value
}

func (i *Int) Type() ir.Type {
	return ir.INT
}

func (i *Int) Value() value.Value {
	return i.Val
}

type Float struct {
	Val value.Value
}

func (f *Float) Type() ir.Type {
	return ir.FLOAT
}

func (f *Float) Value() value.Value {
	return f.Val
}
