package run

import (
	"fmt"
	"reflect"

	"github.com/Nv7-Github/bpp/old/ir"
)

func (r *Runnable) runInstruction(index int) error {
	instr := r.ir.Instructions[index]
	switch i := instr.(type) {
	case *ir.Const:
		r.runConst()
		return nil

	case *ir.Print:
		return r.runPrint(i)

	case *ir.AllocStatic:
		r.runAllocStatic()
		return nil

	case *ir.AllocDynamic:
		r.runAllocDynamic()
		return nil

	case *ir.SetMemory:
		r.runSetMemory(i)
		return nil

	case *ir.SetMemoryDynamic:
		r.runSetMemoryDynamic(i)
		return nil

	case *ir.GetMemory:
		r.runGetMemory(i)
		return nil

	case *ir.GetMemoryDynamic:
		r.runGetMemoryDynamic(i)
		return nil

	case *ir.Math:
		r.runMath(i)
		return nil

	case *ir.Cast:
		return r.runCast(i)

	case *ir.GetArg:
		return r.runGetArg(i)

	case *ir.JmpPoint:
		return nil

	case *ir.Jmp:
		r.runJmp(i)
		return nil

	case *ir.CondJmp:
		r.runCondJmp(i)
		return nil

	case *ir.Compare:
		r.runCompare(i)
		return nil

	case *ir.RandInt:
		r.runRandInt(i)
		return nil

	case *ir.RandFloat:
		r.runRandFloat(i)
		return nil

	case *ir.Concat:
		r.runConcat(i)
		return nil

	case *ir.FunctionCall:
		return r.runFunctionCall(i)

	case *ir.GetParam:
		r.runGetParam(i)
		return nil

	case *ir.Array:
		r.runArray(i)
		return nil

	case *ir.ArrayIndex:
		r.runArrayIndex(i)
		return nil

	case *ir.StringIndex:
		r.runStringIndex(i)
		return nil

	case *ir.PHI:
		r.runPHI(i)
		return nil

	case *ir.ArrayLength:
		r.runArrayLength(i)
		return nil

	case *ir.StringLength:
		r.runStringLength(i)
		return nil

	default:
		return fmt.Errorf("unknown instruction type: %s", reflect.TypeOf(i).String())
	}
}
