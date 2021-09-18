package ir

import (
	"fmt"

	"github.com/Nv7-Github/Bpp/parser"
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
type allocStatic struct {
	typ Type
}

func (a *allocStatic) Type() Type {
	return a.typ
}
func (a *allocStatic) String() string {
	return fmt.Sprintf("AllocStatic<%s>", a.Type().String())
}

func (i *IR) newAllocStatic(typ Type) int {
	instr := &allocStatic{typ: typ}
	return i.AddInstruction(instr)
}

type setMemory struct {
	Mem   int
	Value int

	typ Type
}

func (s *setMemory) Type() Type {
	return NULL
}

func (s *setMemory) String() string {
	return fmt.Sprintf("SetMemory<%s>: (%d, %d)", s.typ.String(), s.Mem, s.Value)
}

func (i *IR) newSetMemory(mem int, val int) int {
	return i.AddInstruction(&setMemory{
		Mem:   mem,
		Value: val,
	})
}

type getMemory struct {
	Mem int

	typ Type
}

func (s *getMemory) Type() Type {
	return s.typ
}

func (s *getMemory) String() string {
	return fmt.Sprintf("GetMemory<%s>: %d", s.typ.String(), s.Mem)
}

func (i *IR) newGetMemory(mem int, typ Type) int {
	return i.AddInstruction(&getMemory{
		Mem: mem,
		typ: typ,
	})
}
