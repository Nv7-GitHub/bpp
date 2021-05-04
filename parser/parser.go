package parser

import "strings"

// Program is the main program, containing the source AST
type Program struct {
	Statements []Statement
}

func (p *Program) Line() int {
	return 1
}

func Parse(code string) (Statement, error) {
	code = strings.ReplaceAll(code, "\n\n", "\n") // Get rid of blank lines
	lns := strings.Split(code, "\n")
	out := make([]Statement, len(lns))

	var err error
	for i, val := range lns {
		out[i], err = ParseLine(val, i+1)
		if err != nil {
			return nil, err
		}
	}
	return &Program{
		Statements: out,
	}, nil
}

func ParseLine(line string, num int) (Statement, error) {
	return nil, nil
}
