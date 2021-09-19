package ir

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Nv7-Github/bpp/parser"
)

func createConst(val *parser.Data) (*Const, error) {
	switch {
	case val.Type().IsEqual(parser.INT):
		return &Const{
			typ:  INT,
			Data: val.Data,
		}, nil

	case val.Type().IsEqual(parser.FLOAT):
		return &Const{
			typ:  FLOAT,
			Data: val.Data,
		}, nil

	case val.Type().IsEqual(parser.STRING):
		return &Const{
			typ:  STRING,
			Data: val.Data,
		}, nil

	default:
		return nil, fmt.Errorf("%v: unknown constant type", val.Pos())
	}
}

type Const struct {
	typ  Type
	Data interface{}
}

func (c *Const) Type() Type {
	return c.typ
}

func (c *Const) String() string {
	return fmt.Sprintf("Const<%s>: %v", c.Type().String(), c.Data)
}

type Cast struct {
	val int
	typ Type
}

func (c *Cast) Type() Type {
	return c.typ
}

func (c *Cast) String() string {
	return fmt.Sprintf("Cast<%s>: %d", c.Type().String(), c.val)
}

func (i *IR) newCast(val int, typ Type) int {
	return i.AddInstruction(&Cast{
		val: val,
		typ: typ,
	})
}

type Array struct {
	Vals    []int
	ValType Type
}

func (a *Array) Type() Type {
	return ARRAY
}

func (a *Array) String() string {
	vals := make([]string, len(a.Vals))
	for i, val := range a.Vals {
		vals[i] = strconv.Itoa(val)
	}
	return fmt.Sprintf("Array<%s>: (%s)", a.Type().String(), strings.Join(vals, ", "))
}

func (i *IR) newArray(vals []int, valType Type) int {
	return i.AddInstruction(&Array{
		Vals:    vals,
		ValType: valType,
	})
}

func (i *IR) addArray(stmt *parser.ArrayStmt) (int, error) {
	vals := make([]int, len(stmt.Values))
	typ := NULL
	for j, val := range stmt.Values {
		ind, err := i.AddStmt(val)
		if err != nil {
			return 0, err
		}

		if typ == NULL {
			typ = i.GetInstruction(ind).Type()
		} else {
			if i.GetInstruction(ind).Type() != typ {
				return 0, fmt.Errorf("%v: value of type \"%s\" doesn't match with array type \"%s\"", stmt.Pos(), i.GetInstruction(ind).Type().String(), typ.String())
			}
		}

		vals[j] = ind
	}

	return i.newArray(vals, typ), nil
}
