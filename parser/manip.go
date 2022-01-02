package parser

import "github.com/Nv7-Github/bpp/types"

type ConcatStmt struct {
	*BasicStmt

	Values []Statement
}

func (c *ConcatStmt) Type() types.Type { return types.STRING }

func addManipStmts() {
	parsers["CONCAT"] = Parser{
		Params: []types.Type{types.STRING, types.VARIADIC},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return &ConcatStmt{
				BasicStmt: NewBasicStmt(pos),

				Values: params,
			}, nil
		},
	}

	parsers["INDEX"] = Parser{
		Params: []types.Type{types.NewMultiType(types.STRING, types.ARRAY), types.INT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return &IndexStmt{
				BasicStmt: NewBasicStmt(pos),

				Val:      params[0],
				Index:    params[1],
				IsString: params[0].Type().Equal(types.STRING),
			}, nil
		},
	}

	parsers["LENGTH"] = Parser{
		Params: []types.Type{types.NewMultiType(types.STRING, types.ARRAY), types.INT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return &LengthStmt{
				BasicStmt: NewBasicStmt(pos),

				Val: params[0],
			}, nil
		},
	}

	parsers["CHOOSE"] = Parser{
		Params: []types.Type{types.NewMultiType(types.STRING, types.ARRAY), types.INT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return &ChooseStmt{
				BasicStmt: NewBasicStmt(pos),

				Val: params[0],
			}, nil
		},
	}

	// Type casts
	parsers["INT"] = Parser{
		Params: []types.Type{types.NewMultiType(types.FLOAT, types.STRING)},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return NewTypeCastStmt(pos, params[0], types.INT), nil
		},
	}
	parsers["FLOAT"] = Parser{
		Params: []types.Type{types.NewMultiType(types.INT, types.STRING)},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return NewTypeCastStmt(pos, params[0], types.FLOAT), nil
		},
	}
	parsers["STRING"] = Parser{
		Params: []types.Type{types.NewMultiType(types.INT, types.FLOAT)},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return NewTypeCastStmt(pos, params[0], types.STRING), nil
		},
	}
}

type IndexStmt struct {
	*BasicStmt

	Val      Statement
	Index    Statement
	IsString bool
}

func (i *IndexStmt) Type() types.Type {
	if i.IsString {
		return types.STRING
	}
	return types.ARRAY
}

type LengthStmt struct {
	*BasicStmt

	Val Statement
}

func (l *LengthStmt) Type() types.Type { return types.INT }

type TypeCastStmt struct {
	*BasicStmt

	Val    Statement
	NewTyp types.Type
}

func (t *TypeCastStmt) Type() types.Type { return t.NewTyp }

func NewTypeCastStmt(pos *Pos, val Statement, newTyp types.Type) *TypeCastStmt {
	return &TypeCastStmt{
		BasicStmt: NewBasicStmt(pos),

		Val:    val,
		NewTyp: newTyp,
	}
}

type ChooseStmt struct {
	*BasicStmt

	Val Statement
}

func (c *ChooseStmt) Type() types.Type {
	if c.Val.Type().Equal(types.STRING) {
		return types.STRING
	}
	return c.Val.Type().(*types.Array).ValType
}
