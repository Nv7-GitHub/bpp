package parser

type ConcatStmt struct {
	*BasicStmt

	Values []Statement
}

func (c *ConcatStmt) Type() Type { return STRING }

func addManipStmts() {
	parsers["CONCAT"] = Parser{
		Params: []Type{STRING, VARIADIC},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return &ConcatStmt{
				BasicStmt: NewBasicStmt(pos),

				Values: params,
			}, nil
		},
	}

	// Type casts
	parsers["INT"] = Parser{
		Params: []Type{NewMultiType(FLOAT, STRING)},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return NewTypeCastStmt(pos, params[0], INT), nil
		},
	}
	parsers["FLOAT"] = Parser{
		Params: []Type{NewMultiType(INT, STRING)},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return NewTypeCastStmt(pos, params[0], FLOAT), nil
		},
	}
	parsers["STRING"] = Parser{
		Params: []Type{NewMultiType(INT, FLOAT)},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return NewTypeCastStmt(pos, params[0], STRING), nil
		},
	}
}

type TypeCastStmt struct {
	*BasicStmt

	Val    Statement
	NewTyp Type
}

func (t *TypeCastStmt) Type() Type { return t.NewTyp }

func NewTypeCastStmt(pos *Pos, val Statement, newTyp Type) *TypeCastStmt {
	return &TypeCastStmt{
		BasicStmt: NewBasicStmt(pos),

		Val:    val,
		NewTyp: newTyp,
	}
}
