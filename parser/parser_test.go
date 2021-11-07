package parser

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParseCode(t *testing.T) {
	code := `[DEFINE i 0][WHILE [COMPARE 1 = 1]


	[VAR i]
		[DEFINE i [MATH [VAR i] + 1]]]
		
		[IF 
			[COMPARE 1 = 1] 
			[BLOCK 
				[VAR i] 
				"ok cool"
			] 
			
			"hi"
		]`
	prog := NewProgram()
	err := prog.Parse(code, "main.bpp")
	if err != nil {
		t.Fatal(err)
	}

	spew.Config.DisablePointerAddresses = true
	spew.Dump(prog)
}
