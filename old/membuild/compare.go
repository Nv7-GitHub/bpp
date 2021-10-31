package membuild

import (
	"fmt"
	"reflect"

	"github.com/Nv7-Github/bpp/old/parser"
)

// CompareStmt compiles a COMPARE statement
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
			ld := left.Value.([]Data)
			rd := right.Value.([]Data)
			ldv := make([]interface{}, len(ld))
			rdv := make([]interface{}, len(rd))
			for i, val := range ld {
				ldv[i] = val.Value
			}
			for i, val := range rd {
				rdv[i] = val.Value
			}
			eq := true
			for i, val := range ld {
				if rd[i] != val {
					eq = false
				}
			}
			if stm.Operation == parser.EQUAL {
				return getBoolVal(eq), nil
			}
			return getBoolVal(!eq), nil
		} else if right.Type.IsEqual(parser.FLOAT) || left.Type.IsEqual(parser.FLOAT) {
			l, err = getFloat(left.Value, stm.Pos(), "COMPARE")
			if err != nil {
				return NewBlankData(), err
			}
			r, err = getFloat(right.Value, stm.Pos(), "COMPARE")
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

func getFloat(val interface{}, pos *parser.Pos, funcName string) (float64, error) {
	v, ok := val.(float64)
	if ok {
		return v, nil
	}
	a, ok := val.(int)
	if ok {
		return float64(a), nil
	}
	return 0, fmt.Errorf("%v: unknown type %s in %s", pos, reflect.TypeOf(val), funcName)
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
