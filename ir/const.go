package ir

import (
	"fmt"

	"github.com/Nv7-Github/bpp/types"
)

type Const struct {
	Data interface{}
	Typ  types.Type
}

func (c *Const) Type() types.Type {
	return c.Typ
}

func (c *Const) String() string {
	return fmt.Sprintf("Const<%s>: %v", c.Type().String(), c.Data)
}

func (i *IR) AddConst(value interface{}, typ types.Type) int {
	return i.AddInstruction(&Const{
		Data: value,
		Typ:  typ,
	})
}

type Cast struct {
	Val    int
	NewTyp types.Type
}

func (c *Cast) Type() types.Type {
	return c.NewTyp
}

func (c *Cast) String() string {
	return fmt.Sprintf("Cast<%s>: %d", c.Type().String(), c.Val)
}

func (i *IR) NewCast(val int, newTyp types.Type) int {
	// TODO: Array => String cast
	return i.AddInstruction(&Cast{
		Val:    val,
		NewTyp: newTyp,
	})
}
