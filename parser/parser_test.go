package parser

import "testing"

func TestParseCode(t *testing.T) {
	code := `[DEFINE
	 i 
	                      0]
	       [LOOP 10 [DEFINE i [MATH [VAR i] + 1]] [VAR i]]`
	ParseCode(code, NewPos("test.bpp"))
}
