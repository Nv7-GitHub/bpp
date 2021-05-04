package parser

// ArrayStmt is the equivalent of [ARRAY stmt.Values[0] stmt.Values[1] etc...]
type ArrayStmt struct {
	Values []Statement

	line int
}

func (a *ArrayStmt) Type() DataType {
	return ARRAY
}

func (a *ArrayStmt) Line() int {
	return a.line
}

// IndexStmt is the equivalent of [VAR stmt.Label stmt.Index]
type IndexStmt struct {
	Value Statement
	Index Statement

	line int
}

func (i *IndexStmt) Type() DataType {
	return ANY
}

func (i *IndexStmt) Line() int {
	return i.line
}

func SetupArrays() {
	parsers["ARRAY"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &ArrayStmt{
				Values: args,
				line:   line,
			}, nil
		},
		Signature: []DataType{ANY, VARIADIC},
	}

	parsers["INDEX"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &IndexStmt{
				Value: args[0],
				Index: args[1],
				line:  line,
			}, nil
		},
		Signature: []DataType{STRING | ARRAY, INT},
	}
}
