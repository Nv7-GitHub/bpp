package membuild

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/Nv7-Github/Bpp/parser"
)

func RandintStmt(p *Program, stm *parser.RandintStmt) (Instruction, error) {
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
		return Data{
			Type:  parser.INT,
			Value: rand.Intn(up.Value.(int)-low.Value.(int)) + low.Value.(int),
		}, nil
	}, nil
}

func ConcatStmt(p *Program, stm *parser.ConcatStmt) (Instruction, error) {
	argDat := make([]Instruction, len(stm.Strings))
	var err error
	for i, str := range stm.Strings {
		argDat[i], err = BuildStmt(p, str)
		if err != nil {
			return nil, err
		}
	}
	return func(p *Program) (Data, error) {
		args := make([]string, len(argDat))
		for i, arg := range argDat {
			v, err := arg(p)
			if err != nil {
				return NewBlankData(), err
			}
			args[i] = fmt.Sprintf("%v", v.Value)
		}
		return Data{
			Type:  parser.STRING,
			Value: strings.Join(args, ""),
		}, nil
	}, nil
}

func IndexStmt(p *Program, stm *parser.IndexStmt) (Instruction, error) {
	val, err := BuildStmt(p, stm.Value)
	if err != nil {
		return nil, err
	}
	ind, err := BuildStmt(p, stm.Index)
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
				Value: string(str[i.Value.(int)]),
			}, nil
		}

		arr, ok := v.Value.([]Data)
		if ok {
			return arr[i.Value.(int)], nil
		}
		return NewBlankData(), fmt.Errorf("line %d: parameter to INDEX must be STRING or ARRAY", stm.Line())
	}, nil
}

func ArgsStmt(p *Program, stm *parser.ArgsStmt) (Instruction, error) {
	ind, err := BuildStmt(p, stm.Index)
	if err != nil {
		return nil, err
	}
	return func(p *Program) (Data, error) {
		i, err := ind(p)
		if err != nil {
			return NewBlankData(), err
		}

		if i.Value.(int) >= len(p.Args) {
			return NewBlankData(), fmt.Errorf("line %d: argument index out of bounds", stm.Line())
		}
		return Data{
			Type:  parser.STRING,
			Value: p.Args[i.Value.(int)],
		}, nil
	}, nil
}

func ArrayStmt(p *Program, stm *parser.ArrayStmt) (Instruction, error) {
	vals := make([]Instruction, len(stm.Values))
	var err error
	for i, val := range stm.Values {
		vals[i], err = BuildStmt(p, val)
		if err != nil {
			return nil, err
		}
	}
	return func(p *Program) (Data, error) {
		vs := make([]Data, len(vals))
		for i, instruction := range vals {
			vs[i], err = instruction(p)
			if err != nil {
				return NewBlankData(), err
			}
		}
		return Data{
			Type:  parser.ARRAY,
			Value: vs,
		}, nil
	}, nil
}
