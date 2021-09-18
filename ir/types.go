package ir

import (
	"fmt"
	"strings"
)

type Type int

const (
	INT Type = iota
	FLOAT
	STRING
	ARRAY
	NULL
)

var typeNames = map[Type]string{
	INT:    "int",
	FLOAT:  "float",
	STRING: "string",
	ARRAY:  "array",
	NULL:   "null",
}

func (t Type) String() string {
	return typeNames[t]
}

type Instruction interface {
	fmt.Stringer

	Type() Type
}

type IR struct {
	Instructions []Instruction

	vars map[string]varData
}

func (i *IR) AddInstruction(instr Instruction) int {
	ind := len(i.Instructions)
	i.Instructions = append(i.Instructions, instr)
	return ind
}

func (i *IR) String() string {
	out := &strings.Builder{}
	for i, instr := range i.Instructions {
		fmt.Fprintf(out, "%d: %s\n", i, instr.String())
	}
	return strings.TrimSpace(out.String())
}

func (i *IR) GetInstruction(index int) Instruction {
	return i.Instructions[index]
}

func Indent(val string) string {
	lines := strings.Split(val, "\n")
	for i, line := range lines {
		lines[i] = "\t" + line
	}
	return strings.Join(lines, "\n")
}
