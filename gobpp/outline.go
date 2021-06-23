package gobpp

import (
	"fmt"
	"go/ast"
)

// AddFile adds an *ast.File to the program's source code
func (p *Program) AddFile(f *ast.File) error {
	for _, decl := range f.Decls {
		err := p.AddDecl(decl)
		if err != nil {
			return err
		}
	}
	return nil
}

// AddDecl adds a Go declaration to the program's source code
func (p *Program) AddDecl(decl ast.Decl) error {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		return p.addFuncDecl(d)

	default:
		return fmt.Errorf("%s: unknown declaration type %T", p.NodePos(decl), decl)
	}
}

// AddExpr adds an expression to the program's source code
func (p *Program) AddExpr(expr ast.Expr) error {
	switch e := expr.(type) {
	case *ast.BasicLit:
		return p.addBasicLit(e)

	case *ast.BinaryExpr:
		return p.addBinaryExpr(e)

	case *ast.CallExpr:
		return p.addCallExpr(e)

	case *ast.Ident:
		p.addIdent(e)
		return nil

	case *ast.IndexExpr:
		return p.addIndexExpr(e)

	case *ast.CompositeLit:
		return p.addCompositeLit(e)

	case *ast.UnaryExpr:
		return p.AddExpr(e.X)

	case *ast.ParenExpr:
		return p.AddExpr(e.X)

	default:
		return fmt.Errorf("%s: unknown expression type %T", p.NodePos(expr), expr)
	}
}

// AddStmt adds a statement to the program's source code
func (p *Program) AddStmt(stmt ast.Stmt) error {
	switch s := stmt.(type) {
	case *ast.ExprStmt:
		return p.addCallExpr(s.X.(*ast.CallExpr))

	case *ast.AssignStmt:
		return p.addAssignStmt(s)

	case *ast.ReturnStmt:
		return p.addReturnStmt(s)

	case *ast.IncDecStmt:
		return p.addIncDecStmt(s)

	case *ast.IfStmt:
		return p.addIfStmt(s)

	case *ast.ForStmt:
		return p.addForStmt(s)

	default:
		return fmt.Errorf("%s: unknown statement type %T", p.NodePos(stmt), stmt)
	}
}
