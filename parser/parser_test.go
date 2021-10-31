package parser

import "testing"

func TestParseCode(t *testing.T) {
	code := `[DEFINE l "etaoinshrdlucmfwypvbgkjqxz"]
	[DEFINE x [POW 26 0.5]]
	[INDEX
	  [VAR l]
	  [FLOOR
		[POW [RANDOM 0 [VAR x]] 2]
	  ]
	]`
	ParseCode(code, NewPos("test.bpp"))
}
