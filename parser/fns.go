package parser

import "fmt"

func GetStatement(fnName string, args []Statement) (Statement, error) {
	fmt.Println("GET STATEMENT:", fnName)

	return nil, nil
}

func GetConst(text string) Statement {
	fmt.Println("CONST", text)
	return nil
}
