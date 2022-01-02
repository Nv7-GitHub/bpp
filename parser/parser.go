package parser

import (
	"github.com/Nv7-Github/bpp/types"
)

var parsers map[string]Parser

type Parser struct {
	Parse  func(params []Statement, prog *Program, pos *types.Pos) (Statement, error)
	Params []types.Type
}

func NewBasicStmt(pos *types.Pos) *BasicStmt {
	return &BasicStmt{Position: pos}
}

type BasicStmt struct {
	Position *types.Pos
}

func (b *BasicStmt) Pos() *types.Pos {
	return b.Position
}

type Function struct {
	Name    string
	Params  []FunctionParam
	RetType types.Type

	Statements []Statement
}

type ExternalFunction struct {
	ParTypes []types.Type
	RetType  types.Type
}

type empty struct{}

type Program struct {
	Functions         map[string]*Function
	ExternalFunctions map[string]ExternalFunction
	VarTypes          map[string]types.Type
	Statements        []Statement

	// Multifile
	Files map[string]string
	Added map[string]empty

	// Functions
	InFunction  bool
	FuncName    string
	OldVarTypes map[string]types.Type
}

type FunctionParam struct {
	Name string
	Type types.Type
}

func (p *Program) Close() {
	p.VarTypes = nil
	p.OldVarTypes = nil
	p.Files = nil
	p.Added = nil
}

type Statement interface {
	Type() types.Type
	Pos() *types.Pos
}

func init() {
	parsers = make(map[string]Parser)
	addVariableParsers()
	addMathStmts()
	addConditionals()
	addManipStmts()
	addConstStmts()
	addLoops()
	addFunctionStmts()
	addMultifileParser()
}
