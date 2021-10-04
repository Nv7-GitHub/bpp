package ir

import (
	"fmt"

	"github.com/Nv7-Github/bpp/parser"
)

type RandInt struct {
	Min int
	Max int
}

func (r *RandInt) String() string {
	return fmt.Sprintf("RandInt: (%d, %d)", r.Min, r.Max)
}

func (r *RandInt) Type() Type {
	return INT
}

func (i *IR) newRandint(min, max int) int {
	return i.AddInstruction(&RandInt{Min: min, Max: max})
}

func (i *IR) addChoose(stmt *parser.ChooseStmt) (int, error) {
	val, err := i.AddStmt(stmt.Data)
	if err != nil {
		return 0, err
	}

	length := i.newLength(val)
	zero := i.AddInstruction(&Const{Data: 0, Typ: INT})
	ind := i.newRandint(zero, length)
	return i.newIndex(val, ind), nil
}

type RandFloat struct {
	Min int
	Max int
}

func (r *RandFloat) String() string {
	return fmt.Sprintf("RandFloat: (%d, %d)", r.Min, r.Max)
}

func (r *RandFloat) Type() Type {
	return FLOAT
}

func (i *IR) newRandfloat(min, max int) int {
	return i.AddInstruction(&RandFloat{Min: min, Max: max})
}

func (i *IR) addRandom(stmt *parser.RandomStmt) (int, error) {
	min, err := i.AddStmt(stmt.Lower)
	if err != nil {
		return 0, err
	}

	max, err := i.AddStmt(stmt.Upper)
	if err != nil {
		return 0, err
	}

	if i.GetInstruction(min).Type() != FLOAT {
		min = i.newCast(min, FLOAT)
	}

	if i.GetInstruction(max).Type() != FLOAT {
		max = i.newCast(max, FLOAT)
	}

	return i.newRandfloat(min, max), nil
}

func (i *IR) addRandint(stmt *parser.RandintStmt) (int, error) {
	min, err := i.AddStmt(stmt.Lower)
	if err != nil {
		return 0, err
	}

	max, err := i.AddStmt(stmt.Upper)
	if err != nil {
		return 0, err
	}

	if i.GetInstruction(min).Type() != INT {
		min = i.newCast(min, INT)
	}

	if i.GetInstruction(max).Type() != INT {
		max = i.newCast(max, INT)
	}

	return i.newRandint(min, max), nil
}
