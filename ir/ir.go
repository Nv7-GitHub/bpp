package ir

import "github.com/Nv7-Github/bpp/types"

type Instruction interface {
	String() string
	Type() types.Type
}
