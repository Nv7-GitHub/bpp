package parser

import "fmt"

func GetStatement(fnName string, args []Statement) (Statement, error) {
	fmt.Println("GET STATEMENT:", fnName)

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
