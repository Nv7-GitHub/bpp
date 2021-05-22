package parser

// Statement stores the data for everything in B++
type Statement interface {
	Line() int
	Type() DataType
}

// Block is a statement that supports being multiple types
type Block interface {
	Line() int
	Type() DataType

	Keywords() []string
	EndSignature() []DataType
	End(keyword string, arguments []Statement, statements []Statement) bool // Returns whether closed or not
}

type StatementParser struct {
	Parse     func(args []Statement, line int) (Statement, error)
	Signature []DataType
}

type BlockParser struct {
	Parse     func(args []Statement, line int) (Block, error)
	Signature []DataType
}

type BasicStatement struct {
	line int
}

func (b *BasicStatement) Line() int {
	return b.line
}

func (b *BasicStatement) Type() DataType {
	return NULL
}

// Parsers
var parsers = make(map[string]StatementParser)
var blocks = make(map[string]BlockParser)

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
	PARAMETER                                             // For PARAM statements
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
