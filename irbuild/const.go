package irbuild

import "github.com/Nv7-Github/bpp/parser"

func (i *IRBuilder) addConst(s *parser.Const) (int, error) {
	return i.AddConst(s.Val, s.Type()), nil
}
