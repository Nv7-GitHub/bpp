package ir

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Nv7-Github/bpp/old/parser"
)

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
	Val   int
	Index int
}

func (i *StringIndex) Type() Type {
	return STRING
}

func (i *StringIndex) String() string {
	return fmt.Sprintf("StringIndex: (%d, %d)", i.Val, i.Index)
}

func (i *IR) newIndex(array int, index int) int {
	typ := i.GetInstruction(array).Type()
	if typ == ARRAY {
		var valType Type
		arr, ok := i.GetInstruction(array).(*Array)
		if ok {
			valType = arr.ValType
		} else {
			valType = i.GetInstruction(array).(*GetMemoryDynamic).ValType
		}
		return i.AddInstruction(&ArrayIndex{
			Array: array,
			Index: index,
			typ:   valType,
		})
	}
	return i.AddInstruction(&StringIndex{
		Val:   array,
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

type ArrayLength struct {
	Val int
}

func (a *ArrayLength) String() string {
	return fmt.Sprintf("ArrayLength: %d", a.Val)
}

func (a *ArrayLength) Type() Type {
	return INT
}

type StringLength struct {
	Val int
}

func (a *StringLength) String() string {
	return fmt.Sprintf("StringLength: %d", a.Val)
}

func (a *StringLength) Type() Type {
	return INT
}

func (i *IR) newLength(val int) int {
	typ := i.GetInstruction(val).Type()
	if typ == ARRAY {
		return i.AddInstruction(&ArrayLength{
			Val: val,
		})
	}
	return i.AddInstruction(&StringLength{
		Val: val,
	})
}

func (i *IR) addLength(stmt *parser.LengthStmt) (int, error) {
	val, err := i.AddStmt(stmt.Value)
	if err != nil {
		return 0, err
	}

	return i.newLength(val), nil
}
