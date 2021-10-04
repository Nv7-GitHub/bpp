package builder

import (
	"github.com/Nv7-Github/bpp/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (b *builder) addRandInt(s *ir.RandInt) {
	min := b.registers[s.Min].(*Int).Value()
	max := b.registers[s.Max].(*Int).Value()

	randV := b.block.NewCall(b.stdFn("rand"))

	var diff value.Value = b.block.NewSub(max, min)
	diff = b.block.NewAdd(diff, constant.NewInt(types.I64, 1))

	var out value.Value = b.block.NewSRem(randV, diff)
	out = b.block.NewAdd(out, min)

	b.registers[b.index] = &Int{Val: b.block.NewSExt(out, types.I64)}
}

// Using this method: https://stackoverflow.com/a/64286825/11388343
func (b *builder) addRandFloat(s *ir.RandFloat) {
	min := b.registers[s.Min].(*Float).Value()
	max := b.registers[s.Max].(*Float).Value()

	rand1 := b.block.NewCall(b.stdFn("rand"))
	rand2 := b.block.NewCall(b.stdFn("rand"))
	mul := b.block.NewMul(rand1, rand2)
	fp := b.block.NewSIToFP(mul, types.Double)
	v := b.block.NewCall(b.stdFn("sin"), fp)

	coeff := b.block.NewFSub(max, min)
	var out value.Value = b.block.NewFMul(v, coeff)
	out = b.block.NewFAdd(out, min)

	b.registers[b.index] = &Float{Val: out}
}
