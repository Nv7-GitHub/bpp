package membuild

import (
	"math/rand"

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
