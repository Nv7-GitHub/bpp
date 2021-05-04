package parser

// DefineStmt is the equivalent of [DEFINE stmt.Label stmt.Value]
type DefineStmt struct {
	*BasicStatement
	Label Statement
	Value Statement
}

func (d *DefineStmt) Type() DataType {
	return NULL
}

// VarStmt is the equivalent of [VAR stmt.Label]
type VarStmt struct {
	*BasicStatement
	Label Statement
}

func (v *VarStmt) Type() DataType {
	return ANY
}

func SetupVariables() {
	parsers["DEFINE"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &DefineStmt{
				Label:          args[0],
				Value:          args[1],
				BasicStatement: &BasicStatement{line: line},
			}, nil
		},
		Signature: []DataType{IDENTIFIER, ANY},
	}

	parsers["VAR"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &VarStmt{
				Label:          args[0],
				BasicStatement: &BasicStatement{line: line},
			}, nil
		},
		Signature: []DataType{IDENTIFIER},
	}
}
