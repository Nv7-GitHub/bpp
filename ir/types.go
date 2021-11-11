package ir

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
)

func (a BasicType) Equal(b Type) bool {
	return a == b.BasicType()
}

func (b BasicType) String() string {
	return typeNames[b]
}

func (b BasicType) BasicType() BasicType {
	return b
}

var typeNames = map[BasicType]string{
	INT:    "int",
	FLOAT:  "float",
	STRING: "string",
	ARRAY:  "array",
	NULL:   "null",
}

type ArrayType struct {
	ElemType Type
}

func (a *ArrayType) BasicType() BasicType {
	return ARRAY
}

func (a *ArrayType) Equal(b Type) bool {
	if b.BasicType() != ARRAY {
		return false
	}
	return a.ElemType.Equal(b.(*ArrayType).ElemType)
}

func (a *ArrayType) String() string {
	return fmt.Sprintf("array<%s>", a.ElemType.String())
}
