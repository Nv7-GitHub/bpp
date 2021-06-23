package gobpp

import (
	"go/ast"
)

func (p *Program) addIfStmt(i *ast.IfStmt) error {
	p.WriteString("[IFB ")
	err := p.AddExpr(i.Cond)
	if err != nil {
		return err
	}
	p.WriteString("]\n")

	err = p.AddBlock(i.Body.List)
	if err != nil {
		return err
	}

	if i.Else != nil {
		p.WriteString("[ELSE]\n")
		err = p.AddBlock(i.Else.(*ast.BlockStmt).List)
		if err != nil {
			return err
		}
	}

	p.WriteString("[ENDIF]")
	return nil
}

func (p *Program) addForStmt(stm *ast.ForStmt) error {
	if stm.Init != nil {
		err := p.AddStmt(stm.Init)
		if err != nil {
			return err
		}

		p.WriteString("\n")
	}

	p.WriteString("[WHILE ")
	err := p.AddExpr(stm.Cond)
	if err != nil {
		return err
	}
	p.WriteString("]\n")

	err = p.AddBlock(stm.Body.List)
	if err != nil {
		return err
	}

	if stm.Post != nil {
		err = p.AddStmt(stm.Post)
		if err != nil {
			return err
		}
		p.WriteString("\n")
	}

	p.WriteString("[ENDWHILE]")
	return nil
}
