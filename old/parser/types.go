package parser

import "fmt"

// NewPos creates a new initialized Pos with the values supplied.
func NewPos(filename string, line int) *Pos {
	return &Pos{
		Filename: filename,
		Line:     line,
	}
}

// Pos defines a position in a B++ program, commonly used in debug prints
type Pos struct {
	Filename string
	Line     int
}

// String allows Pos to implement the Stringer interface
func (p *Pos) String() string {
	return fmt.Sprintf("%s:%d", p.Filename, p.Line)
}

// Statement stores the data for everything in B++
type Statement interface {
	Pos() *Pos
	Type() DataType
}

// Block is a statement that supports being multiple types
type Block interface {
	Pos() *Pos
	Type() DataType

	Keywords() []string
	EndSignature() []DataType
	End(keyword string, arguments []Statement, statements []Statement) (bool, error) // Returns whether closed or not
}

// StatementParser defines the type for a statement parser - a cusotm function that can parse the statement and a signature of its parameters
type StatementParser struct {
	Parse     func(args []Statement, pos *Pos) (Statement, error)
	Signature []DataType
}

// BlockParser defines the type of a block parser - a function to parse the first statement of a block and return a block object based on that, and the signature of the first statement in the block
type BlockParser struct {
	Parse     func(args []Statement, pos *Pos) (Block, error)
	Signature []DataType
}

// BasicStatement allows other statements to implement the Statement interface
type BasicStatement struct {
	pos *Pos
}

// Pos gives the pos of a basic statement
func (b *BasicStatement) Pos() *Pos {
	return b.pos
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

var operatorNames = map[Operator]string{
	EQUAL:          "EQUAL",
	NOTEQUAL:       "NOTEQUAL",
	GREATER:        "GREATER",
	LESS:           "LESS",
	GREATEREQUAL:   "GREATEREQUAL",
	LESSEQUAL:      "LESSEQUAL",
	ADDITION:       "ADDITION",
	SUBTRACTION:    "SUBTRACTION",
	MULTIPLICATION: "MULTIPLICATION",
	DIVISION:       "DIVISION",
	POWER:          "POWER",
}

func (o Operator) String() string {
	return operatorNames[o]
}

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
