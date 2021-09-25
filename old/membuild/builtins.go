package membuild

import (
	"fmt"
	"math/rand"

	"github.com/Nv7-Github/bpp/parser"
)

// ChooseStmt compiles a CHOOSE statement
func ChooseStmt(p *Program, stm *parser.ChooseStmt) (Instruction, error) {
	val, err := BuildStmt(p, stm.Data)
	if err != nil {
		return nil, err
	}
	return func(p *Program) (Data, error) {
		v, err := val(p)
		if err != nil {
			return NewBlankData(), err
		}

		str, ok := v.Value.(string)
		if ok {
			return Data{
				Type:  parser.STRING,
				Value: string(str[rand.Intn(len(str))]),
			}, nil
		}

		arr, ok := v.Value.([]Data)
		if ok {
			return arr[rand.Intn(len(arr))], nil
		}
		return NewBlankData(), fmt.Errorf("%v: parameter to CHOOSE must be STRING or ARRAY", stm.Pos())
	}, nil
}

// RandomStmt compiles a RANDOM statement
func RandomStmt(p *Program, stm *parser.RandomStmt) (Instruction, error) {
	lower, err := BuildStmt(p, stm.Lower)
	if err != nil {
		return nil, err
	}
	upper, err := BuildStmt(p, stm.Upper)
	if err != nil {
		return nil, err
	}
	return func(p *Program) (Data, error) {
		low, err := lower(p)
		if err != nil {
			return NewBlankData(), err
		}
		up, err := upper(p)
		if err != nil {
			return NewBlankData(), err
		}
		l, err := getFloat(low.Value, stm.Pos(), "RANDOM")
		if err != nil {
			return NewBlankData(), err
		}
		u, err := getFloat(up.Value, stm.Pos(), "RANDOM")
		if err != nil {
			return NewBlankData(), err
		}
		return Data{
			Type:  parser.FLOAT,
			Value: l + rand.Float64()*(u-l),
		}, nil
	}, nil
}
