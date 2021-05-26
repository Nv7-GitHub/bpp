package membuild

import (
	"fmt"
	"strconv"

	"github.com/Nv7-Github/Bpp/parser"
)

func TypeCastStmt(p *Program, stm *parser.TypeCastStmt) (Instruction, error) {
	val, err := BuildStmt(p, stm.Value)
	if err != nil {
		return nil, err
	}

	return func(p *Program) (Data, error) {
		val, err := val(p)
		if err != nil {
			return NewBlankData(), err
		}

		switch stm.NewType {
		case parser.INT:
			return convertInt(val, stm.Line())

		case parser.FLOAT:
			return convertFloat(val, stm.Line())

		case parser.STRING:
			return Data{
				Type:  parser.STRING,
				Value: fmt.Sprintf("%v", val.Value),
			}, nil
		}
		return NewBlankData(), nil
	}, nil
}

func convertInt(val Data, num int) (Data, error) {
	switch {
	case val.Type.IsEqual(parser.STRING):
		val, err := strconv.Atoi(val.Value.(string))
		if err != nil {
			return NewBlankData(), fmt.Errorf("line %d: parameter to INT is not integer", num)
		}
		return Data{
			Type:  parser.INT,
			Value: val,
		}, nil

	case val.Type.IsEqual(parser.FLOAT):
		return Data{
			Type:  parser.INT,
			Value: int(val.Value.(float64)),
		}, nil
	}
	return NewBlankData(), fmt.Errorf("line %d: cannot convert type to integer", num)
}

func convertFloat(val Data, num int) (Data, error) {
	switch {
	case val.Type.IsEqual(parser.STRING):
		val, err := strconv.ParseFloat(val.Value.(string), 64)
		if err != nil {
			return NewBlankData(), fmt.Errorf("line %d: parameter to FLOAT is not a float", num)
		}
		return Data{
			Type:  parser.FLOAT,
			Value: val,
		}, nil

	case val.Type.IsEqual(parser.INT):
		return Data{
			Type:  parser.FLOAT,
			Value: float64(val.Value.(int)),
		}, nil
	}

	return NewBlankData(), fmt.Errorf("line %d: cannot convert type to float", num)
}
