package parser

// TypeCastStmt is the equivalent of [stmt.NewType stmt.Val]
type TypeCastStmt struct {
	*BasicStatement
	Value   Statement
	NewType DataType
}

// Type gives the return type of a type-cast statement (stmt.NewType)
func (t *TypeCastStmt) Type() DataType {
	return t.NewType
}

// SetupTypes adds the type-cast parsers for STRING, INT, FLOAT, and LIST
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
