package parser

type Type interface {
	BasicType() BasicType
	Equal(Type) bool
}

type BasicType int

const (
	INT BasicType = iota
	FLOAT
	STRING
	IDENTIFIER
	ARRAY
	NULL
	VARIADIC
	STATEMENT
)

func (b BasicType) BasicType() BasicType {
	return b
}

func (b BasicType) Equal(t Type) bool {
	return b == t.BasicType()
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
