package builder

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

func (b *builder) stdFn(name string) *ir.Func {
	fn, exists := b.stdlib[name]
	if exists {
		return fn
	}

	switch name {
	case "pow":
		fn = b.mod.NewFunc("pow", types.Double, ir.NewParam("input", types.Double), ir.NewParam("power", types.Double))

	case "free":
		fn = b.mod.NewFunc("free", types.Void, ir.NewParam("src", types.I8Ptr))

	case "malloc":
		fn = b.mod.NewFunc("malloc", types.I8Ptr, ir.NewParam("size", types.I64))

	case "memcpy":
		fn = b.mod.NewFunc("memcpy", types.I8Ptr, ir.NewParam("dst", types.I8Ptr), ir.NewParam("src", types.I8Ptr), ir.NewParam("cnt", types.I64))

	case "printf":
		fn = b.mod.NewFunc("printf", types.I32, ir.NewParam("format", types.I8Ptr))
		fn.Sig.Variadic = true

	case "calloc":
		fn = b.mod.NewFunc("calloc", types.I8Ptr, ir.NewParam("num", types.I64), ir.NewParam("size", types.I64))
	}

	b.stdlib[name] = fn
	return fn
}
