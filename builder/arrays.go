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

	isDynamic   bool
	dynamicType ir.Type
}

func (a *Array) Type() ir.Type {
	return ir.ARRAY
}

func (a *Array) Value() value.Value {
	return a.Val
}

func (a *Array) Data(b *builder) value.Value {
	dat := b.block.NewGetElementPtr(stringType, a.Val, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	return b.block.NewLoad(types.I8Ptr, dat)
}

func (a *Array) Length(b *builder) value.Value {
	len := b.block.NewGetElementPtr(stringType, a.Val, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	return b.block.NewLoad(types.I64, len)
}

func (a *Array) ElemSize(b *builder) value.Value {
	len := b.block.NewGetElementPtr(stringType, a.Val, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 2))
	return b.block.NewLoad(types.I64, len)
}

// TODO: If dynamic, go through values in for loop (LLVM-side) and free them, free based on the dynamicTypes
func (a *Array) Free(b *builder) {
	b.block.NewCall(b.stdFn("free"), a.Data(b))
}

func (a *Array) Own(b *builder) {
	delete(b.autofree, a.freeind)
}

func (a *Array) Duplicate(b *builder) DynamicValue {
	dat := a.Data(b)
	len := a.Length(b)
	size := a.ElemSize(b)

	datLen := b.block.NewMul(size, len)
	newMem := b.block.NewCall(b.stdFn("malloc"), datLen)
	b.block.NewCall(b.stdFn("memcpy"), newMem, dat, datLen)

	newArr := b.block.NewAlloca(arrayType)

	valPtr := b.block.NewGetElementPtr(arrayType, newArr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	b.block.NewStore(newMem, valPtr)

	sizePtr := b.block.NewGetElementPtr(arrayType, newArr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	b.block.NewStore(len, sizePtr)

	elemSizePtr := b.block.NewGetElementPtr(arrayType, newArr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 2))
	b.block.NewStore(size, elemSizePtr)

	freeind := b.autofreeCnt
	b.autofreeCnt++
	arrV := &Array{Val: newArr, freeind: freeind}

	b.autofree[freeind] = arrV

	return arrV
}

// TODO: Duplicate dynamic values, and own them
func (b *builder) addArray(i *ir.Array) {
	arr := b.block.NewAlloca(arrayType)
	valPtr := b.block.NewGetElementPtr(arrayType, arr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))

	firstVal := b.registers[i.Vals[0]].(Value)
	size := b.sizeof(firstVal.Value())
	mem := b.block.NewCall(b.stdFn("malloc"), b.block.NewMul(size, constant.NewInt(types.I64, int64(len(i.Vals)))))
	for j, val := range i.Vals {
		v := b.registers[val].(Value)
		ptr := b.block.NewGetElementPtr(types.I8, mem, b.block.NewMul(size, constant.NewInt(types.I64, int64(j))))
		b.block.NewCall(b.stdFn("memcpy"), ptr, v.Value(), size)
	}

	b.block.NewStore(mem, valPtr)

	sizePtr := b.block.NewGetElementPtr(arrayType, arr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
	b.block.NewStore(constant.NewInt(types.I64, int64(len(i.Vals))), sizePtr)

	elemSizePtr := b.block.NewGetElementPtr(arrayType, arr, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 2))
	b.block.NewStore(size, elemSizePtr)

	freeind := b.autofreeCnt
	b.autofreeCnt++

	arrV := &Array{Val: arr, freeind: freeind}
	b.autofree[freeind] = arrV
	b.registers[b.index] = arrV
}

func (b *builder) sizeof(val value.Value) value.Value {
	typ := val.Type()
	sizePtr := b.block.NewGetElementPtr(typ, constant.NewNull(types.NewPointer(types.NewArray(2, typ))), constant.NewInt(types.I32, 1))
	return b.block.NewPtrToInt(sizePtr, types.I64)
}
