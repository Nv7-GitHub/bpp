package builder

import (
	"github.com/Nv7-Github/bpp/ir"
	"github.com/Nv7-Github/bpp/old/parser"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (b *builder) addCompare(s *ir.Compare) {
	var out value.Value
	val1 := b.registers[s.Val1].(Value).Value()
	val2 := b.registers[s.Val2].(Value).Value()

	switch s.Type() {
	case ir.INT:
		switch s.Op {
		case parser.EQUAL:
			out = b.block.NewICmp(enum.IPredEQ, val1, val2)

		case parser.NOTEQUAL:
			out = b.block.NewICmp(enum.IPredNE, val1, val2)

		case parser.LESS:
			out = b.block.NewICmp(enum.IPredSLT, val1, val2)

		case parser.LESSEQUAL:
			out = b.block.NewICmp(enum.IPredSLE, val1, val2)

		case parser.GREATER:
			out = b.block.NewICmp(enum.IPredSGT, val1, val2)

		case parser.GREATEREQUAL:
			out = b.block.NewICmp(enum.IPredSGE, val1, val2)
		}

	case ir.FLOAT:
		switch s.Op {
		case parser.EQUAL:
			out = b.block.NewFCmp(enum.FPredOEQ, val1, val2)

		case parser.NOTEQUAL:
			out = b.block.NewFCmp(enum.FPredONE, val1, val2)

		case parser.LESS:
			out = b.block.NewFCmp(enum.FPredOLT, val1, val2)

		case parser.LESSEQUAL:
			out = b.block.NewFCmp(enum.FPredOLE, val1, val2)

		case parser.GREATER:
			out = b.block.NewFCmp(enum.FPredOGE, val1, val2)

		case parser.GREATEREQUAL:
			out = b.block.NewFCmp(enum.FPredOGE, val1, val2)
		}

	case ir.STRING:
		v1 := b.registers[s.Val1].(*String)
		v1l := v1.Length(b)
		v2 := b.registers[s.Val2].(*String)
		v2l := v2.Length(b)

		// Add null terminators
		v1v := b.block.NewCall(b.stdFn("calloc"), constant.NewInt(types.I64, 0), b.block.NewAdd(v1l, constant.NewInt(types.I64, 1)))
		v2v := b.block.NewCall(b.stdFn("calloc"), constant.NewInt(types.I64, 0), b.block.NewAdd(v2l, constant.NewInt(types.I64, 1)))
		b.block.NewCall(b.stdFn("memcpy"), v1v, v1.StringVal(b), v1l)
		b.block.NewCall(b.stdFn("memcpy"), v2v, v2.StringVal(b), v2l)

		// Compare & cleanup
		cmp := b.block.NewCall(b.stdFn("strcmp"), v1v, v2v)
		b.block.NewCall(b.stdFn("free"), v1v)
		b.block.NewCall(b.stdFn("free"), v2v)

		switch s.Op {
		case parser.EQUAL:
			out = b.block.NewICmp(enum.IPredEQ, cmp, constant.NewInt(types.I64, 0))

		case parser.NOTEQUAL:
			out = b.block.NewICmp(enum.IPredNE, cmp, constant.NewInt(types.I64, 0))

		case parser.LESS:
			out = b.block.NewICmp(enum.IPredSLT, cmp, constant.NewInt(types.I64, 0))

		case parser.LESSEQUAL:
			out = b.block.NewICmp(enum.IPredSLE, cmp, constant.NewInt(types.I64, 0))

		case parser.GREATER:
			out = b.block.NewICmp(enum.IPredSGT, cmp, constant.NewInt(types.I64, 0))

		case parser.GREATEREQUAL:
			out = b.block.NewICmp(enum.IPredSGE, cmp, constant.NewInt(types.I64, 0))
		}
	}

	intV := b.block.NewZExt(out, types.I64)
	b.registers[b.index] = &Int{Val: intV}
}
