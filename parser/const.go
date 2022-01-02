package parser

import (
	"strconv"

	"github.com/Nv7-Github/bpp/types"
)

type Const struct {
	*BasicStmt

	Val interface{}
	Typ types.BasicType
}

func (c *Const) Type() types.Type {
	return c.Typ
}

func GetConst(text string, pos *Pos) Statement {
	intV, err := strconv.Atoi(text)
	if err == nil {
		return &Const{BasicStmt: NewBasicStmt(pos), Val: intV, Typ: types.INT}
	}
	floatV, err := strconv.ParseFloat(text, 64)
	if err == nil {
		return &Const{BasicStmt: NewBasicStmt(pos), Val: floatV, Typ: types.FLOAT}
	}
	if text[0] == '"' && text[len(text)-1] == '"' {
		return &Const{BasicStmt: NewBasicStmt(pos), Val: text[1 : len(text)-1], Typ: types.STRING}
	}
	return &Const{BasicStmt: NewBasicStmt(pos), Val: text, Typ: types.STRING}
}

type ArrayStmt struct {
	*BasicStmt

	Vals []Statement
	Typ  *types.Array
}

func (a *ArrayStmt) Type() types.Type { return a.Typ }

type ArgsStmt struct {
	*BasicStmt

	Index Statement
}

func (a *ArgsStmt) Type() types.Type { return types.STRING }

func addConstStmts() {
	parsers["ARRAY"] = Parser{
		Params: []types.Type{types.STATEMENT, types.VARIADIC},
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
				Typ:       types.NewArrayType(params[0].Type()),
			}, nil
		},
	}

	parsers["ARGS"] = Parser{
		Params: []types.Type{types.INT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return &ArgsStmt{
				BasicStmt: NewBasicStmt(pos),

				Index: params[0],
			}, nil
		},
	}
}
