package builder

import (
	"github.com/Nv7-Github/bpp/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type empty struct{}

var stringType = types.NewStruct(types.I8Ptr, types.I64)

type String struct {
	Val value.Value

	freeind int
	owners  map[int]empty

	parent Parent
}

func (s *String) Type() ir.Type {
	return ir.STRING
}

func (s *String) Value() value.Value {
	return s.Val
}

func (s *String) Free(b *builder, owner int) {
	delete(s.owners, owner)
	if len(s.owners) == 0 {
		b.block.NewCall(b.stdFn("free"), s.StringVal(b))
		if s.freeind != -1 {
			delete(b.autofree, s.freeind)
		}
	}
	if s.parent != nil {
		s.parent.Free(owner)
	}
}

func (s *String) StringVal(b *builder) value.Value {
	str := b.block.NewGetElementPtr(stringType, s.Val, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	return b.block.NewLoad(types.I8Ptr, str)
}

func (s *String) Length(b *builder) value.Value {
	length := b.block.NewGetElementPtr(stringType, s.Val, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	return b.block.NewLoad(types.I64, length)
}

func (s *String) Own(b *builder, index int) {
	if s.freeind != -1 {
		delete(b.autofree, s.freeind)
		s.freeind = -1
	}
	s.owners[index] = empty{}
	if s.parent != nil {
		s.parent.Own(index)
	}
}

func (s *String) AddParent(p Parent) {
	s.parent = p
}

func (s *String) Size(_ *builder) value.Value {
	return constant.NewInt(types.I64, 16)
}

func newString(length value.Value, mem value.Value, b *builder) *String {
	str := b.entry.NewAlloca(stringType)
	valPtr := b.block.NewGetElementPtr(stringType, str, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	b.block.NewStore(mem, valPtr)

	lenPtr := b.block.NewGetElementPtr(stringType, str, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	b.block.NewStore(length, lenPtr)

	return newStringFromStruct(str, b, true)
}

func newStringFromStruct(val value.Value, bld *builder, autofree bool) *String {
	v := &String{Val: val, owners: make(map[int]empty), freeind: -1}
	if autofree {
		v.freeind = bld.autofreeCnt
		bld.autofreeCnt++

		bld.autofree[v.freeind] = v
	}
	return v
}

func (s *String) print(b *builder) {
	strVal := s.StringVal(b)
	length := s.Length(b)
	b.block.NewCall(b.stdFn("printf"), b.stdV("fmt"), length, strVal)
}

func (b *builder) addPrint(s *ir.Print) {
	str := b.registers[s.Val].(*String)
	str.print(b)
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

	b.registers[b.index] = newString(length, out, b)
}

func (b *builder) addStringIndex(i *ir.StringIndex) {
	v := b.registers[i.Val].(*String)
	ind := b.registers[i.Index].(*Int)

	char := b.block.NewCall(b.stdFn("malloc"), constant.NewInt(types.I64, 1))

	str := v.StringVal(b)
	ptr := b.block.NewGetElementPtr(types.I8, str, ind.Value())
	b.block.NewCall(b.stdFn("memcpy"), char, ptr, constant.NewInt(types.I64, 1))

	b.registers[b.index] = newString(constant.NewInt(types.I64, 1), char, b)
}

func (b *builder) addStringLength(s *ir.StringLength) {
	b.registers[b.index] = &Int{Val: b.registers[s.Val].(*String).Length(b)}
}
