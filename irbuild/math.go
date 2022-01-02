package irbuild

import "github.com/Nv7-Github/bpp/parser"

func (i *IRBuilder) addMath(s *parser.MathStmt) (int, error) {
	v1, err := i.AddStmt(s.Val1)
	if err != nil {
		return 0, err
	}
	v2, err := i.AddStmt(s.Val2)
	if err != nil {
		return 0, err
	}
	// Cast types
	if !s.Val1.Type().Equal(s.OutTyp) {
		v1 = i.NewCast(v1, s.OutTyp)
	}
	if !s.Val2.Type().Equal(s.OutTyp) {
		v2 = i.NewCast(v2, s.OutTyp)
	}

	// Math instruction
	return i.NewMath(s.Op, v1, v2, s.OutTyp), nil
}
