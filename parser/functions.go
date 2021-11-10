package parser

import (
	"regexp"
	"strings"
)

func (p *Program) BeginFunction(name string, pos *Pos) error {
	if p.InFunction {
		return pos.NewError("nested functions aren't allowed")
	}

	p.InFunction = true
	p.FuncName = name
	p.OldVarTypes = p.VarTypes
	p.VarTypes = make(map[string]Type)
	p.Functions[name] = &Function{
		Name:   name,
		Params: make([]FunctionParam, 0),
	}

	return nil
}

var paramTypeNames = map[string]BasicType{
	"INT":    INT,
	"FLOAT":  FLOAT,
	"STRING": STRING,
}

var arrayTypeNameParser = regexp.MustCompile(`^ARRAY\{(.+)\}$`)

type ParamStmt struct{ *BasicStmt }

func (p *ParamStmt) Type() Type { return NULL }

type RetTypeStmt struct{ *BasicStmt }

func (r *RetTypeStmt) Type() Type { return NULL }

type ReturnStmt struct {
	*BasicStmt

	Val Statement
}

func (r *ReturnStmt) Type() Type { return r.Val.Type() }

func addFunctionStmts() {
	parsers["PARAM"] = Parser{
		Params: []Type{STRING, STRING},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			// Check if in param section
			if !prog.InFunction || prog.Functions[prog.FuncName].RetType != nil {
				return nil, pos.NewError("misplaced PARAM statement")
			}

			nameval, ok := params[0].(*Const)
			if !ok {
				return nil, pos.NewError("name of parameter must be constant")
			}
			name := nameval.Val.(string)

			// Get type
			typeval, ok := params[1].(*Const)
			if !ok {
				return nil, pos.NewError("type of parameter must be constant")
			}
			typName := typeval.Val.(string)
			typ, err := ParseTypeString(typName, pos)
			if err != nil {
				return nil, err
			}

			// Save param
			par := FunctionParam{Name: name, Type: typ}
			prog.Functions[prog.FuncName].Params = append(prog.Functions[prog.FuncName].Params, par)

			// Add to var types
			prog.VarTypes[name] = typ

			return &ParamStmt{NewBasicStmt(pos)}, nil
		},
	}

	parsers["RETURNS"] = Parser{
		Params: []Type{STRING},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			// Check if in param section
			if !prog.InFunction || prog.Functions[prog.FuncName].RetType != nil {
				return nil, pos.NewError("misplaced RETURNS statement")
			}

			typName, ok := params[0].(*Const)
			if !ok {
				return nil, pos.NewError("type of return value must be constant")
			}
			typ, err := ParseTypeString(typName.Val.(string), pos)
			if err != nil {
				return nil, err
			}

			prog.Functions[prog.FuncName].RetType = typ

			return &RetTypeStmt{NewBasicStmt(pos)}, nil
		},
	}

	parsers["FUNCTION"] = Parser{
		Params: []Type{STRING, NULL, VARIADIC, STATEMENT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			// Get rid of name
			params = params[1:]

			blkV := params[len(params)-1]
			blk, ok := blkV.(*BlockStmt)
			if !ok {
				return nil, pos.NewError("last argument to function must be BLOCK")
			}

			// Verify type of NULL params
			parTyps := params[:len(params)-1]
			_, ok = parTyps[len(parTyps)-1].(*RetTypeStmt)
			if !ok {
				return nil, pos.NewError("last parameter type to FUNCTION must be RETURNS")
			}
			parTyps = parTyps[:len(parTyps)-1]
			for _, parTyp := range parTyps {
				_, ok = parTyp.(*ParamStmt)
				if !ok {
					return nil, pos.NewError("parameter types to FUNCTION must be PARAM")
				}
			}

			// Save function
			prog.Functions[prog.FuncName].Statements = blk.Body
			prog.VarTypes = prog.OldVarTypes

			return nil, nil
		},
	}

	parsers["RETURN"] = Parser{
		Params: []Type{STATEMENT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			// Check if in right section
			if !prog.InFunction {
				return nil, pos.NewError("RETURN must be in function")
			}
			fnRetType := prog.Functions[prog.FuncName].RetType
			if fnRetType == nil {
				return nil, pos.NewError("RETURN must be in BLOCK")
			}

			if !params[0].Type().Equal(fnRetType) {
				return nil, pos.NewError("expected return type of \"%s\", got \"%s\"", fnRetType.String(), params[0].Type().String())
			}

			return &ReturnStmt{
				BasicStmt: NewBasicStmt(pos),

				Val: params[0],
			}, nil
		},
	}
}

func ParseTypeString(typName string, pos *Pos) (Type, error) {
	// Check if basic type
	var typ Type
	var ok bool
	typ, ok = paramTypeNames[typName]
	if !ok {
		arrTyp := strings.HasPrefix(typName, "ARRAY")
		if arrTyp {
			valtyp := arrayTypeNameParser.FindAllStringSubmatch(typName, 1)
			if len(valtyp) != 1 {
				return nil, pos.NewError("invalid array type")
			}
			// Recursively parse type
			var err error
			typ, err = ParseTypeString(valtyp[0][0], pos)
			if err != nil {
				return nil, err
			}
			typ = &Array{
				ValType: typ,
			}
		} else {
			return nil, pos.NewError("unknown type \"%s\"", typName)
		}
	}

	return typ, nil
}
