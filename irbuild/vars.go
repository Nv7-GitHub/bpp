package irbuild

import (
	"github.com/Nv7-Github/bpp/types"

	"github.com/Nv7-Github/bpp/parser"
)

func isDynamic(typ types.Type) bool {
	return typ.BasicType().Equal(types.STRING) || typ.BasicType().Equal(types.ARRAY)
}

func (i *IRBuilder) addDefine(s *parser.DefineStmt) (int, error) {
	valind, err := i.AddStmt(s.Val)
	if err != nil {
		return 0, err
	}
	typ := i.GetInstruction(valind).Type()

	name := s.Variable
	v, exists := i.vars[name]
	if !exists {
		// Alloc memory
		dynamic := isDynamic(typ)
		var mem int
		if !dynamic {
			mem = i.NewAllocStatic(typ)
		} else {
			mem = i.NewAllocDynamic(typ)
		}
		v = varData{
			Mem: mem,
			Typ: typ,
		}
		i.vars[name] = v
	}

	// Set memory
	if isDynamic(v.Typ) {
		return i.NewSetMemoryDynamic(v.Mem, valind), nil
	}
	return i.NewSetMemory(v.Mem, valind), nil
}
