package parser

import (
	"fmt"
	"math"
)

func mathFunc() {
	funcs["MATH"] = func(args []string, line int) (Executable, error) {
		if len(args) != 3 {
			return Executable{}, fmt.Errorf("line %d: invalid argument amount for function %s", line, "MATH")
		}
		ex1, err := parseStmt(args[0], line)
		if err != nil {
			return Executable{}, err
		}
		ex2, err := parseStmt(args[1], line)
		if err != nil {
			return Executable{}, err
		}
		ex3, err := parseStmt(args[2], line)
		if err != nil {
			return Executable{}, err
		}
		return Executable{
			Exec: func(p *Program) (Variable, error) {
				val1, err := ex1.Exec(p)
				if err != nil {
					return Variable{}, err
				}
				op, err := ex2.Exec(p)
				if err != nil {
					return Variable{}, err
				}
				val2, err := ex3.Exec(p)
				if err != nil {
					return Variable{}, err
				}
				isFloat := false
				if val1.Type == FLOAT || val2.Type == FLOAT {
					isFloat = true
					if val1.Type == INT {
						val1.Data = float64(val1.Data.(int))
					}
					if val2.Type == INT {
						val2.Data = float64(val2.Data.(int))
					}
				}
				if (op.Type & STRING) != STRING {
					return Variable{}, fmt.Errorf("line %d: parameter 2 of MATH must be string", line)
				}
				switch op.Data.(string) {
				case "+":
					if isFloat {
						return Variable{
							Type: FLOAT,
							Data: val1.Data.(float64) + val2.Data.(float64),
						}, nil
					}
					return Variable{
						Type: INT,
						Data: val1.Data.(int) + val2.Data.(int),
					}, nil
				case "-":
					if isFloat {
						return Variable{
							Type: FLOAT,
							Data: val1.Data.(float64) - val2.Data.(float64),
						}, nil
					}
					return Variable{
						Type: INT,
						Data: val1.Data.(int) - val2.Data.(int),
					}, nil
				case "*":
					if isFloat {
						return Variable{
							Type: FLOAT,
							Data: val1.Data.(float64) * val2.Data.(float64),
						}, nil
					}
					return Variable{
						Type: INT,
						Data: val1.Data.(int) * val2.Data.(int),
					}, nil
				case "/":
					if isFloat {
						return Variable{
							Type: FLOAT,
							Data: val1.Data.(float64) / val2.Data.(float64),
						}, nil
					}
					return Variable{
						Type: INT,
						Data: val1.Data.(int) / val2.Data.(int),
					}, nil
				case "^":
					if isFloat {
						return Variable{
							Type: FLOAT,
							Data: math.Pow(val1.Data.(float64), val2.Data.(float64)),
						}, nil
					}
					return Variable{
						Type: INT,
						Data: intPow(val1.Data.(int), val2.Data.(int)),
					}, nil
				}
				return Variable{}, fmt.Errorf("line %d: invalid operation", line)
			},
		}, nil
	}
}

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
