package builder

import (
	"fmt"

	"github.com/Nv7-Github/bpp/ir"
	llir "github.com/llir/llvm/ir"
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
		b.registers[b.index] = newString(b.block, len(str), mem)
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

	Free(*builder)
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

var stringType = types.NewStruct(types.I8Ptr, types.I64)

type String struct {
	Val value.Value
}

func (s *String) Type() ir.Type {
	return ir.STRING
}

func (s *String) Value() value.Value {
	return s.Val
}

func (s *String) Free(b *builder) {
	b.block.NewCall(b.stdFn("free"), s.StringVal(b))
}

func (s *String) StringVal(b *builder) value.Value {
	str := b.block.NewGetElementPtr(stringType, s.Val, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	return b.block.NewLoad(types.I8Ptr, str)
}

func (s *String) Length(b *builder) value.Value {
	len := b.block.NewGetElementPtr(stringType, s.Val, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	return b.block.NewLoad(types.I64, len)
}

func newString(b *llir.Block, length int, mem value.Value) *String {
	str := b.NewAlloca(stringType)
	valPtr := b.NewGetElementPtr(stringType, str, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	b.NewStore(mem, valPtr)

	lenPtr := b.NewGetElementPtr(stringType, str, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	b.NewStore(constant.NewInt(types.I64, int64(length)), lenPtr)

	return &String{Val: str}
}

func (b *builder) addPrint(s *ir.Print) {
	str := b.registers[s.Val].(*String)
	strVal := str.StringVal(b)

	len := str.Length(b)
	cstr := b.block.NewCall(b.stdFn("calloc"), constant.NewInt(types.I64, 0), b.block.NewAdd(len, constant.NewInt(types.I64, 1)))
	b.block.NewCall(b.stdFn("memcpy"), cstr, strVal, len)
	b.block.NewCall(b.stdFn("printf"), b.formatter, cstr)
	b.block.NewCall(b.stdFn("free"), cstr)
}
