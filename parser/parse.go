package parser

import (
	"fmt"
	"strings"
)

// ParseStmt parses a B++ statement's source code and returns the parsed statement
func ParseStmt(line string, pos *Pos, scope ...*ScopeStack) (Statement, error) {
	if strings.ContainsRune(line, '#') {
		line = line[:strings.IndexRune(line, '#')]
	}
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return nil, nil
	}
	if line[0] == '[' && line[len(line)-1] == ']' {
		split := strings.SplitN(line[1:], " ", 2)
		funcName := split[0]
		if len(split) == 1 {
			funcName = funcName[:len(funcName)-1]
		}

		parser, hasParser := parsers[funcName]
		var block Block
		var bParser BlockParser
		var fnType FunctionType
		isBParser := 0

		if !hasParser {
			// Is it a custom function
			var exists bool
			fnType, exists = functionTypes[funcName]
			if !exists {
				// No parser, is it a block end?
				if len(scope) > 0 {
					s := scope[0].GetScope()
					if !s.HasKeyword(funcName) {
						// Not a block end, is it a block start?
						bparser, exists := blocks[funcName]
						if exists {
							isBParser = 1
							bParser = bparser
						} else {
							return nil, fmt.Errorf("%v: No such function '%s'", pos, funcName)
						}
					} else {
						block = s.Block
					}
				} else {
					return nil, fmt.Errorf("%v: No such function '%s'", pos, funcName)
				}
			} else {
				isBParser = 2
			}
		}

		args := make([]string, 0)
		openedBrackets := 0
		openQuotation := false
		argTxt := strings.TrimSpace(line[len(funcName)+1 : len(line)-1])
		arg := ""

		for i, char := range argTxt {
			arg += string(char)

			switch char {
			case '[':
				openedBrackets++
			case ']':
				openedBrackets--
			case '"':
				openQuotation = !openQuotation
			}

			if (char == ' ' || i == len(argTxt)-1) && openedBrackets == 0 && !openQuotation {
				args = append(args, arg)
				arg = ""
				continue
			}
		}

		// Type checking
		argDat, err := ParseArgs(args, pos)
		if err != nil {
			return nil, err
		}

		if hasParser {
			err = MatchTypes(argDat, pos, parser.Signature)
			if err != nil {
				return nil, err
			}

			return parser.Parse(argDat, pos)
		} else if isBParser == 2 {
			err = MatchTypes(argDat, pos, fnType.Signature)
			if err != nil {
				return nil, err
			}

			return &FunctionCallStmt{
				BasicStatement: &BasicStatement{pos: pos},
				ReturnType:     fnType.ReturnType,
				Name:           funcName,
				Args:           argDat,
			}, nil
		} else if isBParser == 1 {
			err = MatchTypes(argDat, pos, bParser.Signature)
			if err != nil {
				return nil, err
			}

			block, err = bParser.Parse(argDat, pos)
			if err != nil {
				return nil, err
			}

			s := NewScope(block)
			scope[0].AddScope(s)

			return nil, nil
		} else {
			err = MatchTypes(argDat, pos, block.EndSignature())
			if err != nil {
				return nil, err
			}

			scope[0].FinishScope(funcName, argDat)

			return nil, nil
		}
	}
	return ParseData(line, pos), nil
}
