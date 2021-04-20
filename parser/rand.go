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
			if !val1.Type.IsEqual(INT) {
				return Variable{}, fmt.Errorf("line %d: parameter 1 of DEFINE must be integer", line)
			}
			if !val2.Type.IsEqual(INT) {
				return Variable{}, fmt.Errorf("line %d: parameter 1 of DEFINE must be integer", line)
			}
			return Variable{
				Type: INT,
				Data: rand.Intn(val2.Data.(int)-val1.Data.(int)) + val1.Data.(int),
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
			if !val1.Type.IsEqual(FLOAT) && !val1.Type.IsEqual(INT) {
				return Variable{}, fmt.Errorf("line %d: parameter 1 of RANDOM must be float or integer", line)
			}
			if !val2.Type.IsEqual(FLOAT) && !val2.Type.IsEqual(INT) {
				return Variable{}, fmt.Errorf("line %d: parameter 2 of RANDOM must be float or integer", line)
			}
			if val1.Type.IsEqual(INT) {
				val1.Type = FLOAT
				val1.Data = float64(val1.Data.(int))
			}
			if val2.Type.IsEqual(INT) {
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

	funcs["CHOOSE"] = func(args []string, line int) (Executable, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("line %d: invalid argument amount for function %s", line, "CHOOSE")
		}
		ex1, err := parseStmt(args[0], line)
		if err != nil {
			return nil, err
		}
		return func(p *Program) (Variable, error) {
			arr, err := ex1(p)
			if err != nil {
				return Variable{}, err
			}
			if !arr.Type.IsEqual(ARRAY) && !arr.Type.IsEqual(STRING) {
				return Variable{}, fmt.Errorf("line %d: parameter 1 of CHOOSE must be array or string", line)
			}
			if arr.Type.IsEqual(STRING) {
				index := rand.Intn(len(arr.Data.(string)))
				return Variable{
					Type: STRING,
					Data: string(arr.Data.(string)[index]),
				}, nil
			}
			index := rand.Intn(len(arr.Data.([]Variable)))
			return arr.Data.([]Variable)[index], nil
		}, nil
	}
	funcs["CHOOSECHAR"] = funcs["CHOOSE"]
}
