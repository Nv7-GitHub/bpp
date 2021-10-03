package run

import "github.com/Nv7-Github/bpp/ir"

func (r *Runnable) runAllocStatic() {
	r.registers[r.Index] = len(r.vars)
	r.vars[len(r.vars)] = nil
}

func (r *Runnable) runAllocDynamic() {
	r.registers[r.Index] = len(r.vars)
	r.vars[len(r.vars)] = nil
}

func (r *Runnable) runSetMemory(i *ir.SetMemory) {
	mem := r.registers[i.Mem].(int)
	r.vars[mem] = r.registers[i.Value]
}

func (r *Runnable) runSetMemoryDynamic(i *ir.SetMemoryDynamic) {
	mem := r.registers[i.Mem].(int)
	r.vars[mem] = r.registers[i.Value]
}

func (r *Runnable) runGetMemory(i *ir.GetMemory) {
	mem := r.registers[i.Mem].(int)
	r.registers[r.Index] = r.vars[mem]
}

func (r *Runnable) runGetMemoryDynamic(i *ir.GetMemoryDynamic) {
	mem := r.registers[i.Mem].(int)
	r.registers[r.Index] = r.vars[mem]
}
