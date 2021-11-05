package parser

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
	MutualTyp Type
}

func (c *CompareStmt) Type() Type {
	return INT
}

func addConditionals() {
	parsers["COMPARE"] = Parser{
		Params: []Type{STATEMENT, STRING, STATEMENT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			opV, ok := params[1].(*Const)
			if !ok {
				return nil, pos.NewError("compare operation must be constant")
			}
			op, exists := opNames[opV.Val.(string)]
			if !exists {
				return nil, pos.NewError("unknown compare operation \"%s\"", opV.Val.(string))
			}

			var outTyp Type
			typ1 := params[0].Type()
			typ2 := params[2].Type()

			// Numbers
			if typ1.Equal(INT) || typ1.Equal(FLOAT) {
				if typ2.Equal(INT) || typ2.Equal(FLOAT) {
					if typ1.Equal(typ2) {
						outTyp = typ1
					} else {
						outTyp = FLOAT // if not equal, one is a float
					}
				} else { // if one is a number and the other isn't
					return nil, pos.NewError("can only compare numbers")
				}
			} else if typ1.Equal(STRING) {
				if typ2.Equal(STRING) {
					outTyp = STRING
				} else {
					return nil, pos.NewError("can only compare strings to strings")
				}
			} else if typ1.Equal(ARRAY) {
				if typ2.Equal(typ1) {
					outTyp = typ1
				} else {
					return nil, pos.NewError("can only compare arrays to arrays")
				}
			} else {
				return nil, pos.NewError("unknown type comparison")
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
