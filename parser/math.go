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

	parsers["RANDOM"] = Parser{
		Params: []Type{FLOAT, FLOAT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return &RandomStmt{
				BasicStmt: NewBasicStmt(pos),

				Min: params[0],
				Max: params[1],
			}, nil
		},
	}
	parsers["RANDINT"] = Parser{
		Params: []Type{INT, INT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return &RandintStmt{
				BasicStmt: NewBasicStmt(pos),

				Min: params[0],
				Max: params[1],
			}, nil
		},
	}

	// Math Functions
	parsers["ROUND"] = Parser{
		Params: []Type{FLOAT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return NewMathFunction(pos, params[0], MathFunctionRound, INT), nil
		},
	}
	parsers["CEIL"] = Parser{
		Params: []Type{FLOAT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return NewMathFunction(pos, params[0], MathFunctionCeil, INT), nil
		},
	}
	parsers["FLOOR"] = Parser{
		Params: []Type{FLOAT},
		Parse: func(params []Statement, prog *Program, pos *Pos) (Statement, error) {
			return NewMathFunction(pos, params[0], MathFunctionFloor, INT), nil
		},
	}
}

type RandomStmt struct {
	*BasicStmt

	Min Statement
	Max Statement
}

func (r *RandomStmt) Type() Type { return FLOAT }

type RandintStmt struct {
	*BasicStmt

	Min Statement
	Max Statement
}

func (r *RandintStmt) Type() Type { return INT }

type MathFunction int

const (
	MathFunctionRound MathFunction = iota
	MathFunctionCeil
	MathFunctionFloor
)

type MathFunctionStmt struct {
	*BasicStmt

	Func   MathFunction
	Val    Statement
	OutTyp Type
}

func (m *MathFunctionStmt) Type() Type {
	return m.OutTyp
}

func NewMathFunction(pos *Pos, val Statement, fn MathFunction, outTyp Type) *MathFunctionStmt {
	return &MathFunctionStmt{
		BasicStmt: NewBasicStmt(pos),

		Func:   fn,
		Val:    val,
		OutTyp: outTyp,
	}
}
