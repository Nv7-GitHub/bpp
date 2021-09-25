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
		arr := i.GetInstruction(val).(*Array)
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

	return i.AddInstruction(&Cast{
		Val: val,
		Typ: typ,
	})
}
