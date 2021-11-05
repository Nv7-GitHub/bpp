package parser

import "fmt"

var parsers map[string]Parser

type Parser struct {
	Parse  func(params []Statement, prog *Program, pos *Pos) (Statement, error)
	Params []Type
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
	Statements []Statement
	Args       []Argument
	RetType    Type
}

type Argument struct {
	Type Type
	Name string
}

type Program struct {
	Functions  map[string]*Function
	VarTypes   map[string]Type
	Statements []Statement
}

func NewProgram() *Program {
	return &Program{
		Functions: make(map[string]*Function),
		VarTypes:  make(map[string]Type),
	}
}

func (p *Program) Parse(code string, filename string) error {
	built, err := ParseCode(code, NewPos(filename), p)
	if err != nil {
		return err
	}
	p.Statements = built
	return nil
}

type Statement interface {
	Type() Type
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
}
