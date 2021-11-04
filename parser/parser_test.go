package parser

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParseCode(t *testing.T) {
	code := `[DEFINE l "etaoinshrdlucmfwypvbgkjqxz"]`
	prog := NewProgram()
	err := prog.Parse(code, "main.bpp")
	if err != nil {
		t.Fatal(err)
	}

	spew.Dump(prog)
}
