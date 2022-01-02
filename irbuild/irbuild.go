package irbuild

import (
	"github.com/Nv7-Github/bpp/types"

	"github.com/Nv7-Github/bpp/ir"
)

type IRBuilder struct {
	*ir.IR

	vars map[string]varData
}

type varData struct {
	Mem int
	Typ types.Type
}
