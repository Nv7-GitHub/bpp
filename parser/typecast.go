package parser

// TypeCastStmt is the equivalent of [stmt.NewType stmt.Val]
type TypeCastStmt struct {
	*BasicStatement
	Value   Statement
	NewType DataType
}

func (t *TypeCastStmt) Type() DataType {
	return t.NewType
}

func SetupTypes() {
	parsers["STRING"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &TypeCastStmt{
				BasicStatement: &BasicStatement{line: line},
				Value:          args[0],
				NewType:        STRING,
			}, nil
		},
		Signature: []DataType{INT | FLOAT},
	}

	parsers["INT"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &TypeCastStmt{
				BasicStatement: &BasicStatement{line: line},
				Value:          args[0],
				NewType:        INT,
			}, nil
		},
		Signature: []DataType{STRING | FLOAT},
	}

	parsers["FLOAT"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &TypeCastStmt{
				BasicStatement: &BasicStatement{line: line},
				Value:          args[0],
				NewType:        FLOAT,
			}, nil
		},
		Signature: []DataType{STRING | INT},
	}

	parsers["LIST"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &TypeCastStmt{
				BasicStatement: &BasicStatement{line: line},
				Value:          args[0],
				NewType:        ARRAY,
			}, nil
		},
		Signature: []DataType{STRING},
	}
}
