package parser

import (
	_ "embed"
	"testing"
)

//go:embed test.bpp
var code string

func TestParser(t *testing.T) {
	_, err := Parse(code, "test.bpp")
	if err != nil {
		panic(err)
	}
}
