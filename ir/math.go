package ir

import (
	"fmt"

	"github.com/Nv7-Github/bpp/types"
)

type Math struct {
	Op   types.MathOp
	Val1 int
	Val2 int
	Typ  types.Type
}

func (m *Math) Type() types.Type {
	return m.Typ
}

func (m *Math) String() string {
	return fmt.Sprintf("Math<%s, %s>: (%d, %d)", m.Typ.String(), m.Op.String(), m.Val1, m.Val2)
}

func (i *IR) NewMath(op types.MathOp, val1, val2 int, typ types.Type) int {
	return i.AddInstruction(&Math{
		Op:   op,
		Val1: val1,
		Val2: val2,
		Typ:  typ,
	})
}
