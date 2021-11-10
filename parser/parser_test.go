package parser

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParseCode(t *testing.T) {
	code := `[FUNCTION ADD [PARAM a INT] [PARAM b INT] [RETURNS INT] [BLOCK
		[RETURN [MATH [VAR a] + [VAR b]]]
	]]
	
	[ADD 1 2]`
	prog, err := Parse(code, "main.bpp")
	if err != nil {
		t.Fatal(err)
	}

	spew.Config.DisablePointerAddresses = true
	spew.Dump(prog)
}
