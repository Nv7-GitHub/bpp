package membuild

import (
	"github.com/Nv7-Github/Bpp/parser"
)

// Builds the program from an AST to an array of functions, which can then be executed really quickly and multiple times

type Instruction func(p *Program) (Data, error)

type Program struct {
	Instructions []Instruction
	Memory       map[string]Data
	Sections     map[string]int
	Args         []string
}

type Data struct {
	Type  parser.DataType
	Value interface{}
}

func Build(prog *parser.Program) (*Program, error) {
	p := &Program{
		Instructions: make([]Instruction, len(prog.Statements)),
		Memory:       make(map[string]Data),
		Sections:     make(map[string]int),
	}
	var err error
	for i, stmt := range prog.Statements {
		p.Instructions[i], err = BuildStmt(p, stmt, i)
		if err != nil {
			return p, err
		}
	}
	return p, nil
}

const (
	GOTO parser.DataType = -1
)

func NewBlankData() Data {
	return Data{Type: parser.NULL}
}

func NewBlankInstruction() Instruction {
	return func(p *Program) (Data, error) { return NewBlankData(), nil }
}

func ParserDataToData(d *parser.Data) Data {
	return Data{
		Type:  d.Type(),
		Value: d.Data,
	}
}
