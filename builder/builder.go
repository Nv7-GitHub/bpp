package builder

import (
	"github.com/Nv7-Github/bpp/ir"
	llir "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type builder struct {
	mod   *llir.Module
	fn    *llir.Func
	block *llir.Block

	tmpCount int
	index    int

	registers []interface{}
	ir        *ir.IR
	stdlib    map[string]*llir.Func
	stdv      map[string]value.Value

	autofreeCnt int
	autofree    map[int]DynamicValue
	autofreeMem map[int]empty
}

func Build(ir *ir.IR) (string, error) {
	m := llir.NewModule()
	fn := m.NewFunc("main", types.I32, llir.NewParam("argc", types.I32), llir.NewParam("argv", types.NewPointer(types.I8Ptr)))
	b := fn.NewBlock("")

	builder := &builder{
		mod:   m,
		fn:    fn,
		block: b,

		tmpCount: 0,
		index:    0,

		registers: make([]interface{}, len(ir.Instructions)),
		ir:        ir,
		stdlib:    make(map[string]*llir.Func),
		stdv:      make(map[string]value.Value),

		autofree:    make(map[int]DynamicValue),
		autofreeMem: make(map[int]empty),
	}
	err := builder.build()
	if err != nil {
		return "", err
	}

	return builder.mod.String(), nil
}

// CALL THIS BEFORE JUMPS
func (b *builder) cleanup() {
	for _, val := range b.autofree {
		val.Free(b, -1)
	}
	for ind := range b.autofreeMem {
		mem := b.registers[ind].(DynamicMem)
		mem.Val.Free(b, mem.Index)
	}
}

func (b *builder) build() error {
	for _, instr := range b.ir.Instructions {
		err := b.addInstruction(instr)
		if err != nil {
			return err
		}
		b.index++
	}
	b.cleanup()
	b.block.NewRet(constant.NewInt(types.I32, 0))
	return nil
}
