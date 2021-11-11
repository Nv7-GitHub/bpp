package ir

type Instruction interface {
	String() string
	Type() Type
}
