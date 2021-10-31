package ir

import (
	"fmt"

	"github.com/Nv7-Github/bpp/old/parser"
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
	Typ  Type
}

func (c *Compare) Type() Type {
	return c.Typ
}

func (c *Compare) String() string {
	return fmt.Sprintf("Compare<%s>: (%d, %d)", c.Type(), c.Val1, c.Val2)
}

func (i *IR) newCompare(op parser.Operator, val1, val2 int, typ Type) int {
	return i.AddInstruction(&Compare{Op: op, Val1: val1, Val2: val2, Typ: typ})
}

func (i *IR) addIf(stmt *parser.IfStmt) (int, error) {
	cond, err := i.AddStmt(stmt.Condition)
	if err != nil {
		return 0, err
	}

	jmp := i.newCondJmp(cond)

	ifTrue := i.newJmpPoint()
	ifTrueVal, err := i.AddStmt(stmt.Body)
	if err != nil {
		return 0, err
	}
	ifTrueEnd := i.newJmp()

	ifFalse := i.newJmpPoint()
	ifFalseVal, err := i.AddStmt(stmt.Else)
	if err != nil {
		return 0, err
	}
	ifFalseEnd := i.newJmp()

	end := i.newJmpPoint()
	i.SetCondJmpPoint(jmp, ifTrue, ifFalse)
	i.SetJmpPoint(ifTrueEnd, end)
	i.SetJmpPoint(ifFalseEnd, end)

	var phiTyp Type
	ifTrueVal, ifFalseVal, phiTyp = i.makeComparable(ifTrueVal, ifFalseVal)
	return i.newPHI(cond, ifTrueVal, ifFalseVal, phiTyp), nil
}

type Jmp struct {
	Target int
}

func (j *Jmp) Type() Type {
	return NULL
}

func (j *Jmp) String() string {
	return fmt.Sprintf("Jmp: %d", j.Target)
}

func (i *IR) newJmp() int {
	return i.AddInstruction(&Jmp{Target: -1})
}

type CondJmp struct {
	Cond        int
	TargetTrue  int
	TargetFalse int
}

func (j *CondJmp) Type() Type {
	return NULL
}

func (j *CondJmp) String() string {
	return fmt.Sprintf("CondJmp: %d => (%d, %d)", j.Cond, j.TargetTrue, j.TargetFalse)
}

func (i *IR) newCondJmp(cond int) int {
	return i.AddInstruction(&CondJmp{Cond: cond, TargetTrue: -1, TargetFalse: -1})
}

type JmpPoint struct{}

func (j *JmpPoint) Type() Type {
	return NULL
}

func (j *JmpPoint) String() string {
	return "JmpPoint"
}

func (i *IR) newJmpPoint() int {
	return i.AddInstruction(&JmpPoint{})
}

type PHI struct {
	Cond     int
	ValTrue  int
	ValFalse int
	Typ      Type
}

func (p *PHI) Type() Type {
	return p.Typ
}

func (p *PHI) String() string {
	return fmt.Sprintf("PHI<%s>: %d => (%d, %d)", p.Type().String(), p.Cond, p.ValTrue, p.ValFalse)
}

func (i *IR) newPHI(cond, valTrue, valFalse int, typ Type) int {
	return i.AddInstruction(&PHI{Cond: cond, ValTrue: valTrue, ValFalse: valFalse, Typ: typ})
}
