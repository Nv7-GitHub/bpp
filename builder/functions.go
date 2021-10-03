package builder

import (
	"errors"
	"fmt"

	"github.com/Nv7-Github/bpp/ir"
	llir "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var fnTypeMap = map[ir.Type]types.Type{
	ir.INT:    types.I64,
	ir.FLOAT:  types.Double,
	ir.STRING: stringType,
}

func (b *builder) addFn(index int) error {
	fn := b.ir.Functions[index]
	pars := make([]*llir.Param, len(fn.ParTypes))
	for i, parType := range fn.ParTypes {
		parName := fmt.Sprintf("par%d", i)
		pars[i] = llir.NewParam(parName, fnTypeMap[parType])
	}
	irfn := b.mod.NewFunc(fmt.Sprintf("fn%d", index), fnTypeMap[fn.RetType], pars...)
	b.block = irfn.NewBlock("")
	b.params = pars
	b.parTypes = fn.ParTypes
	b.fns[index] = irfn

	b.setup(len(fn.Instructions))
	for _, instr := range fn.Instructions {
		err := b.addInstruction(instr)
		if err != nil {
			return err
		}
		b.index++
	}

	ret := b.registers[fn.Ret].(Value)
	_, ok := ret.(DynamicValue)
	if ok {
		ret.(DynamicValue).Own(b, -1)
		b.cleanup()
		b.block.NewRet(b.block.NewLoad(stringType, ret.Value())) // String is only dynamic value returnable, so can assume
	} else {
		b.cleanup()
		b.block.NewRet(ret.Value())
	}

	return nil
}

func (b *builder) setup(instrcount int) {
	b.index = 0
	b.autofree = make(map[int]DynamicValue)
	b.autofreeMem = make(map[int]empty)
	b.registers = make([]interface{}, instrcount)
	b.stdv = make(map[string]value.Value)
}

func (b *builder) addGetParam(i *ir.GetParam) error {
	var out Value
	switch b.parTypes[i.Index] {
	case ir.INT:
		out = &Int{Val: b.params[i.Index]}

	case ir.FLOAT:
		out = &Float{Val: b.params[i.Index]}

	case ir.STRING:
		mem := b.block.NewAlloca(stringType)
		par := b.params[i.Index]

		ptr := b.block.NewGetElementPtr(stringType, mem, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
		val := b.block.NewExtractValue(par, 0)
		b.block.NewStore(val, ptr)

		ptr = b.block.NewGetElementPtr(stringType, mem, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
		val = b.block.NewExtractValue(par, 1)
		b.block.NewStore(val, ptr)

		out = newStringFromStruct(mem, b, false)
		out.(DynamicValue).Own(b, b.index)
	}
	b.registers[b.index] = out
	return nil
}

func (b *builder) addFunctionCall(i *ir.FunctionCall) error {
	pars := make([]value.Value, len(i.Params))
	for j, par := range i.Params {
		switch b.ir.Functions[i.Fn].ParTypes[j] {
		case ir.INT, ir.FLOAT:
			pars[j] = b.registers[par].(Value).Value()

		case ir.STRING:
			pars[j] = b.block.NewLoad(stringType, b.registers[par].(Value).Value())
		}
	}
	ret := b.block.NewCall(b.fns[i.Fn], pars...)

	var out Value
	switch b.ir.Functions[i.Fn].RetType {
	case ir.INT:
		out = &Int{Val: ret}

	case ir.FLOAT:
		out = &Int{Val: ret}

	case ir.STRING:
		mem := b.block.NewAlloca(stringType)

		ptr := b.block.NewGetElementPtr(stringType, mem, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
		val := b.block.NewExtractValue(ret, 0)
		b.block.NewStore(val, ptr)

		ptr = b.block.NewGetElementPtr(stringType, mem, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 1))
		val = b.block.NewExtractValue(ret, 1)
		b.block.NewStore(val, ptr)

		out = newStringFromStruct(mem, b, true)

	case ir.ARRAY:
		return errors.New("return ARRAYs is currently not supported")
	}

	b.registers[b.index] = out
	return nil
}
