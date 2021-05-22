package parser

// IfBlock is the equivalent of [IFB stmt.Condition] stmt.Body [ELSE] stmt.Else [ENDIF], the stmt.Else may be nil
type IfBlock struct {
	*BasicStatement
	Statements []Statement

	Condition Statement
	Body      []Statement
	Else      []Statement
}

func (i *IfBlock) Type() DataType {
	return NULL
}

func (i *IfBlock) Keywords() []string {
	return []string{"ELSE", "ENDIF"}
}

func (i *IfBlock) EndSignature() []DataType {
	return []DataType{}
}

func (i *IfBlock) End(kind string, _ []Statement, statements []Statement) bool {
	if kind == "ELSE" {
		i.Body = statements
		return false
	}

	if i.Body == nil {
		i.Body = statements
	} else {
		i.Else = statements
	}
	return true
}

func SetupBlocks() {
	blocks["IFB"] = BlockParser{
		Parse: func(args []Statement, line int) (Block, error) {
			return &IfBlock{
				BasicStatement: &BasicStatement{line: line},
				Condition:      args[0],
			}, nil
		},
		Signature: []DataType{INT},
	}

	blocks["WHILE"] = BlockParser{
		Parse: func(args []Statement, line int) (Block, error) {
			return &WhileBlock{
				BasicStatement: &BasicStatement{line: line},
				Condition:      args[0],
			}, nil
		},
		Signature: []DataType{INT},
	}
}

// WhileBlock is the equivalent of [WHILE stmt.Condition]  [ENDWHILE]
type WhileBlock struct {
	*BasicStatement
	Statements []Statement

	Condition Statement
	Body      []Statement
}

func (w *WhileBlock) Type() DataType {
	return NULL
}

func (w *WhileBlock) Keywords() []string {
	return []string{"ENDWHILE"}
}

func (w *WhileBlock) EndSignature() []DataType {
	return []DataType{}
}

func (w *WhileBlock) End(kind string, _ []Statement, statements []Statement) bool {
	w.Body = statements
	return true
}
