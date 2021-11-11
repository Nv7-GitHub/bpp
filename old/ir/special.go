package ir

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Nv7-Github/bpp/old/parser"
)

type Concat struct {
	Vals []int
}

func (c *Concat) Type() Type {
	return STRING
}

func (c *Concat) String() string {
	nums := make([]string, len(c.Vals))
	for i, val := range c.Vals {
		nums[i] = strconv.Itoa(val)
	}
	return fmt.Sprintf("Concat: (%s)", strings.Join(nums, ", "))
}

func (i *IR) newConcat(vals []int) int {
	return i.AddInstruction(&Concat{Vals: vals})
}

func (i *IR) addConcat(stmt *parser.ConcatStmt) (int, error) {
	vals := make([]int, len(stmt.Strings))
	for j, val := range stmt.Strings {
		ind, err := i.AddStmt(val)
		if err != nil {
			return 0, err
		}
		vals[j] = ind
	}

	return i.newConcat(vals), nil
}

func getType(typ parser.DataType) Type {
	switch {
	case typ.IsEqual(parser.STRING):
		return STRING
	case typ.IsEqual(parser.ARRAY):
		return ARRAY
	case typ.IsEqual(parser.FLOAT):
		return FLOAT
	case typ.IsEqual(parser.INT):
		return INT
	}
	return NULL
}

func (i *IR) addTypeCast(stmt *parser.TypeCastStmt) (int, error) {
	val, err := i.AddStmt(stmt.Value)
	if err != nil {
		return 0, err
	}

	return i.newCast(val, getType(stmt.NewType)), nil
}

type GetArg struct {
	Index int
}

func (g *GetArg) Type() Type {
	return STRING
}

func (g *GetArg) String() string {
	return fmt.Sprintf("GetArg: %d", g.Index)
}

func (i *IR) newGetArg(index int) int {
	return i.AddInstruction(&GetArg{Index: index})
}

func (i *IR) addArgs(stmt *parser.ArgsStmt) (int, error) {
	ind, err := i.AddStmt(stmt.Index)
	if err != nil {
		return 0, err
	}
	return i.newGetArg(ind), nil
}
