package ir

import (
	"fmt"

	"github.com/Nv7-Github/bpp/types"
)

// Memory instructions
type AllocStatic struct {
	Typ types.Type
}

func (a *AllocStatic) Type() types.Type {
	return a.Typ
}

func (a *AllocStatic) String() string {
	return fmt.Sprintf("AllocStatic<%s>", a.Type().String())
}

func (i *IR) NewAllocStatic(typ types.Type) int {
	instr := &AllocStatic{Typ: typ}
	return i.AddInstruction(instr)
}

type SetMemory struct {
	Mem   int
	Value int

	Typ types.Type
}

func (s *SetMemory) Type() types.Type {
	return types.NULL
}

func (s *SetMemory) String() string {
	return fmt.Sprintf("SetMemory<%s>: (%d, %d)", s.Typ.String(), s.Mem, s.Value)
}

func (i *IR) NewSetMemory(mem int, val int) int {
	return i.AddInstruction(&SetMemory{
		Mem:   mem,
		Value: val,
		Typ:   i.GetInstruction(val).Type(),
	})
}

type GetMemory struct {
	Mem int

	Typ types.Type
}

func (s *GetMemory) Type() types.Type {
	return s.Typ
}

func (s *GetMemory) String() string {
	return fmt.Sprintf("GetMemory<%s>: %d", s.Typ.String(), s.Mem)
}

func (i *IR) NewGetMemory(mem int) int {
	return i.AddInstruction(&GetMemory{
		Mem: mem,
		Typ: i.GetInstruction(mem).(*AllocStatic).Type(),
	})
}

// Dynamic types
type AllocDynamic struct {
	Typ types.Type
}

func (a *AllocDynamic) Type() types.Type {
	return a.Typ
}

func (a *AllocDynamic) String() string {
	return fmt.Sprintf("AllocDynamic<%s>", a.Type().String())
}

func (i *IR) NewAllocDynamic(typ types.Type) int {
	instr := &AllocDynamic{Typ: typ}
	return i.AddInstruction(instr)
}

type SetMemoryDynamic struct {
	Mem   int
	Value int

	Typ types.Type
}

func (s *SetMemoryDynamic) Type() types.Type {
	return types.NULL
}

func (s *SetMemoryDynamic) String() string {
	return fmt.Sprintf("SetMemoryDynamic<%s>: (%d, %d)", s.Typ.String(), s.Mem, s.Value)
}

func (i *IR) NewSetMemoryDynamic(mem int, val int) int {
	return i.AddInstruction(&SetMemoryDynamic{
		Mem:   mem,
		Value: val,
		Typ:   i.GetInstruction(val).Type(),
	})
}

type GetMemoryDynamic struct {
	Mem int

	Typ types.Type
}

func (s *GetMemoryDynamic) Type() types.Type {
	return s.Typ
}

func (s *GetMemoryDynamic) String() string {
	return fmt.Sprintf("GetMemory<%s>: %d", s.Typ.String(), s.Mem)
}

func (i *IR) NewGetMemoryDynamic(mem int) int {
	m := i.GetInstruction(mem).(*AllocDynamic)
	return i.AddInstruction(&GetMemoryDynamic{
		Mem: mem,
		Typ: m.Type(),
	})
}
