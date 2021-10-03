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
	gob.Register(&GetArg{})   // TODO
	gob.Register(&GetParam{}) // TODO
	gob.Register(&Cast{})
	gob.Register(&Math{})
	gob.Register(&Concat{})
	gob.Register(&PHI{})         // TODO
	gob.Register(&Jmp{})         // TODO
	gob.Register(&CondJmp{})     // TODO
	gob.Register(&JmpPoint{})    // TODO
	gob.Register(&Compare{})     // TODO
	gob.Register(&RandInt{})     // TODO
	gob.Register(&RandFloat{})   // TODO
	gob.Register(&ArrayIndex{})  // TODO
	gob.Register(&StringIndex{}) // TODO
	gob.Register(&Array{})
}

func (i *IR) Save(f io.Writer) error {
	return gob.NewEncoder(f).Encode(i)
}

func (i *IR) Load(f io.Reader) error {
	return gob.NewDecoder(f).Decode(i)
}
