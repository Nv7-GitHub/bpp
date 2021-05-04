package parser

// Statements
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

// Parsers
var parsers map[string]StatementParser = make(map[string]StatementParser)

// Operators
type Operator int

const (
	EQUAL          Operator = iota // =
	NOTEQUAL                       // !=
	GREATER                        // >
	LESS                           // <
	GREATEREQUAL                   // >=
	LESSEQUAL                      // <=
	ADDITION                       // +
	SUBTRACTION                    // -
	MULTIPLICATION                 // *
	DIVISION                       // /
	POWER                          // ^
)

// Data Types
type DataType int

func (a DataType) IsEqual(b DataType) bool {
	return (a&b) != 0 || (a&b) == a
}

const (
	STRING     DataType                       = 1 << iota // string
	INT                                                   // int
	FLOAT                                                 // float64
	ARRAY                                                 // []Data
	IDENTIFIER                                            // string
	NULL                                                  // nil
	VARIADIC                                              // Multiple args
	NUMBER     = INT | FLOAT                              // interface{}
	ANY        = STRING | INT | FLOAT | ARRAY             // interface{}
)

type Data struct {
	*BasicStatement
	kind DataType
	Data interface{}
}

func (d *Data) Type() DataType {
	return d.kind
}
