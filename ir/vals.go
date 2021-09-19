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

// NOTE: Index is an instruction index
type ArrayIndex struct {
	Array int
	Index int
	typ   Type
}

func (i *ArrayIndex) Type() Type {
	return i.typ
}

func (i *ArrayIndex) String() string {
	return fmt.Sprintf("ArrayIndex<%s>: (%d, %d)", i.typ.String(), i.Array, i.Index)
}

// NOTE: Index is an instruction index
type StringIndex struct {
	Array int
	Index int
}

func (i *StringIndex) Type() Type {
	return STRING
}

func (i *StringIndex) String() string {
	return fmt.Sprintf("StringIndex: (%d, %d)", i.Array, i.Index)
}

func (i *IR) newIndex(array int, index int) int {
	typ := i.GetInstruction(array).Type()
	if typ == ARRAY {
		return i.AddInstruction(&ArrayIndex{
			Array: array,
			Index: index,
			typ:   i.GetInstruction(array).(*Array).ValType,
		})
	}
	return i.AddInstruction(&StringIndex{
		Array: array,
		Index: index,
	})
}

func (i *IR) addIndex(stmt *parser.IndexStmt) (int, error) {
	array, err := i.AddStmt(stmt.Value)
	if err != nil {
		return 0, err
	}

	ind, err := i.AddStmt(stmt.Index)
	if err != nil {
		return 0, err
	}

	return i.newIndex(array, ind), nil
}
