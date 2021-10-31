package parser

import "testing"

func TestParseCode(t *testing.T) {
	code := `[LOOP 
	10 [DEFINE i 
	
				[MATH [
					
				VAR i] + 1]
				
				
				] [VAR 
				
	i]]`
	ParseCode(code, NewPos("test.bpp"))
}
