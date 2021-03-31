package parser

import (
	"fmt"
)

func concatFunc() {
	funcs["CONCAT"] = func(args []string, line int) (Executable, error) {
		if len(args) < 1 {
			return nil, fmt.Errorf("line %d: invalid argument amount for function %s", line, "CONCAT")
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
			vals := make([]Variable, len(exs))
			for i, ex := range exs {
				vals[i], err = ex(p)
				if err != nil {
					return Variable{}, err
				}
			}
			var tp Type = -1
			for _, val := range vals {
				if (val.Type != STRING) && val.Type.IsEqual(STRING) {
					val.Type = STRING
				}
				if val.Type != tp && tp != -1 {
					return Variable{}, fmt.Errorf("line %d: all arguments to CONCAT must be of the same type", line)
				} else if tp == -1 && val.Type != tp {
					tp = val.Type
				}
			}
			if (tp != ARRAY) && (tp != STRING) {
				return Variable{}, fmt.Errorf("line %d: CONCAT only accepts array and string", line)
			}
			if tp == ARRAY {
				out := make([]Variable, 0)
				for _, val := range vals {
					out = append(out, val.Data.([]Variable)...)
				}
				return Variable{
					Type: ARRAY,
					Data: out,
				}, nil
			}
			out := ""
			for _, val := range vals {
				out += val.Data.(string)
			}
			return Variable{
				Type: STRING,
				Data: out,
			}, nil
		}, nil
	}
}
