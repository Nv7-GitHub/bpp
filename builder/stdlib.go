package builder

import (
	"fmt"

	"github.com/llir/irutil"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
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

	case "sprintf":
		fn = b.mod.NewFunc("sprintf", types.I32, ir.NewParam("dst", types.I8Ptr), ir.NewParam("format", types.I8Ptr))
		fn.Sig.Variadic = true

	case "gcvt":
		fn = b.mod.NewFunc("gcvt", types.I8Ptr, ir.NewParam("input", types.Double), ir.NewParam("length", types.I32), ir.NewParam("out", types.I8Ptr))

	case "strtol":
		fn = b.mod.NewFunc("strtol", types.I64, ir.NewParam("input", types.I8Ptr), ir.NewParam("remaining", types.NewPointer(types.I8Ptr)), ir.NewParam("base", types.I32))

	case "strtod":
		fn = b.mod.NewFunc("strtod", types.Double, ir.NewParam("input", types.I8Ptr), ir.NewParam("remaining", types.NewPointer(types.I8Ptr)))

	case "strlen":
		fn = b.mod.NewFunc("strlen", types.I64, ir.NewParam("src", types.I8Ptr))

	case "strcmp":
		fn = b.mod.NewFunc("strcmp", types.I32, ir.NewParam("a", types.I8Ptr), ir.NewParam("b", types.I8Ptr))

	case "rand":
		fn = b.mod.NewFunc("rand", types.I32)

	case "sin":
		fn = b.mod.NewFunc("sin", types.Double, ir.NewParam("input", types.Double))

	case "fabs":
		fn = b.mod.NewFunc("fabs", types.Double, ir.NewParam("input", types.Double))

	case "time":
		fn = b.mod.NewFunc("time", types.I64, ir.NewParam("time", types.I64Ptr))

	case "srand":
		fn = b.mod.NewFunc("srand", types.Void, ir.NewParam("seed", types.I32))
	}

	b.stdlib[name] = fn
	return fn
}

func (b *builder) stdV(name string) value.Value {
	v, exists := b.stdv[name]
	if exists {
		return v
	}

	switch name {
	case "fmt":
		b.newStdval("fmt", "%s\n")
	case "intfmt":
		b.newStdval("intfmt", "%d")
	case "ptrfmt":
		b.newStdval("ptrfmt", "%p\n")
	case "intprint":
		b.newStdval("intprint", "%d\n")
	}

	return b.stdv[name]
}

func (b *builder) newStdval(name string, val string) {
	glob := b.mod.NewGlobalDef(fmt.Sprintf("%s%d", name, b.tmpCount), irutil.NewCString(val))
	b.tmpCount++
	v := b.block.NewGetElementPtr(types.NewArray(uint64(len(val)+1), types.I8), glob, constant.NewInt(types.I32, 0), constant.NewInt(types.I32, 0))
	b.stdv[name] = v
}
