package compiler

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var printf *ir.Func

//var strlen *ir.Func
var strcat *ir.Func
var strcmp *ir.Func
var sprintf *ir.Func
var gcvt *ir.Func
var pow *ir.Func

var intFmt *ir.Global
var strFmt *ir.Global
var fltFmt *ir.Global

func generateBuiltins() {
	printf = m.NewFunc("printf", types.I32, ir.NewParam("format", types.I8Ptr))
	printf.Sig.Variadic = true
	//strlen = m.NewFunc("strlen", types.I32, ir.NewParam("src", types.I8Ptr)) // For when LENGTH comes
	strcat = m.NewFunc("strcat", types.I8Ptr, ir.NewParam("dst", types.I8Ptr), ir.NewParam("src", types.I8Ptr))
	strcmp = m.NewFunc("strcmp", types.I32, ir.NewParam("a", types.I8Ptr), ir.NewParam("b", types.I8Ptr))
	sprintf = m.NewFunc("sprintf", types.I32, ir.NewParam("dst", types.I8Ptr), ir.NewParam("format", types.I8Ptr))
	sprintf.Sig.Variadic = true
	gcvt = m.NewFunc("gcvt", types.I8Ptr, ir.NewParam("input", types.Double), ir.NewParam("length", types.I32), ir.NewParam("out", types.I8Ptr))
	pow = m.NewFunc("pow", types.Double, ir.NewParam("input", types.Double), ir.NewParam("power", types.Double))

	intFmt = m.NewGlobalDef("intfmt", constant.NewCharArrayFromString("%ld\n"+string(rune(0))))
	strFmt = m.NewGlobalDef("strfmt", constant.NewCharArrayFromString("%s\n"+string(rune(0))))
	fltFmt = m.NewGlobalDef("fltfmt", constant.NewCharArrayFromString("%f\n"+string(rune(0))))
}

func getStrPtr(val value.Value, block *ir.Block) value.Value {
	glob, ok := val.(*ir.Global)
	if ok {
		return block.NewGetElementPtr(glob.ContentType, val, constant.NewInt(types.I64, 0), constant.NewInt(types.I64, 0))
	}
	return val
}

var tmpUsed int

func getTmp() string {
	name := fmt.Sprintf("tmp%d", tmpUsed)
	tmpUsed++
	return name
}

func getStr(str string) *ir.Global {
	return m.NewGlobalDef(getTmp(), constant.NewCharArrayFromString(str+string(rune(0))))
}
