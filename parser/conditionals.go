package parser

import "fmt"

func conditionalFuncs() {
	funcs["COMPARE"] = func(args []string, line int) (Executable, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("line %d: invalid argument amount for function %s", line, "COMPARE")
		}
		ex1, err := parseStmt(args[0], line)
		if err != nil {
			return nil, err
		}
		ex2, err := parseStmt(args[1], line)
		if err != nil {
			return nil, err
		}
		ex3, err := parseStmt(args[2], line)
		if err != nil {
			return nil, err
		}
		return func(p *Program) (Variable, error) {
			val1, err := ex1(p)
			if err != nil {
				return Variable{}, err
			}
			op, err := ex2(p)
			if err != nil {
				return Variable{}, err
			}
			val2, err := ex3(p)
			if err != nil {
				return Variable{}, err
			}
			if (op.Type & STRING) != STRING {
				return Variable{}, fmt.Errorf("line %d: parameter 2 of MATH must be string", line)
			}
			if op.Data.(string) == "=" {
				return Variable{
					Type: BOOLEAN,
					Data: val1.Data == val2.Data,
				}, nil
			}
			if op.Data.(string) == "!=" {
				return Variable{
					Type: BOOLEAN,
					Data: val1.Data != val2.Data,
				}, nil
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
			switch op.Data.(string) {
			case ">":
				if isFloat {
					return Variable{
						Type: BOOLEAN,
						Data: val1.Data.(float64) > val2.Data.(float64),
					}, nil
				}
				return Variable{
					Type: BOOLEAN,
					Data: val1.Data.(int) > val2.Data.(int),
				}, nil
			case "<":
				if isFloat {
					return Variable{
						Type: BOOLEAN,
						Data: val1.Data.(float64) < val2.Data.(float64),
					}, nil
				}
				return Variable{
					Type: BOOLEAN,
					Data: val1.Data.(int) < val2.Data.(int),
				}, nil
			case ">=":
				if isFloat {
					return Variable{
						Type: BOOLEAN,
						Data: val1.Data.(float64) >= val2.Data.(float64),
					}, nil
				}
				return Variable{
					Type: BOOLEAN,
					Data: val1.Data.(int) >= val2.Data.(int),
				}, nil
			case "<=":
				if isFloat {
					return Variable{
						Type: BOOLEAN,
						Data: val1.Data.(float64) <= val2.Data.(float64),
					}, nil
				}
				return Variable{
					Type: BOOLEAN,
					Data: val1.Data.(int) <= val2.Data.(int),
				}, nil
			}
			return Variable{}, fmt.Errorf("line %d: invalid operation", line)
		}, nil
	}

	funcs["IF"] = func(args []string, line int) (Executable, error) {
		if len(args) != 3 {
			return nil, fmt.Errorf("line %d: invalid argument amount for function %s", line, "IF")
		}
		opEx, err := parseStmt(args[0], line)
		if err != nil {
			return nil, err
		}
		ex1, err := parseStmt(args[1], line)
		if err != nil {
			return nil, err
		}
		ex2, err := parseStmt(args[2], line)
		if err != nil {
			return nil, err
		}
		return func(p *Program) (Variable, error) {
			op, err := opEx(p)
			if err != nil {
				return Variable{}, err
			}
			if (op.Type & BOOLEAN) != BOOLEAN {
				return Variable{}, fmt.Errorf("line %d: parameter 1 of IF must be a boolean", line)
			}
			if op.Data.(bool) {
				return ex1(p)
			}
			return ex2(p)
		}, nil
	}
}
