package parser

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

var lineRegex = regexp.MustCompile(`\[([A-Z]+?) (.+)\]`)

func Parse(src string) (*Program, error) {
	lines := strings.Split(src, "\n")
	prg := &Program{
		Memory:  make(map[string]Variable),
		Program: make([]Executable, len(lines)),
	}
	var err error
	for i, line := range lines {
		prg.Program[i], err = parseStmt(line, i+1)
		if err != nil {
			return prg, err
		}
	}
	return prg, nil
}

func parseStmt(src string, line int) (Executable, error) {
	if src[0] != '[' && src[len(src)-1] != ']' {
		vr := parseVariable(src)
		return Executable{
			Exec: func(*Program, []Executable) (Variable, error) {
				return vr, nil
			},
		}, nil
	}
	matches := lineRegex.FindAllStringSubmatch(src, -1)
	if len(matches) < 1 || len(matches[0]) < 3 {
		return Executable{}, fmt.Errorf("line %d: unable to parse", line)
	}
	funcName := matches[0][1]
	inpVals := strings.Split(matches[0][2], " ")
	args := make([]string, 0)
	openBrackets := 0
	openQuotations := 0
	arg := ""
	for _, val := range inpVals {
		openBrackets += strings.Count(val, "[")
		openBrackets -= strings.Count(val, "]")
		openQuotations += strings.Count(val, `"`)
		if (math.Round(float64(openQuotations)/2) == float64(openQuotations/2)) && (openBrackets == 0) {
			args = append(args, arg+val)
			arg = ""
			openBrackets = 0
			openQuotations = 0
		} else {
			arg += val + " "
		}
	}

	fn, exists := funcs[funcName]
	if !exists {
		return Executable{}, fmt.Errorf("line %d: no such function %s", line, funcName)
	}
	return fn(args, line)
}
