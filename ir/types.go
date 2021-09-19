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
	Functions    []Function

	vars map[string]varData
	fns  map[string]int
}

func (i *IR) Index() int {
	return len(i.Instructions)
}

func (i *IR) AddInstruction(instr Instruction) int {
	ind := len(i.Instructions)
	i.Instructions = append(i.Instructions, instr)
	return ind
}

func (i *IR) String() string {
	out := &strings.Builder{}
	for fnName, fnId := range i.fns {
		fn := i.Functions[fnId]
		fmt.Fprintf(out, "%d [%s]:\n", fnId, fnName)
		for j, instr := range fn.Instructions {
			fmt.Fprintf(out, "\t%d: %s\n", j, instr)
		}
	}
	if len(i.Functions) > 0 {
		fmt.Fprintln(out)
	}

	for i, instr := range i.Instructions {
		fmt.Fprintf(out, "%d: %s\n", i, instr.String())
	}
	return strings.TrimSpace(out.String())
}

func (i *IR) GetInstruction(index int) Instruction {
	return i.Instructions[index]
}

func (i *IR) SetJmpPoint(index int, target int) {
	jmp := i.Instructions[index].(*Jmp)
	jmp.Target = target
	i.Instructions[index] = jmp
}

func (i *IR) SetCondJmpPoint(index int, targetTrue int, targetFalse int) {
	jmp := i.Instructions[index].(*CondJmp)
	jmp.TargetTrue = targetTrue
	jmp.TargetFalse = targetFalse
	i.Instructions[index] = jmp
}

func Indent(val string) string {
	lines := strings.Split(val, "\n")
	for i, line := range lines {
		lines[i] = "\t" + line
	}
	return strings.Join(lines, "\n")
}

type empty struct{}
