package parser

type MathOp int

const (
	MathOpAdd MathOp = iota
	MathOpSub
	MathOpMul
	MathOpDiv
	MathOpMod
	MathOpPow
)

var mathOpNames = map[string]MathOp{
	"+": MathOpAdd,
	"-": MathOpSub,
	"*": MathOpMul,
	"/": MathOpDiv,
	"%": MathOpMod,
	"^": MathOpPow,
}

type MathStmt struct {
	*BasicStmt

	Op     MathOp
	Val1   Statement
	Val2   Statement
	OutTyp Type // If parameter type not this, cast to this
}

func (m *MathStmt) Type() Type {
	return m.OutTyp
}

func addMathStmts() {
	parsers["MATH"] = Parser{
		Params: []Type{NewMultiType(INT, FLOAT), STRING, NewMultiType(INT, FLOAT)},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			var outType Type
			t1 := params[0].Type()
			t2 := params[2].Type()
			if t1.Equal(t2) {
				outType = t1
			} else {
				outType = FLOAT // If they arent equal, then one is a FLOAT
			}

			op, ok := params[1].(*Const)
			if !ok {
				return nil, pos.NewError("operation must be constant")
			}
			mathOp, exists := mathOpNames[op.Val.(string)]
			if !exists {
				return nil, pos.NewError("unknown math operation: \"%s\"", op.Val.(string))
			}

			return &MathStmt{
				BasicStmt: NewBasicStmt(pos),

				Op:     mathOp,
				Val1:   params[0],
				Val2:   params[2],
				OutTyp: outType,
			}, nil
		},
	}
}
