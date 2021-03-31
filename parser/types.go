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
	IDENTIFIER Type = 4
)

type Executable struct {
	Exec   func(*Program, []Executable) (Variable, error)
	Params []Type
}
