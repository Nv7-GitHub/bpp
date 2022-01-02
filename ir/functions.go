package ir

import "github.com/Nv7-Github/bpp/types"

type Function struct {
	Instructions []Instruction
	Ret          int

	ParTypes []types.Type
	RetType  types.Type
}
