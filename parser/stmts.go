package parser

func (p *Program) GetStatement(fnName string, args []Statement, pos *Pos) (Statement, error) {
	parser, exists := parsers[fnName]
	if !exists {
		return nil, pos.NewError("unknown function: %s", fnName)
	}
	err := MatchTypes(args, parser.Params, pos)
	if err != nil {
		return nil, err
	}
	return parser.Parse(args, p, pos)
}

func MatchTypes(a []Statement, b []Type, pos *Pos) error {
	variadiccnt := 0
	for _, val := range b {
		if val.Equal(VARIADIC) {
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
