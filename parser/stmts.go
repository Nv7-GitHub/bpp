package parser

import "github.com/Nv7-Github/bpp/types"

func (p *Program) GetStatement(fnName string, args []Statement, pos *types.Pos) (Statement, error) {
	parser, exists := parsers[fnName]
	if !exists {
		// Function call?
		_, exists := p.Functions[fnName]
		if exists {
			// Function call!
			fn := p.Functions[fnName]
			// Check types
			parTypes := make([]types.Type, len(fn.Params))
			for i, par := range fn.Params {
				parTypes[i] = par.Type
			}
			err := MatchTypes(args, parTypes, pos)
			if err != nil {
				return nil, err
			}
			return &FunctionCallStmt{
				FnName:  fnName,
				Params:  args,
				RetType: fn.RetType,
			}, nil
		}

		// Not function call, external function call?
		extFn, exists := p.ExternalFunctions[fnName]
		if exists {
			// External function call!
			err := MatchTypes(args, extFn.ParTypes, pos)
			if err != nil {
				return nil, err
			}

			return &ExternalCallStmt{
				FnName:  fnName,
				Params:  args,
				RetType: extFn.RetType,
			}, nil
		}
		return nil, pos.NewError("unknown function: %s", fnName)
	}
	err := MatchTypes(args, parser.Params, pos)
	if err != nil {
		return nil, err
	}
	return parser.Parse(args, p, pos)
}

func MatchTypes(a []Statement, b []types.Type, pos *types.Pos) error {
	variadiccnt := 0
	for _, val := range b {
		if val.Equal(types.VARIADIC) {
			variadiccnt++
		}
	}

	// Not variadic
	if variadiccnt == 0 {
		if len(a) != len(b) {
			return pos.NewError("expected %d arguments, got %d", len(b), len(a))
		}
		for i, par := range a {
			if !b[i].Equal(par.Type()) {
				return pos.NewError("expected type %s, got %s in argument %d", b[i].String(), par.Type().String(), i+1)
			}
		}
		return nil
	}

	// it is variadic
	if len(a) < len(b)-variadiccnt {
		return pos.NewError("expected at least %d arguments, got %d", len(b)-variadiccnt, len(a))
	}
	// TODO: Figure this out
	return nil
}
