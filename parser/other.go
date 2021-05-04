package parser

// ArgsStmt is the equivalent of [ARGS stmt.Index]
type ArgsStmt struct {
	*BasicStatement
	Index Statement
}

func (a *ArgsStmt) Type() DataType {
	return ANY
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
}
