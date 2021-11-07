package parser

type RepeatStmt struct {
	*BasicStmt

	Count Statement
	Body  []Statement
}

func (r *RepeatStmt) Type() Type {
	return NULL
}

type WhileStmt struct {
	*BasicStmt

	Condition Statement
	Body      []Statement
}

func (w *WhileStmt) Type() Type {
	return NULL
}

func addLoops() {
	parsers["REPEAT"] = Parser{
		Params: []Type{INT, STATEMENT, VARIADIC},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return &RepeatStmt{
				BasicStmt: NewBasicStmt(pos),

				Count: params[0],
				Body:  params[1:],
			}, nil
		},
	}

	parsers["WHILE"] = Parser{
		Params: []Type{INT, STATEMENT, VARIADIC},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return &WhileStmt{
				BasicStmt: NewBasicStmt(pos),

				Condition: params[0],
				Body:      params[1:],
			}, nil
		},
	}
}
