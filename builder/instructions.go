package builder

import (
	"fmt"
	"reflect"

	"github.com/Nv7-Github/bpp/ir"
)

func (b *builder) addInstruction(instr ir.Instruction) error {
	switch i := instr.(type) {
	case *ir.Const:
		return b.addConst(i)

	case *ir.AllocStatic:
		b.addAllocStatic(i)
		return nil

	case *ir.SetMemory:
		b.addSetMemory(i)
		return nil

	case *ir.GetMemory:
		b.addGetMemory(i)
		return nil

	case *ir.Math:
		b.addMath(i)
		return nil

	case *ir.Print:
		b.addPrint(i)
		return nil

	default:
		return fmt.Errorf("unknown instruction type: %s", reflect.TypeOf(i).String())
	}
}
