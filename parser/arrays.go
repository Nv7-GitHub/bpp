package parser

// ArrayStmt is the equivalent of [ARRAY stmt.Values[0] stmt.Values[1] etc...]
type ArrayStmt struct {
	*BasicStatement
	Values []Statement
}

// Type gets the type of an ARRAY statement (ARRAY)
func (a *ArrayStmt) Type() DataType {
	return ARRAY
}

// IndexStmt is the equivalent of [VAR stmt.Label stmt.Index]
type IndexStmt struct {
	*BasicStatement
	Value Statement
	Index Statement
}

// Type gets the type of an INDEX statement (ANY)
func (i *IndexStmt) Type() DataType {
	return ANY
}

// SetupArrays adds the ARRAY and INDEX functions
func SetupArrays() {
	parsers["ARRAY"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &ArrayStmt{
				Values:         args,
				BasicStatement: &BasicStatement{line: line},
			}, nil
		},
		Signature: []DataType{ANY, VARIADIC},
	}

	parsers["INDEX"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &IndexStmt{
				Value:          args[0],
				Index:          args[1],
				BasicStatement: &BasicStatement{line: line},
			}, nil
		},
		Signature: []DataType{STRING | ARRAY, INT},
	}
}
