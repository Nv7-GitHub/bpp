package parser

import (
	"fmt"
	"math"
)

func mathFuncs() {
	funcs["FLOOR"] = func(args []string, line int) (Executable, error) {
		if len(args) != 1 {
			return Executable{}, fmt.Errorf("line %d: invalid argument amount for function %s", line, "FLOOR")
		}
		ex1, err := parseStmt(args[0], line)
		if err != nil {
			return Executable{}, err
		}
		return Executable{
			Exec: func(p *Program) (Variable, error) {
				val, err := ex1.Exec(p)
				if err != nil {
					return Variable{}, err
				}
				if val.Type != FLOAT && val.Type != INT {
					return Variable{}, fmt.Errorf("line %d: input for FLOOR must be float or integer", line)
				}
				if val.Type == INT {
					val.Data = float64(val.Data.(int))
				}
				num := math.Floor(val.Data.(float64))
				if val.Type == FLOAT {
					return Variable{
						Type: FLOAT,
						Data: num,
					}, nil
				}
				return Variable{
					Type: INT,
					Data: int(num),
				}, nil
			},
		}, nil
	}
}
