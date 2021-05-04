package membuild

import (
	"fmt"

	"github.com/Nv7-Github/Bpp/parser"
)

func GotoStmt(p *Program, stm *parser.GotoStmt) (Instruction, error) {
	label, err := BuildStmt(p, stm.Label)
	if err != nil {
		return nil, err
	}
	return func(p *Program) (Data, error) {
		lTxt, err := label(p)
		if err != nil {
			return NewBlankData(), err
		}
		pos, exists := p.Sections[lTxt.Value.(string)]
		if !exists {
			return NewBlankData(), fmt.Errorf("line %d: unknown section %s", stm.Line(), lTxt.Value.(string))
		}
		return Data{
			Type:  GOTO,
			Value: pos,
		}, nil
	}, nil
}

func DefineStmt(p *Program, stm *parser.DefineStmt) (Instruction, error) {
	label, err := BuildStmt(p, stm.Label)
	if err != nil {
		return nil, err
	}
	val, err := BuildStmt(p, stm.Label)
	if err != nil {
		return nil, err
	}
	return func(p *Program) (Data, error) {
		lTxt, err := label(p)
		if err != nil {
			return NewBlankData(), err
		}
		p.Memory[lTxt.Value.(string)], err = val(p)
		if err != nil {
			return NewBlankData(), err
		}
		return NewBlankData(), nil
	}, nil
}

func VarStmt(p *Program, stm *parser.VarStmt) (Instruction, error) {
	label, err := BuildStmt(p, stm.Label)
	if err != nil {
		return nil, err
	}
	return func(p *Program) (Data, error) {
		lTxt, err := label(p)
		if err != nil {
			return NewBlankData(), err
		}
		return p.Memory[lTxt.Value.(string)], nil
	}, nil
}

func IfStmt(p *Program, stm *parser.IfStmt) (Instruction, error) {
	cond, err := BuildStmt(p, stm.Condition)
	if err != nil {
		return nil, err
	}
	body, err := BuildStmt(p, stm.Body)
	if err != nil {
		return nil, err
	}
	el, err := BuildStmt(p, stm.Else)
	if err != nil {
		return nil, err
	}
	return func(p *Program) (Data, error) {
		cond, err := cond(p)
		if err != nil {
			return NewBlankData(), err
		}
		if cond.Value.(int) == 1 {
			return body(p)
		}
		return el(p)
	}, nil
}
