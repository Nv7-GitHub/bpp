package ir

import (
	"encoding/gob"
	"io"
)

func init() {
	gob.Register(&FunctionCall{})
	gob.Register(&Print{})
	gob.Register(&Const{})
	gob.Register(&AllocStatic{})
	gob.Register(&AllocDynamic{})
	gob.Register(&SetMemory{})
	gob.Register(&SetMemoryDynamic{})
	gob.Register(&GetMemory{})
	gob.Register(&GetMemoryDynamic{})
	gob.Register(&GetArg{})
	gob.Register(&GetParam{})
	gob.Register(&Cast{})
	gob.Register(&Math{})
	gob.Register(&Concat{})
	gob.Register(&PHI{})
	gob.Register(&Jmp{})
	gob.Register(&CondJmp{})
	gob.Register(&JmpPoint{})
	gob.Register(&Compare{})
	gob.Register(&RandInt{})
	gob.Register(&RandFloat{})
	gob.Register(&ArrayIndex{})
	gob.Register(&StringIndex{})
	gob.Register(&Array{})
}

func (i *IR) Save(f io.Writer) error {
	return gob.NewEncoder(f).Encode(i)
}

func (i *IR) Load(f io.Reader) error {
	return gob.NewDecoder(f).Decode(i)
}
