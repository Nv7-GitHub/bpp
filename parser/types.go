package parser

import "fmt"

type Type interface {
	BasicType() BasicType
	Equal(Type) bool
	String() string
}

type BasicType int

const (
	INT BasicType = iota
	FLOAT
	STRING
	ARRAY
	NULL
	VARIADIC
	STATEMENT
)

func (b BasicType) BasicType() BasicType {
	return b
}

func (b BasicType) Equal(t Type) bool {
	if b == STATEMENT {
		return true
	}
	return b == t.BasicType()
}

var typeNames = map[BasicType]string{
	INT:    "int",
	FLOAT:  "float",
	STRING: "string",
	ARRAY:  "array",
	NULL:   "null",
}

func (b BasicType) String() string {
	return typeNames[b]
}

type Array struct {
	ValType Type
}

func (a *Array) BasicType() BasicType {
	return ARRAY
}

func (a *Array) Equal(t Type) bool {
	if t.BasicType() != ARRAY {
		return false
	}
	return a.ValType.Equal(t.(*Array).ValType)
}

func (a *Array) String() string {
	return fmt.Sprintf("Array<%s>", a.BasicType().String())
}
