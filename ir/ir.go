package ir

import (
	"fmt"
	"strings"

	"github.com/Nv7-Github/bpp/types"
)

type Instruction interface {
	String() string
	Type() types.Type
}

type IR struct {
	Instructions []Instruction
	Functions    []Function
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
	for fnId, fn := range i.Functions {
		_, _ = fmt.Fprintf(out, "%d => %d:\n", fnId, fn.Ret)
		for j, instr := range fn.Instructions {
			_, _ = fmt.Fprintf(out, "\t%d: %s\n", j, instr)
		}
	}
	if len(i.Functions) > 0 {
		_, _ = fmt.Fprintln(out)
	}

	for i, instr := range i.Instructions {
		_, _ = fmt.Fprintf(out, "%d: %s\n", i, instr.String())
	}
	return strings.TrimSpace(out.String())
}

func (i *IR) GetInstruction(index int) Instruction {
	return i.Instructions[index]
}
