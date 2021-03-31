package parser

import "fmt"

func arrayFuncs() {
	funcs["ARRAY"] = func(args []string, line int) (Executable, error) {
		if len(args) < 1 {
			return Executable{}, fmt.Errorf("line %d: invalid argument amount for function %s", line, "ARRAY")
		}
		exs := make([]Executable, len(args))
		var err error
		for i, arg := range args {
			exs[i], err = parseStmt(arg, line)
			if err != nil {
				return Executable{}, err
			}
		}
		return Executable{
			Exec: func(p *Program) (Variable, error) {
				vr := Variable{
					Type: ARRAY,
					Data: make([]Variable, len(exs)),
				}
				for i, ex := range exs {
					vr.Data.([]Variable)[i], err = ex.Exec(p)
					if err != nil {
						return Variable{}, err
					}
				}
				return vr, nil
			},
		}, nil
	}
}
