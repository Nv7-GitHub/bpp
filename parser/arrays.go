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

	funcs["INDEX"] = func(args []string, line int) (Executable, error) {
		if len(args) != 2 {
			return Executable{}, fmt.Errorf("line %d: invalid argument amount for function %s", line, "INDEX")
		}
		ex1, err := parseStmt(args[0], line)
		if err != nil {
			return Executable{}, err
		}
		ex2, err := parseStmt(args[1], line)
		if err != nil {
			return Executable{}, err
		}
		return Executable{
			Exec: func(p *Program) (Variable, error) {
				arr, err := ex1.Exec(p)
				if err != nil {
					return Variable{}, err
				}
				if (arr.Type & ARRAY) != ARRAY {
					return Variable{}, fmt.Errorf("line %d: parameter 1 of INDEX must be array", line)
				}
				index, err := ex2.Exec(p)
				if err != nil {
					return Variable{}, err
				}
				if (index.Type & INT) != INT {
					return Variable{}, fmt.Errorf("line %d: parameter 2 of INDEX must be integer", line)
				}
				return arr.Data.([]Variable)[index.Data.(int)], nil
			},
		}, nil
	}
}
