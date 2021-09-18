package ir

import (
	"fmt"

	"github.com/Nv7-Github/Bpp/parser"
)

func NewIR() *IR {
	ir := &IR{
		vars: make(map[string]varData),
	}
	return ir
}

func CreateIR(prog *parser.Program) (*IR, error) {
	ir := NewIR()
	for _, stmt := range prog.Statements {
		ind, err := ir.AddStmt(stmt)
		if err != nil {
			return nil, err
		}
		typ := ir.GetInstruction(ind).Type()
		if typ != NULL {
			switch typ {
			case STRING:
				ir.newPrint(ind)

			case INT, FLOAT:
				casted := ir.newCast(ind, STRING)
				ir.newPrint(casted)
			}
		}
	}
	return ir, nil
}

type print struct {
	Val int
}

func (p *print) String() string {
	return fmt.Sprintf("Print: %d", p.Val)
}

func (p *print) Type() Type {
	return NULL
}

func (i *IR) newPrint(val int) int {
	return i.AddInstruction(&print{Val: val})
}
