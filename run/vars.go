package run

import "github.com/Nv7-Github/bpp/ir"

func (r *Runnable) runAllocStatic(index int, i *ir.AllocStatic) {
	r.registers[index] = len(r.vars)
	r.vars[len(r.vars)] = nil
}

func (r *Runnable) runAllocDynamic(index int, i *ir.AllocDynamic) {
	r.registers[index] = len(r.vars)
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

func (r *Runnable) runGetMemory(index int, i *ir.GetMemory) {
	mem := r.registers[i.Mem].(int)
	r.registers[index] = r.vars[mem]
}

func (r *Runnable) runGetMemoryDynamic(index int, i *ir.GetMemoryDynamic) {
	mem := r.registers[i.Mem].(int)
	r.registers[index] = r.vars[mem]
}
