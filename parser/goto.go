package parser

// GotoStmt is the equivalent of [GOTO stmt.Label]
type GotoStmt struct {
	Label Statement

	line int
}

func (g *GotoStmt) Type() DataType {
	return NULL
}

func (g *GotoStmt) Line() int {
	return g.line
}

// SectionStmt is the equivalent of [SECTION stmt.Label]
type SectionStmt struct {
	Label Statement

	line int
}

func (s *SectionStmt) Type() DataType {
	return NULL
}

func (s *SectionStmt) Line() int {
	return s.line
}

func SetupGotos() {
	parsers["GOTO"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &GotoStmt{
				Label: args[0],
				line:  line,
			}, nil
		},
		Signature: []DataType{STRING},
	}

	parsers["SECTION"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &GotoStmt{
				Label: args[0],
				line:  line,
			}, nil
		},
		Signature: []DataType{STRING},
	}
}
