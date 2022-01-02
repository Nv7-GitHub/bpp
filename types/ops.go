package types

type MathOp int

const (
	MathOpAdd MathOp = iota
	MathOpSub
	MathOpMul
	MathOpDiv
	MathOpMod
	MathOpPow
)

var mathOpNames = map[MathOp]string{
	MathOpAdd: "+",
	MathOpSub: "-",
	MathOpMul: "*",
	MathOpDiv: "/",
	MathOpMod: "%",
	MathOpPow: "^",
}

func (m MathOp) String() string {
	return mathOpNames[m]
}
