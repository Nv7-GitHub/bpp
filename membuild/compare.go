package membuild

import (
	"fmt"

	"github.com/Nv7-Github/Bpp/parser"
)

func CompareStmt(p *Program, stm *parser.ComparisonStmt) (Instruction, error) {
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

		var l interface{}
		var r interface{}
		var lnum float64 = -7
		var rnum float64 = 7

		if right.Type.IsEqual(parser.STRING) || left.Type.IsEqual(parser.STRING) {
			l = fmt.Sprintf("%v", right.Value)
			r = fmt.Sprintf("%v", left.Value)
		} else if right.Type.IsEqual(parser.ARRAY) || left.Type.IsEqual(parser.ARRAY) {
			ld, err := convArray(p, left.Value.([]parser.Statement))
			if err != nil {
				return NewBlankData(), err
			}
			rd, err := convArray(p, right.Value.([]parser.Statement))
			if err != nil {
				return NewBlankData(), err
			}
			ldv := make([]interface{}, len(ld))
			rdv := make([]interface{}, len(rd))
			for i, val := range ld {
				ldv[i] = val.Value
			}
			for i, val := range rd {
				rdv[i] = val.Value
			}
			l = ldv
			r = rdv
		} else if right.Type.IsEqual(parser.FLOAT) || left.Type.IsEqual(parser.FLOAT) {
			l, err = getFloat(left.Value, stm.Line())
			if err != nil {
				return NewBlankData(), err
			}
			r, err = getFloat(right.Value, stm.Line())
			if err != nil {
				return NewBlankData(), err
			}
			lnum = l.(float64)
			rnum = r.(float64)
		} else {
			l = left.Value.(int)
			r = right.Value.(int)
			lnum = float64(l.(int))
			rnum = float64(r.(int))
		}

		cond := false
		switch stm.Operation {
		case parser.EQUAL:
			cond = l == r
		case parser.NOTEQUAL:
			cond = l != r
		case parser.GREATER:
			cond = lnum > rnum
		case parser.LESS:
			cond = lnum < rnum
		case parser.GREATEREQUAL:
			cond = lnum >= rnum
		case parser.LESSEQUAL:
			cond = lnum <= rnum
		}

		return getBoolVal(cond), nil
	}, nil
}

func getFloat(val interface{}, line int) (float64, error) {
	v, ok := val.(float64)
	if ok {
		return v, nil
	}
	a, ok := val.(int)
	if ok {
		return float64(a), nil
	}
	return 0, fmt.Errorf("line %d: unknown type in COMPARE", line)
}

func getBoolVal(cond bool) Data {
	out := 0
	if cond {
		out = 1
	}
	return Data{
		Type:  parser.INT,
		Value: out,
	}
}

func convArray(p *Program, arr []parser.Statement) ([]Data, error) {
	out := make([]Data, len(arr))
	for i, val := range arr {
		v, err := BuildStmt(p, val)
		if err != nil {
			return out, err
		}
		d, err := v(p)
		if err != nil {
			return out, err
		}
		out[i] = d
	}
	return out, nil
}
