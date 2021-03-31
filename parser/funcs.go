package parser

var funcs map[string]func([]string, int) (Executable, error)
