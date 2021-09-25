package run

import (
	"errors"
	"math/rand"

	"github.com/Nv7-Github/bpp/ir"
)

func (r *Runnable) runGetArg(i *ir.GetArg) error {
	ind := r.registers[i.Index].(int)
	if ind >= len(r.args) {
		return errors.New("GetArg: index out of range")
	}
	r.registers[r.Index] = r.args[ind]
	return nil
}

func (r *Runnable) runCondJmp(i *ir.CondJmp) {
	cond := r.registers[i.Cond].(int) == 1
	if cond {
		r.Index = i.TargetTrue
	} else {
		r.Index = i.TargetFalse
	}
}

func (r *Runnable) runJmp(i *ir.Jmp) {
	r.Index = i.Target
}

func (r *Runnable) runRandInt(i *ir.RandInt) {
	r.registers[r.Index] = rand.Intn(i.Max-i.Min) + i.Min
}

func (r *Runnable) runRandFloat(i *ir.RandFloat) {
	r.registers[r.Index] = rand.Float64()*float64(i.Max-i.Min) + float64(i.Min)
}

func (r *Runnable) runPHI(i *ir.PHI) {
	cond := r.registers[i.Cond].(int) == 1
	if cond {
		r.registers[r.Index] = r.registers[i.ValTrue]
	} else {
		r.registers[r.Index] = r.registers[i.ValFalse]
	}
}
