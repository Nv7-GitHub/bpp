package parser

func init() {
	// Statements
	SetupGotos()
	SetupVariables()
	SetupArrays()
	SetupRandoms()
	SetupComparisons()
	SetupMath()
	SetupOthers()

	// Blocks
	SetupBlocks()
}
