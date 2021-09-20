package run

import (
	"fmt"
	"reflect"

	"github.com/Nv7-Github/bpp/ir"
)

func (r *Runnable) runInstruction(index int) error {
	instr := r.ir.Instructions[index]
	switch i := instr.(type) {
	case *ir.Const:
		r.runConst(index)
		return nil

	case *ir.Print:
		return r.runPrint(i)

	default:
		return fmt.Errorf("unknown instruction type: %s", reflect.TypeOf(i).String())
	}
}
