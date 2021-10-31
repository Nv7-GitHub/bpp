package parser

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
	Col  int
}
