package parser

import "github.com/Nv7-Github/bpp/types"

var mathOpNames = map[string]types.MathOp{
	"+": types.MathOpAdd,
	"-": types.MathOpSub,
	"*": types.MathOpMul,
	"/": types.MathOpDiv,
	"%": types.MathOpMod,
	"^": types.MathOpPow,
}

type MathStmt struct {
	*BasicStmt

	Op     types.MathOp
	Val1   Statement
	Val2   Statement
	OutTyp types.Type // If parameter type not this, cast to this
}

func (m *MathStmt) Type() types.Type {
	return m.OutTyp
}

func addMathStmts() {
	parsers["MATH"] = Parser{
		Params: []types.Type{types.NewMultiType(types.INT, types.FLOAT), types.STRING, types.NewMultiType(types.INT, types.FLOAT)},
		Parse: func(params []Statement, prog *Program, pos *types.Pos) (Statement, error) {
			var outType types.Type
			t1 := params[0].Type()
			t2 := params[2].Type()
			if t1.Equal(t2) {
				outType = t1
			} else {
				outType = types.FLOAT // If they arent equal, then one is a FLOAT
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
		Params: []types.Type{types.FLOAT, types.FLOAT},
		Parse: func(params []Statement, prog *Program, pos *types.Pos) (Statement, error) {
			return &RandomStmt{
				BasicStmt: NewBasicStmt(pos),

				Min: params[0],
				Max: params[1],
			}, nil
		},
	}
	parsers["RANDINT"] = Parser{
		Params: []types.Type{types.INT, types.INT},
		Parse: func(params []Statement, prog *Program, pos *types.Pos) (Statement, error) {
			return &RandintStmt{
				BasicStmt: NewBasicStmt(pos),

				Min: params[0],
				Max: params[1],
			}, nil
		},
	}

	// Math Functions
	parsers["ROUND"] = Parser{
		Params: []types.Type{types.FLOAT},
		Parse: func(params []Statement, prog *Program, pos *types.Pos) (Statement, error) {
			return NewMathFunction(pos, params[0], MathFunctionRound, types.INT), nil
		},
	}
	parsers["CEIL"] = Parser{
		Params: []types.Type{types.FLOAT},
		Parse: func(params []Statement, prog *Program, pos *types.Pos) (Statement, error) {
			return NewMathFunction(pos, params[0], MathFunctionCeil, types.INT), nil
		},
	}
	parsers["FLOOR"] = Parser{
		Params: []types.Type{types.FLOAT},
		Parse: func(params []Statement, prog *Program, pos *types.Pos) (Statement, error) {
			return NewMathFunction(pos, params[0], MathFunctionFloor, types.INT), nil
		},
	}
}

type RandomStmt struct {
	*BasicStmt

	Min Statement
	Max Statement
}

func (r *RandomStmt) Type() types.Type { return types.FLOAT }

type RandintStmt struct {
	*BasicStmt

	Min Statement
	Max Statement
}

func (r *RandintStmt) Type() types.Type { return types.INT }

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
	OutTyp types.Type
}

func (m *MathFunctionStmt) Type() types.Type {
	return m.OutTyp
}

func NewMathFunction(pos *types.Pos, val Statement, fn MathFunction, outTyp types.Type) *MathFunctionStmt {
	return &MathFunctionStmt{
		BasicStmt: NewBasicStmt(pos),

		Func:   fn,
		Val:    val,
		OutTyp: outTyp,
	}
}
