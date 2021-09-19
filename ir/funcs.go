package ir

import (
	"fmt"

	"github.com/Nv7-Github/bpp/parser"
)

type GetParam struct {
	Index int
	typ   Type
}

func (g *GetParam) Type() Type {
	return g.typ
}

func (g *GetParam) String() string {
	return fmt.Sprintf("GetParam<%s>: %d", g.Type().String(), g.Index)
}

func (i *IR) AddFunction(fn *parser.FunctionBlock) error {
	i.Instructions = make([]Instruction, 0)
	i.vars = make(map[string]varData)
	for ind, par := range fn.Signature.Signature {
		typ := getType(par)
		val := i.AddInstruction(&GetParam{Index: ind, typ: typ})

		var mem int
		_, exists := dynamics[typ]
		if exists {
			mem = i.newAllocDynamic(val)
			i.newSetMemoryDynamic(mem, val)
		} else {
			mem = i.newAllocStatic(typ)
			i.newSetMemory(mem, val)
		}

		i.vars[fn.Signature.Names[ind]] = varData{
			Mem: mem,
			Typ: typ,
		}
	}

	for _, stmt := range fn.Body {
		_, err := i.AddStmt(stmt)
		if err != nil {
			return err
		}
	}

	ret, err := i.AddStmt(fn.Return)
	if err != nil {
		return err
	}

	parTypes := make([]Type, len(fn.Signature.Signature))
	for ind, par := range fn.Signature.Signature {
		parTypes[ind] = getType(par)
	}
	i.Functions[fn.Name] = Function{
		ParTypes:     parTypes,
		Ret:          ret,
		Instructions: i.Instructions,
	}

	return nil
}

type Function struct {
	ParTypes     []Type
	Ret          int
	Instructions []Instruction
}
