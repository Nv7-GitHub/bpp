package builder

import (
	"github.com/Nv7-Github/bpp/ir"
	"github.com/llir/irutil"
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

	formatter value.Value

	autofreeCnt int
	autofree    map[int]DynamicValue
}

func Build(ir *ir.IR) (string, error) {
	m := llir.NewModule()
	fn := m.NewFunc("main", types.I32, llir.NewParam("argc", types.I32), llir.NewParam("argv", types.NewPointer(types.I8Ptr)))
	b := fn.NewBlock("")
	formatter := m.NewGlobalDef("format", irutil.NewCString("%s\n"))
	ptr := b.NewGetElementPtr(types.NewArray(4, types.I8), formatter, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))

	builder := &builder{
		mod:   m,
		fn:    fn,
		block: b,

		tmpCount: 0,
		index:    0,

		registers: make([]interface{}, len(ir.Instructions)),
		ir:        ir,
		stdlib:    make(map[string]*llir.Func),

		formatter: ptr,
		autofree:  make(map[int]DynamicValue),
	}
	err := builder.build()
	if err != nil {
		return "", err
	}

	return builder.mod.String(), nil
}

func (b *builder) build() error {
	for _, instr := range b.ir.Instructions {
		err := b.addInstruction(instr)
		if err != nil {
			return err
		}
		b.index++
	}
	for _, val := range b.autofree {
		val.Free(b)
	}
	b.block.NewRet(constant.NewInt(types.I32, 0))
	return nil
}
