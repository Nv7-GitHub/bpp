package parser

// DefineStmt is the equivalent of [DEFINE stmt.Label stmt.Value]
type DefineStmt struct {
	Label Statement
	Value Statement

	line int
}

func (d *DefineStmt) Type() DataType {
	return NULL
}

func (d *DefineStmt) Line() int {
	return d.line
}

// VarStmt is the equivalent of [VAR stmt.Label]
type VarStmt struct {
	Label Statement

	line int
}

func (v *VarStmt) Type() DataType {
	return ANY
}

func (v *VarStmt) Line() int {
	return v.line
}

func SetupVariables() {
	parsers["DEFINE"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &DefineStmt{
				Label: args[0],
				Value: args[1],
				line:  line,
			}, nil
		},
		Signature: []DataType{IDENTIFIER, ANY},
	}

	parsers["VAR"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &VarStmt{
				Label: args[0],
				line:  line,
			}, nil
		},
		Signature: []DataType{IDENTIFIER},
	}
}
