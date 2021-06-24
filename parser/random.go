package parser

// RandintStmt is the equivalent of [RANDINT stmt.Lower stmt.Upper]
type RandintStmt struct {
	*BasicStatement
	Lower Statement
	Upper Statement
}

// Type gives the return type of a RANDINT statement (INT)
func (r *RandintStmt) Type() DataType {
	return INT
}

// RandomStmt is the equivalent of [RANDOM stmt.Lower stmt.Upper]
type RandomStmt struct {
	*BasicStatement
	Lower Statement
	Upper Statement
}

// Type gives the return type of a RANDOM statement (FLOAT)
func (r *RandomStmt) Type() DataType {
	return FLOAT
}

// ChooseStmt is the equivalent of [CHOOSE stmt.Data]
type ChooseStmt struct {
	*BasicStatement
	Data Statement
}

// Type gives the return type of a CHOOSE statement (STRING if the data is a STRING, otherwise ANY)
func (c *ChooseStmt) Type() DataType {
	if c.Data.Type().IsEqual(STRING) {
		return STRING
	}
	return ANY
}

// SetupRandoms adds the RANDINT, RANDOM, CHOOSE, and CHOOSECHAR statements
func SetupRandoms() {
	parsers["RANDINT"] = StatementParser{
		Parse: func(args []Statement, pos *Pos) (Statement, error) {
			return &RandintStmt{
				BasicStatement: &BasicStatement{pos: pos},
				Lower:          args[0],
				Upper:          args[1],
			}, nil
		},
		Signature: []DataType{INT, INT},
	}

	parsers["RANDOM"] = StatementParser{
		Parse: func(args []Statement, pos *Pos) (Statement, error) {
			return &RandomStmt{
				BasicStatement: &BasicStatement{pos: pos},
				Lower:          args[0],
				Upper:          args[1],
			}, nil
		},
		Signature: []DataType{NUMBER, NUMBER},
	}

	parsers["CHOOSE"] = StatementParser{
		Parse: func(args []Statement, pos *Pos) (Statement, error) {
			return &ChooseStmt{
				BasicStatement: &BasicStatement{pos: pos},
				Data:           args[0],
			}, nil
		},
		Signature: []DataType{STRING | ARRAY},
	}

	parsers["CHOOSECHAR"] = StatementParser{
		Parse: func(args []Statement, pos *Pos) (Statement, error) {
			return &ChooseStmt{
				BasicStatement: &BasicStatement{pos: pos},
				Data:           args[0],
			}, nil
		},
		Signature: []DataType{STRING},
	}
}
