package types

import (
	"fmt"
	"strings"
)

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

func NewArrayType(valTyp Type) *Array {
	return &Array{
		ValType: valTyp,
	}
}

func (a *Array) String() string {
	return fmt.Sprintf("Array<%s>", a.ValType.String())
}

type MultiType struct {
	Types []Type
}

func NewMultiType(types ...Type) *MultiType {
	m := &MultiType{
		Types: types,
	}
	return m
}

func (m *MultiType) Equal(t Type) bool {
	for _, typ := range m.Types {
		if typ.Equal(t) {
			return true
		}
	}
	return false
}

func (m *MultiType) BasicType() BasicType {
	return STATEMENT
}

func (m *MultiType) String() string {
	out := &strings.Builder{}
	out.WriteString("MultiType<")
	for i, typ := range m.Types {
		out.WriteString(typ.String())
		if i != len(m.Types)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(">")
	return out.String()
}
