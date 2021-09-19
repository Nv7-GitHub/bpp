package ir

import (
	"fmt"
	"strconv"
	"strings"

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
	_, exists := i.Fns[fn.Name]
	if exists {
		return fmt.Errorf("%v: function \"%s\" already defined", fn.Pos(), fn.Name)
	}
	// Add fn name for recursion
	i.Fns[fn.Name] = len(i.Functions)

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
	i.Functions = append(i.Functions, Function{
		ParTypes:     parTypes,
		Ret:          ret,
		Instructions: i.Instructions,
	})

	return nil
}

type Function struct {
	ParTypes     []Type
	Ret          int
	Instructions []Instruction
}

type FunctionCall struct {
	Fn     int
	Params []int
	typ    Type
}

func (f *FunctionCall) Type() Type {
	return f.typ
}

func (f *FunctionCall) String() string {
	args := make([]string, len(f.Params))
	for ind, par := range f.Params {
		args[ind] = strconv.Itoa(par)
	}
	return fmt.Sprintf("FunctionCall<%s>: (%s) => %d", f.Type().String(), strings.Join(args, ", "), f.Fn)
}

func (i *IR) addFunctionCall(stmt *parser.FunctionCallStmt) (int, error) {
	fn, exists := i.Fns[stmt.Name]
	if !exists {
		return 0, fmt.Errorf("%v: function \"%s\" not defined", stmt.Pos(), stmt.Name)
	}

	pars := make([]int, len(stmt.Args))
	for ind, arg := range stmt.Args {
		v, err := i.AddStmt(arg)
		if err != nil {
			return 0, err
		}
		pars[ind] = v
	}

	return i.AddInstruction(&FunctionCall{
		Fn:     fn,
		Params: pars,
		typ:    i.Functions[fn].Instructions[i.Functions[fn].Ret].Type(),
	}), nil
}
