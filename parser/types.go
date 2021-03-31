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
	STRING     Type = 0
	INT        Type = 1
	FLOAT      Type = 2
	ARRAY      Type = 3
	BOOLEAN    Type = 4
	IDENTIFIER Type = 5
	NULL       Type = 6
)

type Executable struct {
	Exec   func(*Program) (Variable, error)
	Params []Type
}
