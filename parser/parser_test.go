package parser

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParseCode(t *testing.T) {
	util := `[FUNCTION ADD [PARAM a INT] [PARAM b INT] [RETURNS INT] [BLOCK
		[RETURN [MATH [VAR a] + [VAR b]]]
	]]
	[DEFINE i 0]`
	code := `
	[IMPORT "util.bpp"]
	[VAR i]
	[ADD 1 2]`
	prog, err := ParseMultifile(map[string]string{
		"main.bpp": code,
		"util.bpp": util,
	}, "main.bpp")
	if err != nil {
		t.Fatal(err)
	}

	spew.Config.DisablePointerAddresses = true
	spew.Dump(prog)
}
