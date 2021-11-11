package parser

import "strconv"

type Const struct {
	*BasicStmt

	Val interface{}
	Typ BasicType
}

func (c *Const) Type() Type {
	return c.Typ
}

func GetConst(text string, pos *Pos) Statement {
	intV, err := strconv.Atoi(text)
	if err == nil {
		return &Const{BasicStmt: NewBasicStmt(pos), Val: intV, Typ: INT}
	}
	floatV, err := strconv.ParseFloat(text, 64)
	if err == nil {
		return &Const{BasicStmt: NewBasicStmt(pos), Val: floatV, Typ: FLOAT}
	}
	if text[0] == '"' && text[len(text)-1] == '"' {
		return &Const{BasicStmt: NewBasicStmt(pos), Val: text[1 : len(text)-1], Typ: STRING}
	}
	return &Const{BasicStmt: NewBasicStmt(pos), Val: text, Typ: STRING}
}

type ArrayStmt struct {
	*BasicStmt

	Vals []Statement
	Typ  *Array
}

func (a *ArrayStmt) Type() Type { return a.Typ }

type ArgsStmt struct {
	*BasicStmt

	Index Statement
}

func (a *ArgsStmt) Type() Type { return STRING }

func addConstStmts() {
	parsers["ARRAY"] = Parser{
		Params: []Type{STATEMENT, VARIADIC},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			// Type check
			if len(params) > 1 {
				firstTyp := params[0].Type()
				for i, par := range params[1:] {
					if par.Type() != firstTyp {
						return nil, pos.NewError("element %d of array does not match type %s", i+1, firstTyp.String())
					}
				}
			}

			return &ArrayStmt{
				BasicStmt: NewBasicStmt(pos),
				Vals:      params,
				Typ:       NewArrayType(params[0].Type()),
			}, nil
		},
	}

	parsers["ARGS"] = Parser{
		Params: []Type{INT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return &ArgsStmt{
				BasicStmt: NewBasicStmt(pos),

				Index: params[0],
			}, nil
		},
	}
}
