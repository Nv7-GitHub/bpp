// Package parser parses B++ source code into an AST tree.
package parser

import (
	"fmt"
	"reflect"
	"strings"
)

// Program is the main program, containing the source AST
type Program struct {
	Statements []Statement
}

// Type is there to make a Program implement the Statement interface
func (p *Program) Type() DataType {
	return NULL
}

// Pos is there to make a Program implement the Statement interface
func (p *Program) Pos() *Pos {
	return NewPos("", 0)
}

// Keywords is there to make a Program implement the Block interface
func (p *Program) Keywords() []string {
	return []string{}
}

// End is there to make a Program implement the Block interface
func (p *Program) End(_ string, _ []Statement, stmts []Statement) bool {
	p.Statements = stmts
	return true
}

// EndSignature is there to make a Program implement the Block interface
func (p *Program) EndSignature() []DataType {
	return make([]DataType, 0)
}

// Parse parses B++ source code and returns a parsed program
func Parse(filename, code string) (*Program, error) {
	lns := strings.Split(code, "\n")

	functionTypes = make(map[string]FunctionType)

	prog := &Program{}
	scopes := NewScopeStack()
	scopes.AddScope(NewScope(prog))

	for i, val := range lns {
		stmt, err := ParseStmt(val, NewPos(filename, i+1), scopes)
		if err != nil {
			return nil, err
		}
		if stmt != nil {
			scopes.AddStatement(stmt)
		}
	}

	pScope := scopes.GetScope()
	p, ok := pScope.Block.(*Program)
	if !ok {
		return nil, fmt.Errorf("unterminated block: %s", reflect.TypeOf(pScope.Block))
	}
	scopes.FinishScope("", make([]Statement, 0))
	return p, nil
}

// ParseFiles parses multiple B++ source code files, accepting the name of the main file and a map of filename to source code.
func ParseFiles(mainname string, files map[string]string) (*Program, error) {
	code, exists := files[mainname]
	if !exists {
		return nil, fmt.Errorf("no entry point to program found: %s", mainname)
	}

	// Add IMPORT parser
	parsers["IMPORT"] = StatementParser{
		Parse: func(args []Statement, pos *Pos) (Statement, error) {
			fileDat, ok := args[0].(*Data)
			if !ok {
				return nil, fmt.Errorf("%v: argument to IMPORT must be string literal", pos)
			}
			filename, ok := fileDat.Data.(string)
			if !ok {
				return nil, fmt.Errorf("%v: argument to IMPORT must be string literal", pos)
			}

			_, exists := files[filename]
			if !exists {
				return nil, fmt.Errorf("%v: no such file: %s", pos, filename)
			}
			this := parsers["IMPORT"]
			oldFuncs := functionTypes

			parsed, err := ParseFiles(filename, files)
			if err != nil {
				return nil, err
			}

			// Add functions to current scope
			for name, fn := range functionTypes {
				oldFuncs[name] = fn
			}
			functionTypes = oldFuncs
			parsers["IMPORT"] = this

			return &ImportStmt{
				BasicStatement: &BasicStatement{pos: pos},
				Filename:       filename,
				Statements:     parsed.Statements,
			}, nil
		},
		Signature: []DataType{STRING},
	}

	lns := strings.Split(code, "\n")

	functionTypes = make(map[string]FunctionType)

	prog := &Program{}
	scopes := NewScopeStack()
	scopes.AddScope(NewScope(prog))

	for i, val := range lns {
		stmt, err := ParseStmt(val, NewPos(mainname, i+1), scopes)
		if err != nil {
			return nil, err
		}
		if stmt != nil {
			scopes.AddStatement(stmt)
		}
	}

	pScope := scopes.GetScope()
	p, ok := pScope.Block.(*Program)
	if !ok {
		return nil, fmt.Errorf("unterminated block: %s", reflect.TypeOf(pScope.Block))
	}
	scopes.FinishScope("", make([]Statement, 0))
	delete(parsers, "IMPORT") // Remove IMPORT parser
	return p, nil
}
