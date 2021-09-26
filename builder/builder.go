package builder

import (
	"github.com/Nv7-Github/bpp/ir"
	llir "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type builder struct {
	mod   *llir.Module
	fn    *llir.Func
	block *llir.Block

	tmpCount int

	registers []interface{}
	ir        *ir.IR
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

		registers: make([]interface{}, len(ir.Instructions)),
		ir:        ir,
	}
	err := builder.build()
	if err != nil {
		return "", err
	}

	return builder.mod.String(), nil
}

func (b *builder) build() error {
	return nil
}
