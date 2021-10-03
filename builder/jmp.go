package builder

import (
	"github.com/Nv7-Github/bpp/ir"
	llir "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
)

func (b *builder) checkJmpPoint(target int) {
	if b.registers[target] == nil {
		newBlock := b.fn.NewBlock("")
		b.registers[target] = newBlock
	}
}

func (b *builder) addJmp(s *ir.Jmp) {
	b.checkJmpPoint(s.Target)

	blk := b.registers[s.Target].(*llir.Block)
	b.block.NewBr(blk)
}

func (b *builder) addCondJmp(s *ir.CondJmp) {
	b.checkJmpPoint(s.TargetTrue)
	b.checkJmpPoint(s.TargetFalse)

	blkT := b.registers[s.TargetTrue].(*llir.Block)
	blkF := b.registers[s.TargetFalse].(*llir.Block)
	cond := b.block.NewICmp(enum.IPredEQ, b.registers[s.Cond].(*Int).Value(), constant.NewInt(types.I64, 1))
	b.block.NewCondBr(cond, blkT, blkF)
}

func (b *builder) addJmpPoint() {
	b.checkJmpPoint(b.index)

	b.cleanup()
	b.block = b.registers[b.index].(*llir.Block)
}

func (b *builder) addPHI(s *ir.PHI) {
	cond := b.block.NewICmp(enum.IPredSLE, b.registers[s.Cond].(*Int).Value(), constant.NewInt(types.I64, 1))

	ift := b.fn.NewBlock("")
	iff := b.fn.NewBlock("")
	end := b.fn.NewBlock("end")

	b.block.NewCondBr(cond, ift, iff)
	ift.NewBr(end)
	iff.NewBr(end)

	b.block = end

	val1 := b.registers[s.ValTrue].(Value)
	val2 := b.registers[s.ValFalse].(Value)

	val := b.block.NewPhi(llir.NewIncoming(val1.Value(), ift), llir.NewIncoming(val2.Value(), iff))

	var out Value
	switch s.Type() {
	case ir.INT:
		out = &Int{Val: val}

	case ir.FLOAT:
		out = &Float{Val: val}

	case ir.STRING:
		out = newStringFromStruct(val, b, false)

	case ir.ARRAY:
		out = newArrayFromStruct(val, b, make([]Value, 0), false)
	}
	b.registers[b.index] = out
}
