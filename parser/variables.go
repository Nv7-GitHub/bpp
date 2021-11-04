package parser

type DefineStmt struct {
	*BasicStmt

	Val      Statement
	Variable string
}

func (d *DefineStmt) Type() Type {
	return NULL
}

func addVariableParsers() {
	parsers["DEFINE"] = Parser{
		Params: []Type{IDENTIFIER, STATEMENT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			var stmt DefineStmt
			stmt.BasicStmt = NewBasicStmt(pos)
			stmt.Variable = params[0].(*Const).Val.(string)
			stmt.Val = params[1]

			typ, exists := prog.VarTypes[stmt.Variable]
			if exists && !params[1].Type().Equal(typ) {
				return nil, pos.NewError("cannot set variable of type %s to value of type %s", typ.String(), params[1].Type().String())
			}
			if !exists {
				prog.VarTypes[stmt.Variable] = params[1].Type()
			}
			return &stmt, nil
		},
	}
}
