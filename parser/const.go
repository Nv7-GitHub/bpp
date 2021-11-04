package parser

import "strconv"

type Const struct {
	*BasicStmt

	Val interface{}
	Typ BasicType
}

func (c *Const) Type() Type {
	return c.Typ
}

func GetConst(text string, pos *Pos) Statement {
	intV, err := strconv.Atoi(text)
	if err == nil {
		return &Const{BasicStmt: NewBasicStmt(pos), Val: intV, Typ: INT}
	}
	floatV, err := strconv.ParseFloat(text, 64)
	if err == nil {
		return &Const{BasicStmt: NewBasicStmt(pos), Val: floatV, Typ: FLOAT}
	}
	if text[0] == '"' && text[len(text)-1] == '"' {
		return &Const{BasicStmt: NewBasicStmt(pos), Val: text[1 : len(text)-1], Typ: STRING}
	}
	return &Const{BasicStmt: NewBasicStmt(pos), Val: text, Typ: STRING}
}
