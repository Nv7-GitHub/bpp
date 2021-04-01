package parser

type Program struct {
	Memory   map[string]Variable
	Program  []Executable
	Args     []string
	Sections map[string]int
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
	IDENTIFIER
	NULL
	GOTO
)

type Executable func(*Program) (Variable, error)
