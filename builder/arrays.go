package builder

import (
	"github.com/Nv7-Github/bpp/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var arrayType = types.NewStruct(types.I8Ptr, types.I64, types.I64) // Data, len, elementsize

type Array struct {
	Val     value.Value
	freeind int

	toFree []Value
	owners map[int]empty
	index  int
}

func (a *Array) Type() ir.Type {
	return ir.ARRAY
}

func (a *Array) Value() value.Value {
	return a.Val
}

func (a *Array) Data(b *builder) value.Value {
	dat := b.block.NewGetElementPtr(arrayType, a.Val, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	return b.block.NewLoad(types.I8Ptr, dat)
}

func (a *Array) Length(b *builder) value.Value {
	length := b.block.NewGetElementPtr(arrayType, a.Val, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	return b.block.NewLoad(types.I64, length)
}

func (a *Array) ElemSize(b *builder) value.Value {
	length := b.block.NewGetElementPtr(arrayType, a.Val, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 2))
	return b.block.NewLoad(types.I64, length)
}

func (a *Array) Free(b *builder, index int) {
	if index != -1 {
		delete(a.owners, index)
	}
	if len(a.owners) == 0 || index == -1 {
		b.block.NewCall(b.stdFn("free"), a.Data(b))
		for _, val := range a.toFree {
			dv, ok := val.(DynamicValue)
			if ok {
				dv.Free(b, a.index)
			}
		}
	}
}

func (a *Array) Own(b *builder, owner int) {
	if a.freeind != -1 {
		delete(b.autofree, a.freeind)
		a.freeind = -1
	}
	a.owners[owner] = empty{}
}

// TODO: Duplicate dynamic values, and own them
func (b *builder) addArray(i *ir.Array) {
	arr := b.block.NewAlloca(arrayType)
	valPtr := b.block.NewGetElementPtr(arrayType, arr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))

	firstVal := b.registers[i.Vals[0]].(Value)
	size := firstVal.Size(b)
	mem := b.block.NewCall(b.stdFn("malloc"), b.block.NewMul(size, constant.NewInt(types.I64, int64(len(i.Vals)))))
	for j, val := range i.Vals {
		v := b.registers[val].(Value)

		var vPtr value.Value
		_, dynamic := v.(DynamicValue)
		if dynamic {
			vPtr = b.block.NewBitCast(v.Value(), types.I8Ptr)
		} else {
			vPtr = b.staticPtr(v)
		}
		ptr := b.block.NewGetElementPtr(types.I8, mem, b.block.NewMul(size, constant.NewInt(types.I64, int64(j))))
		b.block.NewCall(b.stdFn("memcpy"), ptr, vPtr, size)
	}

	b.block.NewStore(mem, valPtr)

	sizePtr := b.block.NewGetElementPtr(arrayType, arr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	b.block.NewStore(constant.NewInt(types.I64, int64(len(i.Vals))), sizePtr)

	elemSizePtr := b.block.NewGetElementPtr(arrayType, arr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 2))
	b.block.NewStore(size, elemSizePtr)

	b.registers[b.index] = newArrayFromStruct(arr, b, true)
}

func newArrayFromStruct(val value.Value, b *builder, autofree bool) *Array {
	arrV := &Array{Val: val, owners: make(map[int]empty), index: b.index}
	if autofree {
		freeind := b.autofreeCnt
		b.autofreeCnt++
		b.autofree[freeind] = arrV
		b.registers[b.index] = arrV
		arrV.freeind = freeind
	}
	return arrV
}

func (a *Array) Size(b *builder) value.Value {
	return constant.NewInt(types.I64, 16)
}

func (b *builder) staticPtr(val Value) value.Value {
	var mem value.Value
	switch val.Type() {
	case ir.INT:
		mem = b.block.NewAlloca(types.I64)

	case ir.FLOAT:
		mem = b.block.NewAlloca(types.Float)
	}
	b.block.NewStore(val.Value(), mem)
	ptr := b.block.NewBitCast(mem, types.I8Ptr)
	return ptr
}

func (b *builder) addArrayLength(s *ir.ArrayLength) {
	b.registers[b.index] = &Int{Val: b.registers[s.Val].(*Array).Length(b)}
}
