package parser

import (
	"strings"
)

func ParseCode(code string, pos *Pos, prog *Program) ([]Statement, error) {
	code = strings.TrimSpace(code)
	if len(code) == 0 {
		return make([]Statement, 0), nil
	}

	inFnName := true
	fnName := ""
	arg := ""
	args := make([]string, 0)

	openBrackets := 0
	openQuotation := false

	stmts := make([]Statement, 0)

	for _, char := range code {
		switch char {
		case '[':
			openBrackets++
			if !inFnName {
				arg += string(char)
			}

		case ']':
			openBrackets--

			if !openQuotation && openBrackets == 0 {
				// Add last arg
				args = append(args, arg)

				argVals := make([]Statement, 0)
				for _, arg := range args {
					argV, err := ParseCode(arg, pos, prog)
					if err != nil {
						return nil, err
					}
					argVals = append(argVals, argV...)
				}
				fnName = strings.TrimSpace(fnName)
				stmt, err := GetStatement(fnName, argVals, prog, pos)
				if err != nil {
					return nil, err
				}
				stmts = append(stmts, stmt)

				inFnName = true
				fnName = ""
				arg = ""
				args = make([]string, 0)
			} else if inFnName {
				fnName += string(char)
			} else {
				arg += string(char)
			}

		case '"':
			openQuotation = !openQuotation
			if inFnName {
				fnName += string(char)
			} else {
				arg += string(char)
			}

		case '\n', ' ':
			if char == '\n' {
				pos.NextLine()
			}

			if !openQuotation && openBrackets == 1 {
				if inFnName {
					inFnName = false
				} else {
					args = append(args, arg)
					arg = ""
				}
			} else {
				if inFnName {
					fnName += string(char)
				} else {
					arg += string(char)
				}
			}

		default:
			if inFnName {
				fnName += string(char)
			} else {
				arg += string(char)
			}
		}
	}

	if len(stmts) == 0 {
		return []Statement{GetConst(code, pos)}, nil
	}
	return stmts, nil
}
