package membuild

import (
	"fmt"
	"io"

	"github.com/Nv7-Github/Bpp/parser"
)

func RunProgram(prog *Program, out io.Writer) error {
	i := 0
	for !(i == len(prog.Instructions)) {
		val, err := prog.Instructions[i](prog)
		if err != nil {
			return err
		}
		if val.Type == GOTO {
			i = val.Value.(int)
			continue
		}
		if val.Type != parser.NULL {
			txt := fmt.Sprintf("%v", val.Value)
			if len(txt) > 0 {
				_, err := out.Write([]byte(txt + "\n"))
				if err != nil {
					return err
				}
			}
		}
		i++
	}
	return nil
}
