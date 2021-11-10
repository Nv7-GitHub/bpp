package parser

import (
	"strings"
)

func (p *Program) ParseCode(code string, pos *Pos) ([]Statement, error) {
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

				// Fn Name
				fnName = strings.TrimSpace(fnName)
				// FUNCTION OPEN
				if fnName == "FUNCTION" {
					if len(args) < 1 {
						return nil, pos.NewError("function name not specified")
					}
					err := p.BeginFunction(args[0], pos)
					if err != nil {
						return nil, err
					}
				}

				// Parse args recursively
				argVals := make([]Statement, 0)
				for _, arg := range args {
					argV, err := p.ParseCode(arg, pos.Duplicate())
					if err != nil {
						return nil, err
					}
					if argV == nil {
						return nil, pos.NewError("expected argument")
					}
					argVals = append(argVals, argV...)
				}
				stmt, err := p.GetStatement(fnName, argVals, pos)
				if err != nil {
					return nil, err
				}
				if fnName != "FUNCTION" {
					stmts = append(stmts, stmt)
				}

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
