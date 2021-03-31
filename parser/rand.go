package parser

import (
	"fmt"
	"math/rand"
)

func randFuncs() {
	funcs["RANDINT"] = func(args []string, line int) (Executable, error) {
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
			val1, err := ex1(p)
			if err != nil {
				return Variable{}, err
			}
			val2, err := ex2(p)
			if err != nil {
				return Variable{}, err
			}
			if (val1.Type & INT) != INT {
				return Variable{}, fmt.Errorf("line %d: parameter 1 of DEFINE must be integer", line)
			}
			if (val2.Type & INT) != INT {
				return Variable{}, fmt.Errorf("line %d: parameter 1 of DEFINE must be integer", line)
			}
			return Variable{
				Type: INT,
				Data: rand.Intn(val2.Data.(int)+val1.Data.(int)) - val1.Data.(int),
			}, nil
		}, nil
	}
	funcs["RANDOM"] = func(args []string, line int) (Executable, error) {
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
			val1, err := ex1(p)
			if err != nil {
				return Variable{}, err
			}
			val2, err := ex2(p)
			if err != nil {
				return Variable{}, err
			}
			if ((val1.Type & FLOAT) != FLOAT) && ((val1.Type & INT) != INT) {
				return Variable{}, fmt.Errorf("line %d: parameter 1 of RANDOM must be float or integer", line)
			}
			if ((val2.Type & FLOAT) != FLOAT) && ((val2.Type & INT) != INT) {
				return Variable{}, fmt.Errorf("line %d: parameter 2 of RANDOM must be float or integer", line)
			}
			if (val1.Type & INT) == INT {
				val1.Type = FLOAT
				val1.Data = float64(val1.Data.(int))
			}
			if (val2.Type & INT) == INT {
				val2.Type = FLOAT
				val2.Data = float64(val1.Data.(int))
			}
			v1 := val1.Data.(float64)
			v2 := val2.Data.(float64)
			return Variable{
				Type: FLOAT,
				Data: v1 + rand.Float64()*(v2-v1),
			}, nil
		}, nil
	}
}
