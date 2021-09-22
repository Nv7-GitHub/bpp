package run

import (
	"math"

	"github.com/Nv7-Github/bpp/ir"
	"github.com/Nv7-Github/bpp/parser"
)

func (r *Runnable) runMath(i *ir.Math) {
	typ := r.ir.Instructions[i.Val1].Type()
	if typ == ir.INT {
		switch i.Op {
		case parser.ADDITION:
			r.registers[r.Index] = r.registers[i.Val1].(int) + r.registers[i.Val2].(int)

		case parser.SUBTRACTION:
			r.registers[r.Index] = r.registers[i.Val1].(int) - r.registers[i.Val2].(int)

		case parser.MULTIPLICATION:
			r.registers[r.Index] = r.registers[i.Val1].(int) * r.registers[i.Val2].(int)

		case parser.DIVISION:
			r.registers[r.Index] = r.registers[i.Val1].(int) / r.registers[i.Val2].(int)

		case parser.POWER:
			r.registers[r.Index] = r.registers[i.Val1].(int) ^ r.registers[i.Val2].(int)
		}

		return
	}

	switch i.Op {
	case parser.ADDITION:
		r.registers[r.Index] = r.registers[i.Val1].(float64) + r.registers[i.Val2].(float64)

	case parser.SUBTRACTION:
		r.registers[r.Index] = r.registers[i.Val1].(float64) - r.registers[i.Val2].(float64)

	case parser.MULTIPLICATION:
		r.registers[r.Index] = r.registers[i.Val1].(float64) * r.registers[i.Val2].(float64)

	case parser.DIVISION:
		r.registers[r.Index] = r.registers[i.Val1].(float64) / r.registers[i.Val2].(float64)

	case parser.POWER:
		r.registers[r.Index] = math.Pow(r.registers[i.Val1].(float64), r.registers[i.Val2].(float64))
	}
}

func (r *Runnable) runCompare(i *ir.Compare) {
	val1 := r.registers[i.Val1]
	val2 := r.registers[i.Val2]

	var out bool
	switch i.Type() {
	case ir.INT:
		switch i.Op {
		case parser.EQUAL:
			out = val1.(int) == val2.(int)

		case parser.NOTEQUAL:
			out = val1.(int) != val2.(int)

		case parser.GREATER:
			out = val1.(int) > val2.(int)

		case parser.GREATEREQUAL:
			out = val1.(int) >= val2.(int)

		case parser.LESS:
			out = val1.(int) < val2.(int)

		case parser.LESSEQUAL:
			out = val1.(int) <= val2.(int)
		}

	case ir.FLOAT:
		switch i.Op {
		case parser.EQUAL:
			out = val1.(float64) == val2.(float64)

		case parser.NOTEQUAL:
			out = val1.(float64) != val2.(float64)

		case parser.GREATER:
			out = val1.(float64) > val2.(float64)

		case parser.GREATEREQUAL:
			out = val1.(float64) >= val2.(float64)

		case parser.LESS:
			out = val1.(float64) < val2.(float64)

		case parser.LESSEQUAL:
			out = val1.(float64) <= val2.(float64)
		}

	case ir.STRING:
		switch i.Op {
		case parser.EQUAL:
			out = val1.(string) == val2.(string)

		case parser.NOTEQUAL:
			out = val1.(string) != val2.(string)

		case parser.GREATER:
			out = val1.(string) > val2.(string)

		case parser.GREATEREQUAL:
			out = val1.(string) >= val2.(string)

		case parser.LESS:
			out = val1.(string) < val2.(string)

		case parser.LESSEQUAL:
			out = val1.(string) <= val2.(string)
		}
	}

	if out {
		r.registers[r.Index] = 1
	} else {
		r.registers[r.Index] = 0
	}
}
