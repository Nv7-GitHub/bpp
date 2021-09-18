package ir

import (
	"fmt"

	"github.com/Nv7-Github/Bpp/parser"
)

func (i *IR) addMath(stmt *parser.MathStmt) (int, error) {
	val1, err := i.AddStmt(stmt.Left)
	if err != nil {
		return 0, err
	}
	val2, err := i.AddStmt(stmt.Right)
	if err != nil {
		return 0, err
	}

	v1t := i.GetInstruction(val1).Type()
	v2t := i.GetInstruction(val2).Type()

	var typ Type
	// Cast to appropriate type
	switch {
	case v1t == INT && v2t == INT:
		typ = INT

	case v1t == FLOAT && v2t == FLOAT:
		typ = FLOAT

	case v1t == FLOAT && v2t == INT:
		typ = FLOAT
		val2 = i.newCast(val1, FLOAT)

	case v1t == INT && v2t == FLOAT:
		typ = FLOAT
		val1 = i.newCast(val2, FLOAT)

	default:
		return 0, fmt.Errorf("%v: invalid input to MATH", stmt.Pos())
	}

	// Add Instruction
	return i.AddInstruction(&math{
		Op:   stmt.Operation,
		Val1: val1,
		Val2: val2,
		typ:  typ,
	}), nil
}

type math struct {
	Op   parser.Operator
	Val1 int
	Val2 int
	typ  Type
}

func (m *math) Type() Type {
	return m.typ
}

func (m *math) String() string {
	return fmt.Sprintf("Math<%s, %s>: (%d, %d)", m.typ.String(), m.Op.String(), m.Val1, m.Val2)
}
