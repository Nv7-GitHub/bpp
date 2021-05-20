package parser

// RandintStmt is the equivalent of [RANDINT stmt.Lower stmt.Upper]
type TestBlock struct {
	*BasicStatement
	Statements []Statement
}

func (b *TestBlock) Type() DataType {
	return NULL
}

func (b *TestBlock) Keywords() []string {
	return []string{"ENDTEST"}
}

func (b *TestBlock) EndSignature() []DataType {
	return []DataType{}
}

func (b *TestBlock) End(_ string, _ []Statement, statements []Statement) bool {
	b.Statements = statements
	return true
}

func SetupTestBlocks() {
	blocks["TEST"] = BlockParser{
		Parse: func(args []Statement, line int) (Block, error) {
			return &TestBlock{
				BasicStatement: &BasicStatement{line: line},
			}, nil
		},
		Signature: []DataType{},
	}
}
