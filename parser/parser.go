package parser

// Program is the main program, containing the source AST
type Program struct {
	Statements []Statement
}

// Type is there to make a Program implement the Statement interface
func (p *Program) Type() DataType {
	return NULL
}

// Line is there to make a Program implement the Statement interface
func (p *Program) Line() int {
	return 1
}

// Keywords is there to make a Program implement the Block interface
func (p *Program) Keywords() []string {
	return []string{}
}

// End is there to make a Program implement the Block interface
func (p *Program) End(_ string, _ []Statement, stmts []Statement) bool {
	p.Statements = stmts
	return true
}

// EndSignature is there to make a Program implement the Block interface
func (p *Program) EndSignature() []DataType {
	return make([]DataType, 0)
}
