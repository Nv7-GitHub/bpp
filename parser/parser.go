package parser

import "fmt"

type Type int

const (
	INT Type = iota
	FLOAT
	STRING
	ARRAY
	NULL
	VARIADIC
	STATEMENT
)

type Statement interface {
	Type()
}

type Pos struct {
	Line int
	File string
}

func (p *Pos) String() string {
	return fmt.Sprintf("%s:%d", p.File, p.Line)
}

func (p *Pos) NewError(formatter string, options []interface{}) error {
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
