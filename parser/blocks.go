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

	parsers["BLOCK"] = Parser{
		Params: []Type{STATEMENT, VARIADIC},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return &BlockStmt{
				BasicStmt: NewBasicStmt(pos),

				Body: params,
			}, nil
		},
	}

	parsers["IF"] = Parser{
		Params: []Type{INT, STATEMENT, STATEMENT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			var par1blk []Statement
			var par2blk []Statement
			// Check if par1 is block
			_, ok := params[1].(*BlockStmt)
			if ok {
				par1blk = params[1].(*BlockStmt).Body
				_, ok = params[2].(*BlockStmt)
				if !ok {
					par2blk = []Statement{params[2]}
				} else {
					par2blk = params[2].(*BlockStmt).Body
				}
			}
			// Check if par2 is block
			if par1blk == nil {
				_, ok = params[2].(*BlockStmt)
				if ok {
					par2blk = params[2].(*BlockStmt).Body
					_, ok = params[1].(*BlockStmt)
					if !ok {
						par1blk = []Statement{params[1]}
					} else {
						par1blk = params[1].(*BlockStmt).Body
					}
				}
			}

			// if block, then not ternary
			if par1blk != nil && par2blk != nil {
				return &IfStmt{
					BasicStmt: NewBasicStmt(pos),

					Condition: params[0],
					Body:      par1blk,
					Else:      par2blk,
				}, nil
			}

			// Otherwise its ternary
			commonTyp, err := getCommonType(params[1].Type(), params[2].Type(), pos)
			if err != nil {
				return nil, err
			}
			return &IfTernary{
				BasicStmt: NewBasicStmt(pos),

				Condition: params[0],
				Body:      params[1],
				Else:      params[2],
				OutTyp:    commonTyp,
			}, nil
		},
	}
}

type BlockStmt struct {
	*BasicStmt

	Body []Statement
}

func (b *BlockStmt) Type() Type {
	return NULL
}

type IfStmt struct {
	*BasicStmt

	Condition Statement
	Body      []Statement
	Else      []Statement
}

func (i *IfStmt) Type() Type {
	return NULL
}

type IfTernary struct {
	*BasicStmt

	Condition Statement
	Body      Statement
	Else      Statement
	OutTyp    Type
}

func (i *IfTernary) Type() Type {
	return i.OutTyp
}
