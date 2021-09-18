package ir

import (
	"fmt"

	"github.com/Nv7-Github/bpp/parser"
)

var comparisonTiers = map[Type]int{NULL: 0, ARRAY: 1, INT: 2, FLOAT: 3, STRING: 4}

func (i *IR) makeComparable(val1 int, val2 int) (int, int, Type) {
	typ1 := i.GetInstruction(val1).Type()
	typ2 := i.GetInstruction(val2).Type()
	if typ1 == typ2 {
		return val1, val2, typ1
	}

	c1 := comparisonTiers[typ1]
	c2 := comparisonTiers[typ2]
	if c1 > c2 {
		return val1, i.newCast(val2, typ1), typ1
	}
	return i.newCast(val1, typ2), val2, typ2
}

func (i *IR) addComparison(stmt *parser.ComparisonStmt) (int, error) {
	val1, err := i.AddStmt(stmt.Left)
	if err != nil {
		return 0, err
	}
	val2, err := i.AddStmt(stmt.Right)
	if err != nil {
		return 0, err
	}

	var typ Type
	val1, val2, typ = i.makeComparable(val1, val2)
	return i.newCompare(stmt.Operation, val1, val2, typ), nil
}

type Compare struct {
	Op   parser.Operator
	Val1 int
	Val2 int
	typ  Type
}

func (c *Compare) Type() Type {
	return c.typ
}

func (c *Compare) String() string {
	return fmt.Sprintf("Compare<%s>: (%d, %d)", c.Type(), c.Val1, c.Val2)
}

func (i *IR) newCompare(op parser.Operator, val1, val2 int, typ Type) int {
	return i.AddInstruction(&Compare{Op: op, Val1: val1, Val2: val2, typ: typ})
}
