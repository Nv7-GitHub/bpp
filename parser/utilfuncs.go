package parser

import (
	"fmt"
)

func utilFuncs() {
	funcs["ARGS"] = func(args []string, line int) (Executable, error) {
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
			if val.Type.IsEqual(FLOAT) {
				val.Type = INT
				val.Data = int(val.Data.(float64))
			}
			index := val.Data.(int)
			if index >= len(p.Args) || index < 0 {
				return Variable{}, fmt.Errorf("line %d: argument index out of bounds", line)
			}
			return parseVariable(p.Args[index]), nil
		}, nil
	}

	funcs["GOTO"] = func(args []string, line int) (Executable, error) {
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
			if !(varName.Type.IsEqual(STRING) || varName.Type.IsEqual(IDENTIFIER)) {
				return Variable{}, fmt.Errorf("line %d: parameter 1 of VAR must be string or identifier", line)
			}
			return Variable{
				Type: GOTO,
				Data: p.Sections[varName.Data.(string)],
			}, nil
		}, nil
	}
}
