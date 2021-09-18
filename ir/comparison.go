package ir

import (
	"fmt"
	"strings"

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

/*func (i *IR) addIf(stmt *parser.IfStmt) (int, error) {
	cond, err := i.AddStmt(stmt.Condition)
	if err != nil {
		return 0, err
	}

	ind := i.Index()
	_, err = i.AddStmt(stmt.Body)
	if err != nil {
		return 0, err
	}
	body := i.Instructions[ind:]
	i.Instructions = i.Instructions[:ind]

	typ1 := body[0].Type()
	typ2 := NULL

	var els []Instruction = nil
	if stmt.Else != nil {
		ind = i.Index()
		_, err = i.AddStmt(stmt.Else)
		if err != nil {
			return 0, err
		}
		els = i.Instructions[ind:]
		i.Instructions = i.Instructions[:ind]
		typ2 = els[0].Type()
	}

	i.newIf(cond, body, els)

	c1 := comparisonTiers[typ1]
	c2 := comparisonTiers[typ2]
	var outType Type
	if c1 > c2 {
		outType = typ1
	} else {
		outType = typ2
	}
}*/

type If struct {
	Body []Instruction
	Else []Instruction
	Cond int
}

func (i *If) Type() Type {
	return NULL
}

func (i *If) String() string {
	out := &strings.Builder{}
	fmt.Fprintf(out, "If<%d>:", i.Cond)
	for _, instr := range i.Body {
		fmt.Fprintf(out, "\t%s", instr.String())
	}
	if i.Else != nil {
		fmt.Fprintf(out, "else:")
		for _, instr := range i.Else {
			fmt.Fprintf(out, "\t%s", instr.String())
		}
	}
	return strings.TrimSpace(out.String())
}

func (i *IR) newIf(cond int, body, els []Instruction) {
	i.AddInstruction(&If{Cond: cond, Body: body, Else: els})
}
