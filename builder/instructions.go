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

	case *ir.Concat:
		b.addConcat(i)
		return nil

	case *ir.Array:
		b.addArray(i)
		return nil

	case *ir.AllocDynamic:
		b.addAllocDynamic(i)
		return nil

	case *ir.SetMemoryDynamic:
		b.addSetMemoryDynamic(i)
		return nil

	case *ir.GetMemoryDynamic:
		b.addGetMemoryDynamic(i)
		return nil

	case *ir.ArrayLength:
		b.addArrayLength(i)
		return nil

	case *ir.Cast:
		b.addCast(i)
		return nil

	default:
		return fmt.Errorf("unknown instruction type: %s", reflect.TypeOf(i).String())
	}
}
