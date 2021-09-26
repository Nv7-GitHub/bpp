package builder

import (
	"github.com/Nv7-Github/bpp/ir"
	"github.com/Nv7-Github/bpp/parser"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (b *builder) addMath(i *ir.Math) {
	v1 := b.registers[i.Val1].(Value).Value()
	v2 := b.registers[i.Val2].(Value).Value()

	var out value.Value
	switch i.Type() {
	case ir.INT:
		switch i.Op {
		case parser.ADDITION:
			out = b.block.NewAdd(v1, v2)

		case parser.SUBTRACTION:
			out = b.block.NewSub(v1, v2)

		case parser.MULTIPLICATION:
			out = b.block.NewMul(v1, v2)

		case parser.DIVISION:
			out = b.block.NewSDiv(v1, v2)

		case parser.POWER:
			v1f := b.block.NewSIToFP(v1, types.Double)
			v2f := b.block.NewSIToFP(v2, types.Double)
			pow := b.stdFn("pow")
			outF := b.block.NewCall(pow, v1f, v2f)
			outFV := b.block.NewFAdd(outF, constant.NewFloat(types.Double, 0.5))
			out = b.block.NewFPToSI(outFV, types.I64)
		}
		b.registers[b.index] = &Int{Val: out}

	case ir.FLOAT:
		switch i.Op {
		case parser.ADDITION:
			out = b.block.NewFAdd(v1, v2)

		case parser.SUBTRACTION:
			out = b.block.NewFSub(v1, v2)

		case parser.MULTIPLICATION:
			out = b.block.NewFMul(v1, v2)

		case parser.DIVISION:
			out = b.block.NewFDiv(v1, v2)

		case parser.POWER:
			pow := b.stdFn("pow")
			b.block.NewCall(pow, v1, v2)
		}

		b.registers[b.index] = &Float{Val: out}
	}
}
