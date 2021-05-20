package parser

// Program is the main program, containing the source AST
type Program struct {
	Statements []Statement
}

func (p *Program) Type() DataType {
	return NULL
}

func (p *Program) Line() int {
	return 1
}

func (p *Program) Keywords() []string {
	return []string{}
}

func (p *Program) End(_ string, _ []Statement, stmts []Statement) bool {
	p.Statements = stmts
	return true
}

func (p *Program) EndSignature() []DataType {
	return make([]DataType, 0)
}
