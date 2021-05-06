package membuild

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/Nv7-Github/Bpp/parser"
)

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
		return NewBlankData(), fmt.Errorf("line %d: parameter to CHOOSE must be STRING or ARRAY", stm.Line())
	}, nil
}

func RepeatStmt(p *Program, stm *parser.RepeatStmt) (Instruction, error) {
	val, err := BuildStmt(p, stm.Val)
	if err != nil {
		return nil, err
	}
	ind, err := BuildStmt(p, stm.Count)
	if err != nil {
		return nil, err
	}
	return func(p *Program) (Data, error) {
		v, err := val(p)
		if err != nil {
			return NewBlankData(), err
		}
		i, err := ind(p)
		if err != nil {
			return NewBlankData(), err
		}

		str, ok := v.Value.(string)
		if ok {
			return Data{
				Type:  parser.STRING,
				Value: strings.Repeat(str, i.Value.(int)),
			}, nil
		}

		arr, ok := v.Value.([]Data)
		if ok {
			repeated := make([]Data, len(arr)*i.Value.(int))
			for j := 0; j < i.Value.(int); j++ {
				for k := 0; k < len(arr); k++ {
					repeated[j*len(arr)+k] = arr[k]
				}
			}
			v.Value = repeated
			return v, nil
		}
		return NewBlankData(), fmt.Errorf("line %d: parameter 1 to REPEAT must be STRING or ARRAY", stm.Line())
	}, nil
}

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
		l, err := getFloat(low.Value, stm.Line(), "RANDOM")
		if err != nil {
			return NewBlankData(), err
		}
		u, err := getFloat(up.Value, stm.Line(), "RANDOM")
		if err != nil {
			return NewBlankData(), err
		}
		return Data{
			Type:  parser.FLOAT,
			Value: l + rand.Float64()*(u-l),
		}, nil
	}, nil
}
