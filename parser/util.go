package parser

import (
	"strconv"
)

var funcs = make(map[string]func([]string, int) (Executable, error))

func setupFuncs() {
	variableFuncs()
	mathFunc()
	conditionalFuncs()
	arrayFuncs()
	randFuncs()
	mathFuncs()
}

func parseVariable(text string) Variable {
	if text[0] == '"' && text[len(text)-1] == '"' {
		return Variable{
			Type: STRING,
			Data: text[1 : len(text)-1],
		}
	}

	num, err := strconv.Atoi(text)
	if err == nil {
		return Variable{
			Type: INT,
			Data: num,
		}
	}

	flt, err := strconv.ParseFloat(text, 64)
	if err == nil {
		return Variable{
			Type: FLOAT,
			Data: flt,
		}
	}

	return Variable{
		Type: STRING | IDENTIFIER,
		Data: text,
	}
}
