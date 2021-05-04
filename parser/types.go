package parser

type Statement interface {
	Line() int
	Type() DataType
}

type StatementParser struct {
	Parse     func(args []Statement, line int) (Statement, error)
	Signature []DataType
}

type BasicStatement struct {
	line int
}

func (b *BasicStatement) Line() int {
	return b.line
}

var parsers map[string]StatementParser = make(map[string]StatementParser)
