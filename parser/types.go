package parser

type Program struct {
	Memory  map[string]Variable
	Program []Executable
}

type Variable struct {
	Data interface{}
	Type Type
}

type Type int

const (
	STRING Type = 1 << iota
	INT
	FLOAT
	ARRAY
	BOOLEAN
	IDENTIFIER
	NULL
)

type Executable func(*Program) (Variable, error)
