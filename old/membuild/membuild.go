// Package membuild converts a B++ AST tree into a series of lambdas which can then be executed.
package membuild

import (
	"github.com/Nv7-Github/bpp/parser"
)

// Builds the program from an AST to an array of functions, which can then be executed really quickly and multiple times

// Instruction stores the data for an instruction, or a callable B++ expression
type Instruction func(p *Program) (Data, error)

// Program stores a compiled B++ program with Variables, Functions, and allows for inputs
type Program struct {
	Instructions []Instruction
	Memory       map[string]Data
	Functions    map[string]Function

	// Inputs
	Args   []string
	Runner func(Data) error
}

// Data represents a value in B++
type Data struct {
	Type  parser.DataType
	Value interface{}
}

// Build compiles a B++ program
func Build(prog *parser.Program) (*Program, error) {
	p := &Program{
		Instructions: make([]Instruction, len(prog.Statements)),
		Memory:       make(map[string]Data),
		Functions:    make(map[string]Function),
		Runner:       func(_ Data) error { return nil },
	}
	var err error
	for i, stmt := range prog.Statements {
		p.Instructions[i], err = BuildStmt(p, stmt)
		if err != nil {
			return p, err
		}
	}
	return p, nil
}

// NewBlankData is a utility function to create a valid NULL value
func NewBlankData() Data {
	return Data{Type: parser.NULL}
}

// NewBlankInstruction is a utility function to create an empty, but valid instruction

// ParserDataToData converts parser's Data format to membuild's Data format
func ParserDataToData(d *parser.Data) Data {
	return Data{
		Type:  d.Type(),
		Value: d.Data,
	}
}
