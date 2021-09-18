package ir

import (
	"fmt"

	"github.com/Nv7-Github/bpp/parser"
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
		_, err := ir.AddStmtTop(stmt)
		if err != nil {
			return nil, err
		}
	}
	return ir, nil
}

type Print struct {
	Val int
}

func (p *Print) String() string {
	return fmt.Sprintf("Print: %d", p.Val)
}

func (p *Print) Type() Type {
	return NULL
}

func (i *IR) newPrint(val int) int {
	return i.AddInstruction(&Print{Val: val})
}

func (i *IR) AddStmtTop(stmt parser.Statement) (int, error) {
	ind, err := i.AddStmt(stmt)
	if err != nil {
		return 0, err
	}
	typ := i.GetInstruction(ind).Type()
	if typ != NULL {
		switch typ {
		case STRING:
			i.newPrint(ind)

		case INT, FLOAT:
			casted := i.newCast(ind, STRING)
			i.newPrint(casted)
		}
	}
	return ind, nil
}
