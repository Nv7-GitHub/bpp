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

// RandomStmt is the equivalent of [RANDOM stmt.Lower stmt.Upper]
type RandomStmt struct {
	*BasicStatement
	Lower Statement
	Upper Statement
}

func (r *RandomStmt) Type() DataType {
	return FLOAT
}

// ChooseStmt is the equivalent of [CHOOSE stmt.Data]
type ChooseStmt struct {
	*BasicStatement
	Data Statement
}

func (c *ChooseStmt) Type() DataType {
	return ANY
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

	parsers["RANDOM"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &RandomStmt{
				BasicStatement: &BasicStatement{line: line},
				Lower:          args[0],
				Upper:          args[1],
			}, nil
		},
		Signature: []DataType{FLOAT, FLOAT},
	}

	parsers["CHOOSE"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &ChooseStmt{
				BasicStatement: &BasicStatement{line: line},
				Data:           args[0],
			}, nil
		},
		Signature: []DataType{STRING | ARRAY},
	}

	parsers["CHOOSECHAR"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &ChooseStmt{
				BasicStatement: &BasicStatement{line: line},
				Data:           args[0],
			}, nil
		},
		Signature: []DataType{ARRAY},
	}
}
