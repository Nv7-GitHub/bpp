package ir

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Nv7-Github/bpp/parser"
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
