package gobpp

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func (p *Program) addCallExpr(stm *ast.CallExpr) error {
	name := stm.Fun.(*ast.Ident).Name
	fn, exists := p.Funcs[name]
	if exists {
		return fn(stm.Args)
	}

	p.WriteString("[")
	p.WriteString(strings.ToUpper(name))

	for _, arg := range stm.Args {
		p.WriteString(" ")
		err := p.AddExpr(arg)
		if err != nil {
			return err
		}
	}
	p.WriteString("]")

	return nil
}

func (p *Program) addBasicLit(arg *ast.BasicLit) error {
	switch arg.Kind {
	case token.STRING, token.INT, token.FLOAT:
		p.WriteString(arg.Value)
		return nil

	default:
		return fmt.Errorf("%s: unknown data type: %s", p.NodePos(arg), arg.Kind.String())
	}
}

var opMap = map[token.Token]string{
	token.ADD: "+",
	token.SUB: "-",
	token.MUL: "*",
	token.QUO: "/",
	token.REM: "%",
}

var compMap = map[token.Token]string{
	token.EQL: "=",
	token.NEQ: "!=",
	token.GTR: ">",
	token.LSS: "<",
	token.GEQ: ">=",
	token.LEQ: "<=",
}

func (p *Program) addBinaryExpr(expr *ast.BinaryExpr) error {
	fnName := "MATH"

	op, exists := opMap[expr.Op]
	if !exists {
		fnName = "COMPARE"
		op = compMap[expr.Op]
	}

	p.WriteString("[")
	p.WriteString(fnName)
	p.WriteString(" ")

	err := p.AddExpr(expr.X)
	if err != nil {
		return err
	}

	p.WriteString(" ")
	p.WriteString(op)
	p.WriteString(" ")

	err = p.AddExpr(expr.Y)
	if err != nil {
		return err
	}

	p.WriteString("]")
	return nil
}

func (p *Program) addIdent(expr *ast.Ident) {
	fmt.Fprintf(p, "[VAR %s]", expr.Name)
}

func (p *Program) addIndexExpr(expr *ast.IndexExpr) error {
	p.WriteString("[INDEX ")

	err := p.AddExpr(expr.X)
	if err != nil {
		return err
	}

	p.WriteString(" ")

	err = p.AddExpr(expr.Index)
	if err != nil {
		return err
	}

	p.WriteString("]")

	return nil
}

func (p *Program) addCompositeLit(expr *ast.CompositeLit) error {
	p.WriteString("[ARRAY")
	for _, elt := range expr.Elts {
		p.WriteString(" ")
		err := p.AddExpr(elt)
		if err != nil {
			return err
		}
	}

	p.WriteString("]")
	return nil
}
