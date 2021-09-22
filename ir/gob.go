package ir

import (
	"encoding/gob"
	"io"
)

func init() {
	gob.Register(&FunctionCall{}) // TODO
	gob.Register(&Print{})
	gob.Register(&Const{})
	gob.Register(&AllocStatic{})
	gob.Register(&AllocDynamic{})
	gob.Register(&SetMemory{})
	gob.Register(&SetMemoryDynamic{})
	gob.Register(&GetMemory{})
	gob.Register(&GetMemoryDynamic{})
	gob.Register(&GetArg{})
	gob.Register(&GetParam{}) // TODO
	gob.Register(&Cast{})
	gob.Register(&Math{})
	gob.Register(&Concat{})
	gob.Register(&PHI{}) // TODO
	gob.Register(&Jmp{})
	gob.Register(&CondJmp{})
	gob.Register(&JmpPoint{})
	gob.Register(&Compare{})
	gob.Register(&RandInt{})
	gob.Register(&RandFloat{})
	gob.Register(&ArrayIndex{})  // TODO
	gob.Register(&StringIndex{}) // TODO
	gob.Register(&Array{})       // TODO
}

func (i *IR) Save(f io.Writer) error {
	return gob.NewEncoder(f).Encode(i)
}

func (i *IR) Load(f io.Reader) error {
	return gob.NewDecoder(f).Decode(i)
}
