package parser

// ArgsStmt is the equivalent of [ARGS stmt.Index]
type ArgsStmt struct {
	*BasicStatement
	Index Statement
}

// Type returns the return type of an ARGS statement (STRING)
func (a *ArgsStmt) Type() DataType {
	return STRING
}

// ConcatStmt is the equivalent of [CONCAT stmt.Strings]
type ConcatStmt struct {
	*BasicStatement
	Strings []Statement
}

// Type returns the return type of a CONCAT statement (STRING)
func (c *ConcatStmt) Type() DataType {
	return STRING
}

// RepeatStmt is the equivalent of [REPEAT stmt.Val]
type RepeatStmt struct {
	*BasicStatement
	Val   Statement
	Count Statement
}

// Type gives the return type of a REPEAT statement (STRING)
func (r *RepeatStmt) Type() DataType {
	return STRING
}

// ImportStmt is the AST tree resulted of [IMPORT stmt.Filename]
type ImportStmt struct {
	*BasicStatement
	Statements []Statement
	Filename   string
}

// Type returns the return type of a IMPORT statement (NULL)
func (i *ImportStmt) Type() DataType {
	return NULL
}

// SetupOthers adds the ARGS, CONCAT, and REPEAT statements
func SetupOthers() {
	parsers["ARGS"] = StatementParser{
		Parse: func(args []Statement, pos *Pos) (Statement, error) {
			return &ArgsStmt{
				Index:          args[0],
				BasicStatement: &BasicStatement{pos: pos},
			}, nil
		},
		Signature: []DataType{INT},
	}

	parsers["CONCAT"] = StatementParser{
		Parse: func(args []Statement, pos *Pos) (Statement, error) {
			return &ConcatStmt{
				Strings:        args,
				BasicStatement: &BasicStatement{pos: pos},
			}, nil
		},
		Signature: []DataType{ANY, VARIADIC},
	}

	parsers["REPEAT"] = StatementParser{
		Parse: func(args []Statement, pos *Pos) (Statement, error) {
			return &RepeatStmt{
				Val:            args[0],
				Count:          args[1],
				BasicStatement: &BasicStatement{pos: pos},
			}, nil
		},
		Signature: []DataType{STRING, INT},
	}

	parsers["NULL"] = StatementParser{
		Parse: func(args []Statement, pos *Pos) (Statement, error) {
			return &Data{
				BasicStatement: &BasicStatement{pos: pos},
				kind:           NULL,
			}, nil
		},
		Signature: []DataType{},
	}
}
