package ir

import (
	"fmt"

	"github.com/Nv7-Github/bpp/parser"
)

func createConst(val *parser.Data) (*Const, error) {
	switch {
	case val.Type().IsEqual(parser.INT):
		return &Const{
			Typ:  INT,
			Data: val.Data,
		}, nil

	case val.Type().IsEqual(parser.FLOAT):
		return &Const{
			Typ:  FLOAT,
			Data: val.Data,
		}, nil

	case val.Type().IsEqual(parser.STRING):
		return &Const{
			Typ:  STRING,
			Data: val.Data,
		}, nil

	case val.Type().IsEqual(parser.NULL):
		return &Const{
			Typ:  NULL,
			Data: val.Data,
		}, nil

	default:
		return nil, fmt.Errorf("%v: unknown constant type", val.Pos())
	}
}

type Const struct {
	Typ  Type
	Data interface{}
}

func (c *Const) Type() Type {
	return c.Typ
}

func (c *Const) String() string {
	return fmt.Sprintf("Const<%s>: %v", c.Type().String(), c.Data)
}

type Cast struct {
	Val int
	Typ Type
}

func (c *Cast) Type() Type {
	return c.Typ
}

func (c *Cast) String() string {
	return fmt.Sprintf("Cast<%s>: %d", c.Type().String(), c.Val)
}

func (i *IR) newCast(val int, typ Type) int {
	instr := i.GetInstruction(val)
	if instr.Type() == ARRAY && typ == STRING {
		comma := i.AddInstruction(&Const{Typ: STRING, Data: ", "})
		var vals []int
		arr, ok := i.GetInstruction(val).(*Array)
		if ok {
			// Raw array, can print it like this
			if arr.ValType == STRING {
				vals = arr.Vals
			} else {
				vals = make([]int, len(arr.Vals))
				for j, val := range arr.Vals {
					casted := i.newCast(val, STRING)
					vals[j] = casted
				}
			}

			v := make([]int, len(vals)+(len(vals)-1))
			for i := range v {
				if (i % 2) == 0 {
					v[i] = vals[i/2]
				} else {
					v[i] = comma
				}
			}
			str := i.newConcat(v)
			return str
		}
		// need to loop through vals
		blank := i.AddInstruction(&Const{Typ: STRING, Data: ""})
		out := i.newAllocDynamic(blank)
		i.newSetMemoryDynamic(out, blank)
		length := i.newLength(val)
		one := i.AddInstruction(&Const{Typ: INT, Data: 1})
		subbed := i.AddInstruction(&Math{parser.SUBTRACTION, length, one, INT})

		iter := i.newAllocStatic(INT)
		zer := i.AddInstruction(&Const{Typ: INT, Data: 0})
		i.newSetMemory(iter, zer)

		start := i.newJmpPoint()
		cond := i.newCompare(parser.LESS, iter, subbed, INT)
		condJmp := i.newCondJmp(cond)

		begin := i.newJmpPoint()
		v := i.newGetMemoryDynamic(out, STRING)
		ind := i.newIndex(val, iter)
		concated := i.newConcat([]int{v, ind, comma})
		i.newSetMemoryDynamic(out, concated)

		topJmp := i.newJmp()
		i.SetJmpPoint(topJmp, start)
		end := i.newJmpPoint()
		i.SetCondJmpPoint(condJmp, begin, end)

		lastInd := i.newIndex(val, subbed)
		return i.newConcat([]int{out, lastInd})
	}

	return i.AddInstruction(&Cast{
		Val: val,
		Typ: typ,
	})
}
