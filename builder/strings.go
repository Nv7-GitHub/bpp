package builder

import (
	"github.com/Nv7-Github/bpp/ir"
	llir "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var stringType = types.NewStruct(types.I8Ptr, types.I64)

type String struct {
	Val     value.Value
	freeind int
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

func (s *String) Own(b *builder) {
	delete(b.autofree, s.freeind)
}

func newString(b *llir.Block, length value.Value, mem value.Value, bld *builder) *String {
	str := b.NewAlloca(stringType)
	valPtr := b.NewGetElementPtr(stringType, str, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	b.NewStore(mem, valPtr)

	lenPtr := b.NewGetElementPtr(stringType, str, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	b.NewStore(length, lenPtr)

	v := &String{Val: str}
	v.freeind = bld.autofreeCnt
	bld.autofreeCnt++

	bld.autofree[v.freeind] = v
	return v
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

func (b *builder) addConcat(s *ir.Concat) {
	var length value.Value = constant.NewInt(types.I64, 0)
	for _, s := range s.Vals {
		str := b.registers[s].(*String)
		length = b.block.NewAdd(length, str.Length(b))
	}

	out := b.block.NewCall(b.stdFn("malloc"), length)

	var off value.Value = constant.NewInt(types.I64, 0)
	for i, s := range s.Vals {
		str := b.registers[s].(*String)
		strVal := str.StringVal(b)

		var ptr value.Value
		if i == 0 {
			ptr = out
		} else {
			ptr = b.block.NewGetElementPtr(types.I8, out, off)
		}

		lenV := str.Length(b)
		b.block.NewCall(b.stdFn("memcpy"), ptr, strVal, lenV)
		off = b.block.NewAdd(off, lenV)
	}

	b.registers[b.index] = newString(b.block, length, out, b)
}
