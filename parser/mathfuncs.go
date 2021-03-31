package parser

import (
	"fmt"
	"math"
)

func mathFuncs() {
	funcs["FLOOR"] = func(args []string, line int) (Executable, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("line %d: invalid argument amount for function %s", line, "FLOOR")
		}
		ex1, err := parseStmt(args[0], line)
		if err != nil {
			return nil, err
		}
		return func(p *Program) (Variable, error) {
			val, err := ex1(p)
			if err != nil {
				return Variable{}, err
			}
			if !val.Type.IsEqual(FLOAT) && !val.Type.IsEqual(INT) {
				return Variable{}, fmt.Errorf("line %d: input for FLOOR must be float or integer", line)
			}
			if val.Type.IsEqual(INT) {
				val.Data = float64(val.Data.(int))
			}
			num := math.Floor(val.Data.(float64))
			if val.Type.IsEqual(FLOAT) {
				return Variable{
					Type: FLOAT,
					Data: num,
				}, nil
			}
			return Variable{
				Type: INT,
				Data: int(num),
			}, nil
		}, nil
	}

	funcs["CEIL"] = func(args []string, line int) (Executable, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("line %d: invalid argument amount for function %s", line, "FLOOR")
		}
		ex1, err := parseStmt(args[0], line)
		if err != nil {
			return nil, err
		}
		return func(p *Program) (Variable, error) {
			val, err := ex1(p)
			if err != nil {
				return Variable{}, err
			}
			if !val.Type.IsEqual(FLOAT) && !val.Type.IsEqual(INT) {
				return Variable{}, fmt.Errorf("line %d: input for FLOOR must be float or integer", line)
			}
			if val.Type.IsEqual(INT) {
				val.Data = float64(val.Data.(int))
			}
			num := math.Ceil(val.Data.(float64))
			if val.Type.IsEqual(FLOAT) {
				return Variable{
					Type: FLOAT,
					Data: num,
				}, nil
			}
			return Variable{
				Type: INT,
				Data: int(num),
			}, nil
		}, nil
	}

	funcs["FLOOR"] = func(args []string, line int) (Executable, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("line %d: invalid argument amount for function %s", line, "FLOOR")
		}
		ex1, err := parseStmt(args[0], line)
		if err != nil {
			return nil, err
		}
		return func(p *Program) (Variable, error) {
			val, err := ex1(p)
			if err != nil {
				return Variable{}, err
			}
			if !val.Type.IsEqual(FLOAT) && !val.Type.IsEqual(INT) {
				return Variable{}, fmt.Errorf("line %d: input for FLOOR must be float or integer", line)
			}
			if val.Type.IsEqual(INT) {
				val.Data = float64(val.Data.(int))
			}
			num := math.Round(val.Data.(float64))
			if val.Type.IsEqual(FLOAT) {
				return Variable{
					Type: FLOAT,
					Data: num,
				}, nil
			}
			return Variable{
				Type: INT,
				Data: int(num),
			}, nil
		}, nil
	}
}
