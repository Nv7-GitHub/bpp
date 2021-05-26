package compiler

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var printf *ir.Func

var strlen *ir.Func
var strcmp *ir.Func
var sprintf *ir.Func
var gcvt *ir.Func
var pow *ir.Func
var strtol *ir.Func
var strtod *ir.Func
var srand *ir.Func
var time *ir.Func
var rand *ir.Func
var malloc *ir.Func
var memset *ir.Func
var memcpy *ir.Func
var free *ir.Func
var floor *ir.Func
var ceil *ir.Func
var round *ir.Func
var sin *ir.Func
var fabs *ir.Func

var intFmt *ir.Global
var strFmt *ir.Global
var fltFmt *ir.Global
var newLine *ir.Global
var openBracket *ir.Global
var closeBracket *ir.Global
var comma *ir.Global

var args *ir.Global

func generateBuiltins() {
	printf = m.NewFunc("printf", types.I32, ir.NewParam("format", types.I8Ptr))
	printf.Sig.Variadic = true
	strlen = m.NewFunc("strlen", types.I64, ir.NewParam("src", types.I8Ptr))
	malloc = m.NewFunc("malloc", types.I8Ptr, ir.NewParam("len", types.I64))
	memset = m.NewFunc("memset", types.I8Ptr, ir.NewParam("str", types.I8Ptr), ir.NewParam("char", types.I32), ir.NewParam("len", types.I64)) // DEBUG
	memcpy = m.NewFunc("memcpy", types.I8Ptr, ir.NewParam("dst", types.I8Ptr), ir.NewParam("src", types.I8Ptr), ir.NewParam("cnt", types.I64))
	free = m.NewFunc("free", types.Void, ir.NewParam("src", types.I8Ptr))
	strcmp = m.NewFunc("strcmp", types.I32, ir.NewParam("a", types.I8Ptr), ir.NewParam("b", types.I8Ptr))
	sprintf = m.NewFunc("sprintf", types.I32, ir.NewParam("dst", types.I8Ptr), ir.NewParam("format", types.I8Ptr))
	sprintf.Sig.Variadic = true
	gcvt = m.NewFunc("gcvt", types.I8Ptr, ir.NewParam("input", types.Double), ir.NewParam("length", types.I32), ir.NewParam("out", types.I8Ptr))
	pow = m.NewFunc("pow", types.Double, ir.NewParam("input", types.Double), ir.NewParam("power", types.Double))
	strtol = m.NewFunc("strtol", types.I64, ir.NewParam("input", types.I8Ptr), ir.NewParam("remaining", types.NewPointer(types.I8Ptr)), ir.NewParam("base", types.I32))
	strtod = m.NewFunc("strtod", types.Double, ir.NewParam("input", types.I8Ptr), ir.NewParam("remaining", types.NewPointer(types.I8Ptr)))
	srand = m.NewFunc("srand", types.Void, ir.NewParam("seed", types.I32))
	time = m.NewFunc("time", types.I64, ir.NewParam("time", types.I64Ptr))
	rand = m.NewFunc("rand", types.I32)
	floor = m.NewFunc("floor", types.Double, ir.NewParam("input", types.Double))
	ceil = m.NewFunc("ceil", types.Double, ir.NewParam("input", types.Double))
	round = m.NewFunc("round", types.Double, ir.NewParam("input", types.Double))
	sin = m.NewFunc("sin", types.Double, ir.NewParam("input", types.Double))
	fabs = m.NewFunc("fabs", types.Double, ir.NewParam("input", types.Double))

	intFmt = m.NewGlobalDef("intfmt", constant.NewCharArrayFromString("%ld"+string(rune(0))))
	strFmt = m.NewGlobalDef("strfmt", constant.NewCharArrayFromString("%s"+string(rune(0))))
	fltFmt = m.NewGlobalDef("fltfmt", constant.NewCharArrayFromString("%f"+string(rune(0))))
	newLine = m.NewGlobalDef("newline", constant.NewCharArrayFromString("\n"+string(rune(0))))
	openBracket = m.NewGlobalDef("openbracket", constant.NewCharArrayFromString("["+string(rune(0))))
	closeBracket = m.NewGlobalDef("closebracket", constant.NewCharArrayFromString("]"+string(rune(0))))
	comma = m.NewGlobalDef("comma", constant.NewCharArrayFromString(", "+string(rune(0))))

	args = m.NewGlobalDef("args", constant.NewNull(types.NewPointer(types.I8Ptr)))
}

func initMod(block *ir.Block) {
	// Initialize random
	time := block.NewCall(time, constant.NewNull(types.I64Ptr))
	trunced := block.NewTrunc(time, types.I32)
	block.NewCall(srand, trunced)

	// Initialize program args
	block.NewStore(block.Parent.Params[1], args) // Store argv to args
}

func getStrPtr(val value.Value, block *ir.Block) value.Value {
	glob, ok := val.(*ir.Global)
	if ok {
		return block.NewGetElementPtr(glob.ContentType, val, constant.NewInt(types.I64, 0), constant.NewInt(types.I64, 0))
	}
	newv, ok := val.Type().(*types.PointerType)
	if ok {
		_, ok = newv.ElemType.(*types.ArrayType)
		if ok {
			return block.NewBitCast(val, types.I8Ptr)
		}
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
