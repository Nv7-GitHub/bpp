package compiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

var printf *ir.Func
var strlen *ir.Func
var strcat *ir.Func
var strcmp *ir.Func
var sprintf *ir.Func
var gcvt *ir.Func

func generateBuiltins() {
	printf = m.NewFunc("printf", types.I32, ir.NewParam("format", types.I8Ptr))
	printf.Sig.Variadic = true
	strlen = m.NewFunc("strlen", types.I32, ir.NewParam("src", types.I8Ptr))
	strcat = m.NewFunc("strcat", types.I8Ptr, ir.NewParam("dst", types.I8Ptr), ir.NewParam("src", types.I8Ptr))
	strcmp = m.NewFunc("strcmp", types.I32, ir.NewParam("a", types.I8Ptr), ir.NewParam("b", types.I8Ptr))
	sprintf = m.NewFunc("sprintf", types.I32, ir.NewParam("dst", types.I8Ptr), ir.NewParam("format", types.I8Ptr))
	sprintf.Sig.Variadic = true
	gcvt = m.NewFunc("gcvt", types.I8Ptr, ir.NewParam("input", types.Double), ir.NewParam("length", types.I32), ir.NewParam("out", types.I8Ptr))
}
