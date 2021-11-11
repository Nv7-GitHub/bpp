package run

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Nv7-Github/bpp/old/ir"
)

func (r *Runnable) runConst() {
	r.registers[r.Index] = r.ir.Instructions[r.Index].(*ir.Const).Data
}

func (r *Runnable) runPrint(print *ir.Print) error {
	_, err := r.Stdout.Write([]byte(r.registers[print.Val].(string) + "\n"))
	return err
}

func (r *Runnable) runCast(i *ir.Cast) error {
	typ := r.ir.Instructions[i.Val].Type()
	switch typ {
	case ir.INT:
		switch i.Typ {
		case ir.FLOAT:
			r.registers[r.Index] = float64(r.registers[i.Val].(int))

		case ir.STRING:
			r.registers[r.Index] = strconv.Itoa(r.registers[i.Val].(int))
		}

	case ir.FLOAT:
		switch i.Typ {
		case ir.INT:
			r.registers[r.Index] = int(r.registers[i.Val].(float64))

		case ir.STRING:
			r.registers[r.Index] = fmt.Sprintf("%f", r.registers[i.Val].(float64))
		}

	case ir.STRING:
		switch i.Typ {
		case ir.INT:
			v, err := strconv.Atoi(r.registers[i.Val].(string))
			if err != nil {
				return err
			}
			r.registers[r.Index] = v

		case ir.FLOAT:
			v, err := strconv.ParseFloat(r.registers[i.Val].(string), 64)
			if err != nil {
				return err
			}
			r.registers[r.Index] = v
		}
	}

	return nil
}

func (r *Runnable) runConcat(i *ir.Concat) {
	if len(i.Vals) == 2 {
		r.registers[r.Index] = r.registers[i.Vals[0]].(string) + r.registers[i.Vals[1]].(string)
		return
	}

	out := &strings.Builder{}
	for _, val := range i.Vals {
		out.WriteString(r.registers[val].(string))
	}
	r.registers[r.Index] = out.String()
}
