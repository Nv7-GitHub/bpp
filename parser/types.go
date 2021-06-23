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

// StatementParser defines the type for a statement parser - a cusotm function that can parse the statement and a signature of its parameters
type StatementParser struct {
	Parse     func(args []Statement, line int) (Statement, error)
	Signature []DataType
}

// BlockParser defines the type of a block parser - a function to parse the first statement of a block and return a block object based on that, and the signature of the first statement in the block
type BlockParser struct {
	Parse     func(args []Statement, line int) (Block, error)
	Signature []DataType
}

// BasicStatement allows other statements to implement the Statement interface
type BasicStatement struct {
	line int
}

// Line gives the line of a basic statement
func (b *BasicStatement) Line() int {
	return b.line
}

// Type gives NULL for a basic statement, this method is usually overwritten by the statement embedding this struct
func (b *BasicStatement) Type() DataType {
	return NULL
}

// Parsers
var parsers = make(map[string]StatementParser)
var blocks = make(map[string]BlockParser)

// Operator is an enum for all the B++ math and comparison operators
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

// DataType is an enum for all B++ data types. It also supports combining multiple data types through bit masks. The data type for the B++ Data struct is commented next to the enum value
type DataType int

// IsEqual compares data types using bit masks
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

// Data represents a piece of Data in B++, most often a literal. It implements the Statement interface.
type Data struct {
	*BasicStatement
	kind DataType
	Data interface{}
}

// Type returns the type of a piece of data
func (d *Data) Type() DataType {
	return d.kind
}
