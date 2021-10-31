package parser

import "strconv"

type Const struct {
	Val interface{}
	Typ BasicType
}

func (c *Const) Type() Type {
	return c.Typ
}

func GetConst(text string) Statement {
	intV, err := strconv.Atoi(text)
	if err == nil {
		return &Const{Val: intV, Typ: INT}
	}
	floatV, err := strconv.ParseFloat(text, 64)
	if err == nil {
		return &Const{Val: floatV, Typ: FLOAT}
	}
	if text[0] == '"' && text[len(text)-1] == '"' {
		return &Const{Val: text[1 : len(text)-1], Typ: STRING}
	}
	return &Const{Val: text, Typ: IDENTIFIER}
}
