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
	BOOLEAN
	IDENTIFIER
	NULL
	GOTO
)

type Executable func(*Program) (Variable, error)

/*
// Input Code:
for i := 0; i < 100; i++ {
	print(i)
}

// Output Code:
[GOTO main]

[SECTION tmp0]
[IF [COMPARE [VAR i] < [VAR loopLength]] [GOTO tmp1] ""]

[SECTION main]
[DEFINE loopLength 100]
[DEFINE i 0]

[SECTION tmp1]
[DEFINE i [MATH [VAR i] + 1]]
[VAR i]
[GOTO tmp0]
*/

/*
Input Code:
func main() {
	a := 1
	b := 2
	print(add(a, b))
}

func add(a, b int) int {
	return a + b
}

// Output Code
[GOTO main]

[SECTION func1]
[DEFINE tmp0 [MATH [VAR a] + [VAR b]]]

[SECTION main]
[DEFINE a 1]
[DEFINE b 2]
[GOTO func1]
[SECTION tmp1]
[VAR tmp0]
*/
