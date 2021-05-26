package parser

// ArgsStmt is the equivalent of [ARGS stmt.Index]
type ArgsStmt struct {
	*BasicStatement
	Index Statement
}

func (a *ArgsStmt) Type() DataType {
	return ANY
}

// ConcatStmt is the equivalent of [CONCAT stmt.Strings]
type ConcatStmt struct {
	*BasicStatement
	Strings []Statement
}

func (c *ConcatStmt) Type() DataType {
	return STRING
}

// RepeatStmt is the equivalent of [REPEAT stmt.Val]
type RepeatStmt struct {
	*BasicStatement
	Val   Statement
	Count Statement
}

func (r *RepeatStmt) Type() DataType {
	return STRING
}

func SetupOthers() {
	parsers["ARGS"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &ArgsStmt{
				Index:          args[0],
				BasicStatement: &BasicStatement{line: line},
			}, nil
		},
		Signature: []DataType{INT},
	}

	parsers["CONCAT"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &ConcatStmt{
				Strings:        args,
				BasicStatement: &BasicStatement{line: line},
			}, nil
		},
		Signature: []DataType{ANY, VARIADIC},
	}

	parsers["REPEAT"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &RepeatStmt{
				Val:            args[0],
				Count:          args[1],
				BasicStatement: &BasicStatement{line: line},
			}, nil
		},
		Signature: []DataType{STRING, INT},
	}

	parsers["NULL"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &Data{
				BasicStatement: &BasicStatement{line: line},
				kind:           NULL,
			}, nil
		},
		Signature: []DataType{},
	}
}
