package gobpp

import (
	"go/ast"
	"strings"
)

// Function defines the signature for a function parser
type Function func(args []ast.Expr) error

// Init initializes a program
func (p *Program) Init() {
	p.Builder = &strings.Builder{}
	p.Funcs = make(map[string]Function)

	p.Funcs["print"] = func(args []ast.Expr) error {
		return p.AddExpr(args[0])
	}
}

// End cleans up a program and adds running code
func (p *Program) End() {
	p.WriteString("[MAIN]")
}
