package parser

type Statement interface {
	Line() int
	Type() DataType
}

type StatementParser struct {
	Parse     func(args []Statement, line int) (Statement, error)
	Signature []DataType
}

var parsers map[string]StatementParser
