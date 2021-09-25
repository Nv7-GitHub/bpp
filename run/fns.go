package run

import (
	"github.com/Nv7-Github/bpp/ir"
)

func (r *Runnable) runFunctionCall(i *ir.FunctionCall) error {
	pars := make([]interface{}, len(i.Params))
	for j, p := range i.Params {
		pars[j] = r.registers[p]
	}

	f := r.ir.Functions[i.Fn]

	fn := &Runnable{
		ir: &ir.IR{
			Instructions: f.Instructions,
			Functions:    r.ir.Functions,
		},
		Stdout:    r.Stdout,
		registers: make([]interface{}, len(f.Instructions)),
		vars:      make(map[int]interface{}),
		params:    pars,
	}

	err := fn.Run(r.args)
	if err != nil {
		return err
	}

	r.registers[r.Index] = fn.registers[f.Ret]
	return nil
}

func (r *Runnable) runGetParam(i *ir.GetParam) {
	r.registers[r.Index] = r.params[i.Index]
}
