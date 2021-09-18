package ir

import (
	"fmt"

	"github.com/Nv7-Github/bpp/parser"
)

type varData struct {
	Mem int // index of its memory
	Typ Type
}

func (i *IR) addDefine(stmt *parser.DefineStmt) (int, error) {
	valind, err := i.AddStmt(stmt.Value)
	if err != nil {
		return 0, err
	}
	typ := i.GetInstruction(valind).Type()

	// Check if it already existsm if it doesnt, alloc the memory
	name := stmt.Label.(*parser.Data).Data.(string)
	_, exists := i.vars[name]
	if !exists {
		mem := i.newAllocStatic(typ)
		i.vars[name] = varData{
			Mem: mem,
			Typ: typ,
		}
	}

	// Overwrite data
	val := i.vars[name]
	return i.newSetMemory(val.Mem, valind), nil
}

func (i *IR) addVar(stmt *parser.VarStmt) (int, error) {
	name := stmt.Label.(*parser.Data).Data.(string)
	val, exists := i.vars[name]
	if !exists {
		return 0, fmt.Errorf("%v: variable %s not defined", stmt.Pos().String(), name)
	}

	return i.newGetMemory(val.Mem, val.Typ), nil
}

// Memory instructions
type AllocStatic struct {
	typ Type
}

func (a *AllocStatic) Type() Type {
	return a.typ
}
func (a *AllocStatic) String() string {
	return fmt.Sprintf("AllocStatic<%s>", a.Type().String())
}

func (i *IR) newAllocStatic(typ Type) int {
	instr := &AllocStatic{typ: typ}
	return i.AddInstruction(instr)
}

type SetMemory struct {
	Mem   int
	Value int

	typ Type
}

func (s *SetMemory) Type() Type {
	return NULL
}

func (s *SetMemory) String() string {
	return fmt.Sprintf("SetMemory<%s>: (%d, %d)", s.typ.String(), s.Mem, s.Value)
}

func (i *IR) newSetMemory(mem int, val int) int {
	return i.AddInstruction(&SetMemory{
		Mem:   mem,
		Value: val,
	})
}

type GetMemory struct {
	Mem int

	typ Type
}

func (s *GetMemory) Type() Type {
	return s.typ
}

func (s *GetMemory) String() string {
	return fmt.Sprintf("GetMemory<%s>: %d", s.typ.String(), s.Mem)
}

func (i *IR) newGetMemory(mem int, typ Type) int {
	return i.AddInstruction(&GetMemory{
		Mem: mem,
		typ: typ,
	})
}
