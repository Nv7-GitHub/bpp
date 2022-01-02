package parser

import (
	"fmt"

	"github.com/Nv7-Github/bpp/types"
)

var parsers map[string]Parser

type Parser struct {
	Parse  func(params []Statement, prog *Program, pos *Pos) (Statement, error)
	Params []types.Type
}

func NewBasicStmt(pos *Pos) *BasicStmt {
	return &BasicStmt{Position: pos}
}

type BasicStmt struct {
	Position *Pos
}

func (b *BasicStmt) Pos() *Pos {
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
	Pos() *Pos
}

type Pos struct {
	Line int
	File string
}

func (p *Pos) String() string {
	return fmt.Sprintf("%s:%d", p.File, p.Line+1)
}

func (p *Pos) NewError(formatter string, options ...interface{}) error {
	return fmt.Errorf("%v: "+formatter, append([]interface{}{p}, options...)...)
}

func (p *Pos) NextLine() {
	p.Line++
}

func (p *Pos) Duplicate() *Pos {
	return &Pos{p.Line, p.File}
}

func NewPos(file string) *Pos {
	return &Pos{0, file}
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
