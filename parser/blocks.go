package parser

// IfBlock is the equivalent of [IFB stmt.Condition] stmt.Body [ELSE] stmt.Else [ENDIF], the stmt.Else may be nil
type IfBlock struct {
	*BasicStatement

	Condition Statement
	Body      []Statement
	Else      []Statement
}

// Type gets the type of an IFB statement (NULL)
func (i *IfBlock) Type() DataType {
	return NULL
}

// Keywords gets the keywords of an IFB statement (ELSE and ENDIF)
func (i *IfBlock) Keywords() []string {
	return []string{"ELSE", "ENDIF"}
}

// EndSignature gets the ending signature of an IFB statement (blank)
func (i *IfBlock) EndSignature() []DataType {
	return []DataType{}
}

// End parses the end of an IFB statement
func (i *IfBlock) End(kind string, _ []Statement, statements []Statement) (bool, error) {
	if kind == "ELSE" {
		i.Body = statements
		return false, nil
	}

	if i.Body == nil {
		i.Body = statements
	} else {
		i.Else = statements
	}
	return true, nil
}

// SetupBlocks adds the IFB and WHILE functions
func SetupBlocks() {
	blocks["IFB"] = BlockParser{
		Parse: func(args []Statement, pos *Pos) (Block, error) {
			return &IfBlock{
				BasicStatement: &BasicStatement{pos: pos},
				Condition:      args[0],
			}, nil
		},
		Signature: []DataType{INT},
	}

	blocks["WHILE"] = BlockParser{
		Parse: func(args []Statement, pos *Pos) (Block, error) {
			return &WhileBlock{
				BasicStatement: &BasicStatement{pos: pos},
				Condition:      args[0],
			}, nil
		},
		Signature: []DataType{INT},
	}
}

// WhileBlock is the equivalent of [WHILE stmt.Condition]  [ENDWHILE]
type WhileBlock struct {
	*BasicStatement

	Condition Statement
	Body      []Statement
}

// Type gives the type of a WHILE statement (nothing)
func (w *WhileBlock) Type() DataType {
	return NULL
}

// Keywords give the keywords of a WHILE statement (just ENDWHILE)
func (w *WhileBlock) Keywords() []string {
	return []string{"ENDWHILE"}
}

// EndSignature gets the end signature of a WHILE statement (blank)
func (w *WhileBlock) EndSignature() []DataType {
	return []DataType{}
}

// End parses the ending of a WHILE statement
func (w *WhileBlock) End(_ string, _ []Statement, statements []Statement) (bool, error) {
	w.Body = statements
	return true, nil
}
