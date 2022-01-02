package parser

import "github.com/Nv7-Github/bpp/types"

type CompareOp int

const (
	CompareOpEqual CompareOp = iota
	CompareOpNotEqual
	CompareOpLess
	CompareOpLessOrEqual
	CompareOpGreater
	CompareOpGreaterOrEqual
)

var opNames = map[string]CompareOp{
	"=":  CompareOpEqual,
	"!=": CompareOpNotEqual,
	"<":  CompareOpLess,
	"<=": CompareOpLessOrEqual,
	">":  CompareOpGreater,
	">=": CompareOpGreaterOrEqual,
}

type CompareStmt struct {
	*BasicStmt

	Op        CompareOp
	Val1      Statement
	Val2      Statement
	MutualTyp types.Type
}

func (c *CompareStmt) Type() types.Type {
	return types.INT
}

func getCommonType(typ1, typ2 types.Type, pos *Pos) (types.Type, error) {
	var outTyp types.Type
	if typ1.Equal(types.INT) || typ1.Equal(types.FLOAT) {
		if typ2.Equal(types.INT) || typ2.Equal(types.FLOAT) {
			if typ1.Equal(typ2) {
				outTyp = typ1
			} else {
				outTyp = types.FLOAT // if not equal, one is a float
			}
		} else { // if one is a number and the other isn't
			return nil, pos.NewError("can only compare numbers")
		}
	} else if typ1.Equal(types.STRING) {
		if typ2.Equal(types.STRING) {
			outTyp = types.STRING
		} else {
			return nil, pos.NewError("can only compare strings to strings")
		}
	} else if typ1.Equal(types.ARRAY) {
		if typ2.Equal(typ1) {
			outTyp = typ1
		} else {
			return nil, pos.NewError("can only compare arrays to arrays")
		}
	} else {
		return nil, pos.NewError("unknown type comparison")
	}
	return outTyp, nil
}

func addConditionals() {
	parsers["COMPARE"] = Parser{
		Params: []types.Type{types.STATEMENT, types.STRING, types.STATEMENT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			opV, ok := params[1].(*Const)
			if !ok {
				return nil, pos.NewError("compare operation must be constant")
			}
			op, exists := opNames[opV.Val.(string)]
			if !exists {
				return nil, pos.NewError("unknown compare operation \"%s\"", opV.Val.(string))
			}

			var outTyp types.Type
			typ1 := params[0].Type()
			typ2 := params[2].Type()
			outTyp, err := getCommonType(typ1, typ2, pos)
			if err != nil {
				return nil, err
			}

			return &CompareStmt{
				BasicStmt: NewBasicStmt(pos),

				Op:        op,
				Val1:      params[0],
				Val2:      params[2],
				MutualTyp: outTyp,
			}, nil
		},
	}
}
