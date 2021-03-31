package parser

import "fmt"

func arrayFuncs() {
	funcs["ARRAY"] = func(args []string, line int) (Executable, error) {
		if len(args) < 1 {
			return nil, fmt.Errorf("line %d: invalid argument amount for function %s", line, "ARRAY")
		}
		exs := make([]Executable, len(args))
		var err error
		for i, arg := range args {
			exs[i], err = parseStmt(arg, line)
			if err != nil {
				return nil, err
			}
		}
		return func(p *Program) (Variable, error) {
			vr := Variable{
				Type: ARRAY,
				Data: make([]Variable, len(exs)),
			}
			for i, ex := range exs {
				vr.Data.([]Variable)[i], err = ex(p)
				if err != nil {
					return Variable{}, err
				}
			}
			return vr, nil
		}, nil
	}

	funcs["INDEX"] = func(args []string, line int) (Executable, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("line %d: invalid argument amount for function %s", line, "INDEX")
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
			arr, err := ex1(p)
			if err != nil {
				return Variable{}, err
			}
			if !arr.Type.IsEqual(ARRAY) && !arr.Type.IsEqual(STRING) {
				return Variable{}, fmt.Errorf("line %d: parameter 1 of INDEX must be array or string", line)
			}
			index, err := ex2(p)
			if err != nil {
				return Variable{}, err
			}
			if !index.Type.IsEqual(INT) {
				if index.Type.IsEqual(FLOAT) {
					index.Type = INT
					index.Data = int(index.Data.(float64))
				} else {
					return Variable{}, fmt.Errorf("line %d: parameter 2 of INDEX must be integer", line)
				}
			}
			if arr.Type.IsEqual(STRING) {
				return Variable{
					Type: STRING,
					Data: string(arr.Data.(string)[index.Data.(int)]),
				}, nil
			}
			return arr.Data.([]Variable)[index.Data.(int)], nil
		}, nil
	}

	funcs["REPEAT"] = func(args []string, line int) (Executable, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("line %d: invalid argument amount for function %s", line, "INDEX")
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
			arr, err := ex1(p)
			if err != nil {
				return Variable{}, err
			}
			if !arr.Type.IsEqual(ARRAY) && !arr.Type.IsEqual(STRING) {
				return Variable{}, fmt.Errorf("line %d: parameter 1 of INDEX must be array or string", line)
			}
			cnt, err := ex2(p)
			if err != nil {
				return Variable{}, err
			}
			if !cnt.Type.IsEqual(INT) {
				if cnt.Type.IsEqual(FLOAT) {
					cnt.Type = INT
					cnt.Data = int(cnt.Data.(float64))
				} else {
					return Variable{}, fmt.Errorf("line %d: parameter 2 of INDEX must be integer", line)
				}
			}
			if arr.Type.IsEqual(STRING) {
				out := ""
				for i := 0; i < cnt.Data.(int); i++ {
					out += arr.Data.(string)
				}
				return Variable{
					Type: STRING,
					Data: out,
				}, nil
			}
			out := make([]Variable, len(arr.Data.([]Variable))*cnt.Data.(int))
			for i := 0; i < len(out); i += len(arr.Data.([]Variable)) {
				for j := 0; j < len(arr.Data.([]Variable)); j++ {
					out[i+j] = arr.Data.([]Variable)[j]
				}
			}
			return Variable{
				Type: ARRAY,
				Data: out,
			}, nil
		}, nil
	}
}
