package compiler

import (
	"fmt"

	"github.com/Nv7-Github/bpp/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var funcs map[string]*ir.Func

var paramTypeMap = map[parser.DataType]types.Type{
	parser.INT:    types.I64,
	parser.FLOAT:  types.Double,
	parser.STRING: types.I8Ptr,
	parser.NULL:   types.Void,
}

func AddFunction(fn *parser.FunctionBlock) error {
	tmpUsed = 0
	variables = make(map[string]Variable)
	autofree = make(map[string]empty)

	// Make params
	params := make([]*ir.Param, len(fn.Signature.Signature))
	for i, par := range fn.Signature.Signature {
		params[i] = ir.NewParam(fn.Signature.Names[i], paramTypeMap[par])
	}
	retType, exists := paramTypeMap[fn.Signature.ReturnType]
	if !exists {
		return fmt.Errorf("%v: unknown return type", fn.Return.Pos())
	}

	function := m.NewFunc(fn.Name, retType, params...)

	funcs[fn.Name] = function

	initBlock = function.NewBlock("init")
	entry := function.NewBlock("entry")
	block := entry

	// Load params to variables
	for i, par := range params {
		val := block.NewAlloca(par.Type())
		block.NewStore(par, val)
		variables[fn.Signature.Names[i]] = Variable{
			Val:  val,
			Type: val.Type(),
		}
	}

	var err error
	block, err = CompileBlock(fn.Body, block)
	if err != nil {
		return err
	}

	initBlock.NewBr(entry)

	for val := range autofree {
		block.NewCall(free, block.NewLoad(types.I8Ptr, variables[val].Val))
	}

	var ret value.Value
	ret, block, err = CompileStmt(fn.Return, block)
	if err != nil {
		return err
	}

	block.NewRet(ret)

	return nil
}

func CompileFunctionCall(stm *parser.FunctionCallStmt, block *ir.Block) (value.Value, *ir.Block, error) {
	pars := make([]value.Value, len(stm.Args))
	var ar value.Value
	var err error
	for i, arg := range stm.Args {
		ar, block, err = CompileStmt(arg, block)
		if err != nil {
			return nil, block, err
		}

		pars[i] = ar
	}

	res := block.NewCall(funcs[stm.Name], pars...)

	return res, block, nil
}
