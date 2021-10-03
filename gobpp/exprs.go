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

	_, _ = p.WriteString("[")
	_, _ = p.WriteString(strings.ToUpper(name))

	for _, arg := range stm.Args {
		_, _ = p.WriteString(" ")
		err := p.AddExpr(arg)
		if err != nil {
			return err
		}
	}
	_, _ = p.WriteString("]")

	return nil
}

func (p *Program) addBasicLit(arg *ast.BasicLit) error {
	switch arg.Kind {
	case token.STRING, token.INT, token.FLOAT:
		_, _ = p.WriteString(arg.Value)
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

	_, _ = p.WriteString("[")
	_, _ = p.WriteString(fnName)
	_, _ = p.WriteString(" ")

	err := p.AddExpr(expr.X)
	if err != nil {
		return err
	}

	_, _ = p.WriteString(" ")
	_, _ = p.WriteString(op)
	_, _ = p.WriteString(" ")

	err = p.AddExpr(expr.Y)
	if err != nil {
		return err
	}

	_, _ = p.WriteString("]")
	return nil
}

func (p *Program) addIdent(expr *ast.Ident) {
	_, _ = fmt.Fprintf(p, "[VAR %s]", expr.Name)
}

func (p *Program) addIndexExpr(expr *ast.IndexExpr) error {
	_, _ = p.WriteString("[INDEX ")

	err := p.AddExpr(expr.X)
	if err != nil {
		return err
	}

	_, _ = p.WriteString(" ")

	err = p.AddExpr(expr.Index)
	if err != nil {
		return err
	}

	_, _ = p.WriteString("]")

	return nil
}

func (p *Program) addCompositeLit(expr *ast.CompositeLit) error {
	_, _ = p.WriteString("[ARRAY")
	for _, elt := range expr.Elts {
		_, _ = p.WriteString(" ")
		err := p.AddExpr(elt)
		if err != nil {
			return err
		}
	}

	_, _ = p.WriteString("]")
	return nil
}
