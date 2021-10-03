package builder

import (
	"fmt"

	"github.com/Nv7-Github/bpp/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Value interface {
	Type() ir.Type
	Value() value.Value
	Size(b *builder) value.Value
}

type Parent interface {
	Own(index int)
	Free(index int)
}

type DynamicValue interface {
	Value

	Free(b *builder, ownder int)
	Own(b *builder, index int)
	AddParent(Parent)
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

func (i *Int) Size(_ *builder) value.Value {
	return constant.NewInt(types.I64, 8)
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

func (f *Float) Size(_ *builder) value.Value {
	return constant.NewInt(types.I64, 8)
}

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

func (b *builder) addCast(s *ir.Cast) {
	v := b.registers[s.Val].(Value)
	switch v.Type() {
	case ir.INT:
		switch s.Type() {
		case ir.STRING:
			res := b.block.NewCall(b.stdFn("malloc"), constant.NewInt(types.I64, 21))
			b.block.NewCall(b.stdFn("sprintf"), res, b.stdV("intfmt"), v.Value())

			// remove null terminator
			length := b.block.NewCall(b.stdFn("strlen"), res)
			newV := b.block.NewCall(b.stdFn("malloc"), length)
			b.block.NewCall(b.stdFn("memcpy"), newV, res, length)

			str := newString(b.block, length, newV, b)
			b.registers[b.index] = str
			b.block.NewCall(b.stdFn("free"), res) // Free sprintf output

		case ir.FLOAT:
			b.registers[b.index] = &Float{Val: b.block.NewSIToFP(v.Value(), types.Double)}

		case ir.INT:
			b.registers[b.index] = v
		}

	case ir.FLOAT:
		switch s.Type() {
		case ir.STRING:
			res := b.block.NewCall(b.stdFn("malloc"), constant.NewInt(types.I64, 17))
			b.block.NewCall(b.stdFn("gcvt"), v.Value(), constant.NewInt(types.I32, 17), res)

			length := b.block.NewCall(b.stdFn("strlen"), res)
			newV := b.block.NewCall(b.stdFn("malloc"), length)
			b.block.NewCall(b.stdFn("memcpy"), newV, res, length)

			str := newString(b.block, length, newV, b)
			b.registers[b.index] = str
			b.block.NewCall(b.stdFn("free"), res) // Free gcvt output

		case ir.INT:
			b.registers[b.index] = &Int{Val: b.block.NewFPToSI(v.Value(), types.I64)}

		case ir.FLOAT:
			b.registers[b.index] = v
		}

	case ir.STRING:
		switch s.Type() {
		case ir.STRING:
			b.registers[b.index] = v

		case ir.INT:
			str := v.(*String)
			nullstr := constant.NewNull(types.NewPointer(types.I8Ptr))
			out := b.block.NewCall(b.stdFn("strtol"), str.StringVal(b), nullstr, constant.NewInt(types.I32, 10))
			b.registers[b.index] = &Int{Val: out}

		case ir.FLOAT:
			str := v.(*String)
			nullstr := constant.NewNull(types.NewPointer(types.I8Ptr))
			out := b.block.NewCall(b.stdFn("strtod"), str.StringVal(b), nullstr)
			b.registers[b.index] = &Float{Val: out}
		}
	}
}
