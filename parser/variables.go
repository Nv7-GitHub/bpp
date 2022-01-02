package parser

import "github.com/Nv7-Github/bpp/types"

type DefineStmt struct {
	*BasicStmt

	Val      Statement
	Variable string
}

func (d *DefineStmt) Type() types.Type {
	return types.NULL
}

type VarStmt struct {
	*BasicStmt

	Variable string
	Typ      types.Type
}

func (v *VarStmt) Type() types.Type {
	return v.Typ
}

func addVariableParsers() {
	parsers["DEFINE"] = Parser{
		Params: []types.Type{types.STRING, types.STATEMENT},
		Parse: func(params []Statement, prog *Program, pos *types.Pos) (Statement, error) {
			var stmt DefineStmt
			stmt.BasicStmt = NewBasicStmt(pos)
			_, ok := params[0].(*Const)
			if !ok {
				return nil, pos.NewError("variable names must be constants")
			}
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

	parsers["VAR"] = Parser{
		Params: []types.Type{types.STRING},
		Parse: func(params []Statement, prog *Program, pos *types.Pos) (Statement, error) {
			name, ok := params[0].(*Const)
			if !ok {
				return nil, pos.NewError("variable names must be constants")
			}

			v, exists := prog.VarTypes[name.Val.(string)]
			if !exists {
				return nil, pos.NewError("variable %s not defined", name.Val.(string))
			}

			return &VarStmt{
				BasicStmt: NewBasicStmt(pos),

				Variable: name.Val.(string),
				Typ:      v,
			}, nil
		},
	}
}
