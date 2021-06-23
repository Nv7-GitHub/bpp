package gobpp

import (
	"fmt"
	"go/ast"
	"go/token"
)

var opMapAssign = map[token.Token]string{
	token.ADD_ASSIGN: "+",
	token.SUB_ASSIGN: "-",
	token.MUL_ASSIGN: "*",
	token.QUO_ASSIGN: "/",
}

func (p *Program) addAssignStmt(stm *ast.AssignStmt) error {
	r := stm.Rhs[0]

	name := stm.Lhs[0].(*ast.Ident).Name
	p.WriteString("[DEFINE ")
	p.WriteString(name)
	p.WriteString(" ")
	if stm.Tok == token.ASSIGN || stm.Tok == token.DEFINE {
		p.AddExpr(r)
		p.WriteString("]")
		return nil
	}

	op, exists := opMapAssign[stm.Tok]
	if !exists {
		return fmt.Errorf("%s: unknown operation %v", p.NodePos(stm), stm.Tok)
	}
	fmt.Fprintf(p, "[MATH [VAR %s] %s ", name, op)
	p.AddExpr(r)
	p.WriteString("]]")
	return nil
}

var incDecMap = map[token.Token]string{
	token.INC: "+",
	token.DEC: "-",
}

func (p *Program) addIncDecStmt(i *ast.IncDecStmt) error {
	name := i.X.(*ast.Ident).Name
	op, exists := incDecMap[i.Tok]
	if !exists {
		return fmt.Errorf("%s: unknown operation %v", p.NodePos(i), i.Tok)
	}

	fmt.Fprintf(p, "[DEFINE %s [MATH [VAR %s] %s 1]]", name, name, op)
	return nil
}
