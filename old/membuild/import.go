package membuild

import "github.com/Nv7-Github/bpp/parser"

// ImportStmt compiles an IMPORT statement
func ImportStmt(p *Program, stm *parser.ImportStmt) (Instruction, error) {
	instrs := make([]Instruction, len(stm.Statements))
	var err error
	for i, stmt := range stm.Statements {
		instrs[i], err = BuildStmt(p, stmt, i)
		if err != nil {
			return nil, err
		}
	}

	return func(p *Program) (Data, error) {
		for _, stmt := range instrs {
			out, err := stmt(p)
			if err != nil {
				return NewBlankData(), err
			}

			err = p.Runner(out)
			if err != nil {
				return NewBlankData(), err
			}
		}

		return NewBlankData(), nil
	}, nil
}
