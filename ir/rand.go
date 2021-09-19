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
	zero := i.AddInstruction(&Const{Data: 0, typ: INT})
	return i.newRandint(zero, length), nil
}
