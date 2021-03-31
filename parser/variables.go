package parser

import (
	"fmt"
)

func variableFuncs() {
	funcs["DEFINE"] = func(args []string, line int) (Executable, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("line %d: invalid argument amount for function %s", line, "DEFINE")
		}
		ex1, err := parseStmt(args[0], line)
		if err != nil {
			return nil, err
		}
		ex2, err := parseStmt(args[1], line)
		if err != nil {
			return nil, err
		}
		return func(p *Program) (Variable, error) {
			varName, err := ex1(p)
			if err != nil {
				return Variable{}, err
			}
			val, err := ex2(p)
			if err != nil {
				return Variable{}, err
			}
			if !varName.Type.IsEqual(IDENTIFIER) {
				return Variable{}, fmt.Errorf("line %d: parameter 1 of DEFINE must be identifier", line)
			}
			p.Memory[varName.Data.(string)] = val
			return Variable{
				Type: NULL,
			}, nil
		}, nil
	}

	funcs["VAR"] = func(args []string, line int) (Executable, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("line %d: invalid argument amount for function %s", line, "VAR")
		}
		ex1, err := parseStmt(args[0], line)
		if err != nil {
			return nil, err
		}
		return func(p *Program) (Variable, error) {
			varName, err := ex1(p)
			if err != nil {
				return Variable{}, err
			}
			if !varName.Type.IsEqual(IDENTIFIER) {
				return Variable{}, fmt.Errorf("line %d: parameter 1 of VAR must be identifier", line)
			}
			return p.Memory[varName.Data.(string)], nil
		}, nil
	}
}
