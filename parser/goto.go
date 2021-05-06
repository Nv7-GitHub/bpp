package parser

import "fmt"

// GotoStmt is the equivalent of [GOTO stmt.Label]
type GotoStmt struct {
	*BasicStatement
	Label Statement
}

func (g *GotoStmt) Type() DataType {
	return NULL
}

// SectionStmt is the equivalent of [SECTION stmt.Label]
type SectionStmt struct {
	*BasicStatement
	Label string
}

func (s *SectionStmt) Type() DataType {
	return NULL
}

func SetupGotos() {
	parsers["GOTO"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &GotoStmt{
				Label:          args[0],
				BasicStatement: &BasicStatement{line: line},
			}, nil
		},
		Signature: []DataType{STRING},
	}

	parsers["SECTION"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			data, ok := args[0].(*Data)
			if !ok {
				return nil, fmt.Errorf("line %d: parameter to SECTION must be static", line)
			}
			label, ok := data.Data.(string)
			if !ok {
				return nil, fmt.Errorf("line %d: parameter to SECTION must be of type string", line)
			}
			return &SectionStmt{
				Label:          label,
				BasicStatement: &BasicStatement{line: line},
			}, nil
		},
		Signature: []DataType{IDENTIFIER},
	}
}
