package ir

import (
	"fmt"

	"github.com/Nv7-Github/bpp/parser"
)

type varData struct {
	Mem int // index of its memory
	Typ Type
}

var dynamics = map[Type]empty{
	STRING: {},
	ARRAY:  {},
}

func (i *IR) addDefine(stmt *parser.DefineStmt) (int, error) {
	valind, err := i.AddStmt(stmt.Value)
	if err != nil {
		return 0, err
	}
	typ := i.GetInstruction(valind).Type()

	// Check if it already exists and if it doesnt, alloc the memory
	name := stmt.Label.(*parser.Data).Data.(string)
	_, exists := i.vars[name]
	if !exists {
		// If dynamic, alloc dynamic
		_, exists := dynamics[typ]
		var mem int
		if !exists {
			mem = i.newAllocStatic(typ)
		} else {
			mem = i.newAllocDynamic(valind)
		}
		i.vars[name] = varData{
			Mem: mem,
			Typ: typ,
		}
	}

	// Overwrite data
	val := i.vars[name]
	_, exists = dynamics[typ]
	if exists {
		return i.newSetMemoryDynamic(val.Mem, valind), nil
	}
	return i.newSetMemory(val.Mem, valind), nil
}

func (i *IR) addVar(stmt *parser.VarStmt) (int, error) {
	name := stmt.Label.(*parser.Data).Data.(string)
	val, exists := i.vars[name]
	if !exists {
		return 0, fmt.Errorf("%v: variable %s not defined", stmt.Pos().String(), name)
	}
	_, dynamic := dynamics[val.Typ]
	if dynamic {
		return i.newGetMemoryDynamic(val.Mem, val.Typ), nil
	}

	return i.newGetMemory(val.Mem, val.Typ), nil
}

// Memory instructions
type AllocStatic struct {
	Typ Type
}

func (a *AllocStatic) Type() Type {
	return a.Typ
}
func (a *AllocStatic) String() string {
	return fmt.Sprintf("AllocStatic<%s>", a.Type().String())
}

func (i *IR) newAllocStatic(typ Type) int {
	instr := &AllocStatic{Typ: typ}
	return i.AddInstruction(instr)
}

type SetMemory struct {
	Mem   int
	Value int

	Typ Type
}

func (s *SetMemory) Type() Type {
	return NULL
}

func (s *SetMemory) String() string {
	return fmt.Sprintf("SetMemory<%s>: (%d, %d)", s.Type().String(), s.Mem, s.Value)
}

func (i *IR) newSetMemory(mem int, val int) int {
	return i.AddInstruction(&SetMemory{
		Mem:   mem,
		Value: val,
		Typ:   i.GetInstruction(val).Type(),
	})
}

type GetMemory struct {
	Mem int

	Typ Type
}

func (s *GetMemory) Type() Type {
	return s.Typ
}

func (s *GetMemory) String() string {
	return fmt.Sprintf("GetMemory<%s>: %d", s.Typ.String(), s.Mem)
}

func (i *IR) newGetMemory(mem int, typ Type) int {
	return i.AddInstruction(&GetMemory{
		Mem: mem,
		Typ: typ,
	})
}

// Dynamic types
type AllocDynamic struct {
	Val int
	Typ Type
}

func (a *AllocDynamic) Type() Type {
	return a.Typ
}

func (a *AllocDynamic) String() string {
	return fmt.Sprintf("AllocDynamic<%s>: %d", a.Type().String(), a.Val)
}

func (i *IR) newAllocDynamic(val int) int {
	v := i.GetInstruction(val)
	instr := &AllocDynamic{
		Val: val,
		Typ: v.Type(),
	}
	return i.AddInstruction(instr)
}

type SetMemoryDynamic struct {
	Mem   int
	Value int

	Typ Type
}

func (s *SetMemoryDynamic) Type() Type {
	return NULL
}

func (s *SetMemoryDynamic) String() string {
	return fmt.Sprintf("SetMemoryDynamic<%s>: (%d, %d)", s.Typ.String(), s.Mem, s.Value)
}

func (i *IR) newSetMemoryDynamic(mem int, val int) int {
	return i.AddInstruction(&SetMemoryDynamic{
		Mem:   mem,
		Value: val,
		Typ:   i.GetInstruction(val).Type(),
	})
}

type GetMemoryDynamic struct {
	Mem int

	Typ Type
}

func (s *GetMemoryDynamic) Type() Type {
	return s.Typ
}

func (s *GetMemoryDynamic) String() string {
	return fmt.Sprintf("GetMemory<%s>: %d", s.Typ.String(), s.Mem)
}

func (i *IR) newGetMemoryDynamic(mem int, typ Type) int {
	return i.AddInstruction(&GetMemoryDynamic{
		Mem: mem,
		Typ: typ,
	})
}
