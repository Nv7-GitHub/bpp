package ir

import (
	"fmt"

	"github.com/Nv7-Github/Bpp/parser"
)

func createConst(val *parser.Data) (*Const, error) {
	switch {
	case val.Type().IsEqual(parser.INT):
		return &Const{
			typ:  INT,
			Data: val.Data,
		}, nil

	case val.Type().IsEqual(parser.FLOAT):
		return &Const{
			typ:  FLOAT,
			Data: val.Data,
		}, nil

	case val.Type().IsEqual(parser.STRING):
		return &Const{
			typ:  STRING,
			Data: val.Data,
		}, nil

	default:
		return nil, fmt.Errorf("%v: unknown constant type", val.Pos())
	}
}

type Const struct {
	typ  Type
	Data interface{}
}

func (c *Const) Type() Type {
	return c.typ
}

func (c *Const) String() string {
	return fmt.Sprintf("Const<%s>: %v", c.Type().String(), c.Data)
}

type castStmt struct {
	val int
	typ Type
}

func (c *castStmt) Type() Type {
	return c.typ
}

func (c *castStmt) String() string {
	return fmt.Sprintf("Cast<%s>: %d", c.Type().String(), c.val)
}

func (i *IR) newCast(val int, typ Type) int {
	return i.AddInstruction(&castStmt{
		val: val,
		typ: typ,
	})
}
