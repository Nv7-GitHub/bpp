package parser

func GetStatement(fnName string, args []Statement, pos *Pos) (Statement, error) {
	parser, exists := parsers[fnName]
	if !exists {
		return nil, pos.NewError("unknown function: %s", fnName)
	}
	// TODO: Do something with parser
	// NOTE: Match types
	return nil, nil
}

func MatchTypes(a []Statement, b []Type) {
	i := 0
	j := 0
	for i < len(a) {
		if b[j] == VARIADIC {
			for a[i].Type().Equal(b[j-1]) {
				i++
			}
		}
		if a[i].Type().Equal(b[j]) {
			i++
			j++
		}
	}
}
