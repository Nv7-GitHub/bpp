package parser

// DefineStmt is the equivalent of [DEFINE stmt.Label stmt.Value]
type DefineStmt struct {
	*BasicStatement
	Label Statement
	Value Statement
}

// Type gives the return type of a DEFINE statement (NULL)
func (d *DefineStmt) Type() DataType {
	return NULL
}

// VarStmt is the equivalent of [VAR stmt.Label]
type VarStmt struct {
	*BasicStatement
	Label Statement
}

// Type gives the return type of a VAR statement (ANY)
func (v *VarStmt) Type() DataType {
	return ANY
}

// SetupVariables adds the DEFINE and VAR statements
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
