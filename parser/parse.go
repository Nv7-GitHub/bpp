package parser

import (
	"strings"
)

func ParseCode(code string, pos *Pos) ([]Statement, error) {
	code = strings.TrimSpace(code)
	if len(code) == 0 {
		return make([]Statement, 0), nil
	}
	openBrackets := 0
	openQuotation := false
	stmts := make([]Statement, 0)

	fnName := ""
	arg := ""
	args := make([]string, 0)
	inFnName := true

	for _, char := range code {
		switch char {
		case '[':
			if openBrackets == 0 {
				inFnName = true
			}
			openBrackets++
		case ']':
			openBrackets--

			if openBrackets == 0 && !openQuotation && fnName != "" {
				args = append(args, arg)
				argVals := make([]Statement, 0)
				for _, v := range args {
					v = strings.TrimSpace(v)
					parsed, err := ParseCode(v, pos.Duplicate())
					if err != nil {
						return nil, err
					}
					argVals = append(argVals, parsed...)
				}
				stmt, err := GetStatement(fnName, argVals)
				if err != nil {
					return nil, err
				}
				stmts = append(stmts, stmt)

				fnName = ""
				arg = ""
				args = make([]string, 0)

				continue
			}
		case '"':
			openQuotation = !openQuotation
		case '\n', ' ':
			if char == '\n' {
				pos.NextLine()
			}

			if openBrackets == 1 && !openQuotation {
				if inFnName {
					// [DEFINE a 0]
					//        ^
					fnName = strings.TrimSpace(fnName)
					fnName = fnName[1:]
					inFnName = false
				} else {
					// [DEFINE a 0]
					//          ^
					args = append(args, arg)
					arg = ""
				}

				continue
			}
		}

		if inFnName {
			fnName += string(char)
		} else {
			arg += string(char)
		}
	}

	if len(stmts) == 0 {
		return []Statement{GetConst(code)}, nil
	}

	return stmts, nil
}
