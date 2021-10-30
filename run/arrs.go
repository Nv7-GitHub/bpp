package run

import (
	"github.com/Nv7-Github/bpp/ir"
)

func (r *Runnable) runArray(i *ir.Array) {
	arr := make([]interface{}, len(i.Vals))
	for j, val := range i.Vals {
		arr[j] = r.registers[val]
	}
	r.registers[r.Index] = arr
}

func (r *Runnable) runArrayIndex(i *ir.ArrayIndex) {
	arr := r.registers[i.Array]
	ind := r.registers[i.Index].(int)
	r.registers[r.Index] = arr.([]interface{})[ind]
}

func (r *Runnable) runStringIndex(i *ir.StringIndex) {
	str := r.registers[i.Val]
	ind := r.registers[i.Index].(int)
	r.registers[r.Index] = string(str.(string)[ind])
}

func (r *Runnable) runArrayLength(i *ir.ArrayLength) {
	str := r.registers[i.Val]
	r.registers[r.Index] = len(str.([]interface{}))
}

func (r *Runnable) runStringLength(i *ir.StringLength) {
	str := r.registers[i.Val]
	r.registers[r.Index] = len(str.(string))
}
