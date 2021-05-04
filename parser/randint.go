package parser

// RandintStmt is the equivalent of [RANDINT stmt.Lower stmt.Upper]
type RandintStmt struct {
	*BasicStatement
	Lower Statement
	Upper Statement
}

func (r *RandintStmt) Type() DataType {
	return INT
}

func SetupRandoms() {
	parsers["RANDINT"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &RandintStmt{
				BasicStatement: &BasicStatement{line: line},
				Lower:          args[0],
				Upper:          args[1],
			}, nil
		},
		Signature: []DataType{INT, INT},
	}
}
