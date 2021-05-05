package membuild

import (
	"math"

	"github.com/Nv7-Github/Bpp/parser"
)

func MathStmt(p *Program, stm *parser.MathStmt) (Instruction, error) {
	left, err := BuildStmt(p, stm.Left)
	if err != nil {
		return nil, err
	}
	right, err := BuildStmt(p, stm.Right)
	if err != nil {
		return nil, err
	}
	return func(p *Program) (Data, error) {
		right, err := right(p)
		if err != nil {
			return NewBlankData(), err
		}
		left, err := left(p)
		if err != nil {
			return NewBlankData(), err
		}

		var l float64 = -7
		var r float64 = 7

		if right.Type.IsEqual(parser.FLOAT) || left.Type.IsEqual(parser.FLOAT) {
			l, err = getFloat(left.Value, stm.Line(), "MATH")
			if err != nil {
				return NewBlankData(), err
			}
			r, err = getFloat(right.Value, stm.Line(), "MATH")
			if err != nil {
				return NewBlankData(), err
			}
		} else {
			l = float64(left.Value.(int))
			r = float64(right.Value.(int))
		}

		var out float64 = 0
		switch stm.Operation {
		case parser.ADDITION:
			out = l + r
		case parser.SUBTRACTION:
			out = l - r
		case parser.MULTIPLICATION:
			out = l * r
		case parser.DIVISION:
			out = l / r
		case parser.POWER:
			out = math.Pow(l, r)
		}

		return getDataVal(out), nil
	}, nil
}

func getDataVal(v float64) Data {
	if math.Round(v) == v {
		return Data{
			Type:  parser.INT,
			Value: int(v),
		}
	}
	return Data{
		Type:  parser.FLOAT,
		Value: v,
	}
}
