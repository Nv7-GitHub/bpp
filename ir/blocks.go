package ir

import "github.com/Nv7-Github/bpp/parser"

func (i *IR) addIfB(stmt *parser.IfBlock) (int, error) {
	cond, err := i.AddStmt(stmt.Condition)
	if err != nil {
		return 0, err
	}

	jmp := i.newCondJmp(cond)

	ifTrue := i.newJmpPoint()
	for _, stmt := range stmt.Body {
		_, err = i.AddStmtTop(stmt)
		if err != nil {
			return 0, err
		}
	}
	ifTrueEnd := i.newJmp()

	ifFalse := i.newJmpPoint()
	if stmt.Else != nil {
		for _, stmt := range stmt.Else {
			_, err = i.AddStmtTop(stmt)
			if err != nil {
				return 0, err
			}
		}
	}
	ifFalseEnd := i.newJmp()

	end := i.newJmpPoint()
	i.SetCondJmpPoint(jmp, ifTrue, ifFalse)
	i.SetJmpPoint(ifTrueEnd, end)
	i.SetJmpPoint(ifFalseEnd, end)

	return end, nil
}

func (i *IR) addWhile(stmt *parser.WhileBlock) (int, error) {
	start := i.newJmpPoint()

	cond, err := i.AddStmt(stmt.Condition)
	if err != nil {
		return 0, err
	}
	condJmp := i.newCondJmp(cond)

	body := i.newJmp()
	for _, stmt := range stmt.Body {
		if _, err := i.AddStmtTop(stmt); err != nil {
			return 0, err
		}
	}

	startJmp := i.newJmp()
	i.SetJmpPoint(startJmp, start)

	end := i.newJmpPoint()
	i.SetCondJmpPoint(condJmp, body, end)

	return end, nil
}
