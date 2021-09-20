package run

import "github.com/Nv7-Github/bpp/ir"

func (r *Runnable) runConst(index int) {
	r.registers[index] = r.ir.Instructions[index].(*ir.Const).Data
}

func (r *Runnable) runPrint(print *ir.Print) error {
	_, err := r.Stdout.Write([]byte(r.registers[print.Val].(string) + "\n"))
	return err
}
