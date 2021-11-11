package run

import (
	"errors"
	"math/rand"

	"github.com/Nv7-Github/bpp/old/ir"
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
	max := r.registers[i.Max].(int)
	min := r.registers[i.Min].(int)
	r.registers[r.Index] = rand.Intn(max-min) + min
}

func (r *Runnable) runRandFloat(i *ir.RandFloat) {
	max := r.registers[i.Max].(float64)
	min := r.registers[i.Min].(float64)
	r.registers[r.Index] = rand.Float64()*(max-min) + min
}

func (r *Runnable) runPHI(i *ir.PHI) {
	cond := r.registers[i.Cond].(int) == 1
	if cond {
		r.registers[r.Index] = r.registers[i.ValTrue]
	} else {
		r.registers[r.Index] = r.registers[i.ValFalse]
	}
}
