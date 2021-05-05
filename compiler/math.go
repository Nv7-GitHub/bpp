package compiler

import (
	"fmt"

	"github.com/Nv7-Github/Bpp/parser"
)

func CompileMath(val *parser.MathStmt) (string, parser.DataType, error) {
	left, ldt, err := compileStmtRaw(val.Left)
	if err != nil {
		return "", parser.NULL, err
	}
	right, rdt, err := compileStmtRaw(val.Right)
	if err != nil {
		return "", parser.NULL, err
	}

	if val.Operation == parser.POWER {
		isFloat := false
		if rdt.IsEqual(parser.FLOAT) || ldt.IsEqual(parser.FLOAT) {
			isFloat = true
		}
		if isFloat {
			return fmt.Sprintf("pow((float)%s, (float)%s)", left, right), parser.FLOAT, nil
		}
		return fmt.Sprintf("ipow((int)%s, (int)%s)", left, right), parser.INT, nil
	}

	kind := "int"
	type_ := parser.INT
	if rdt.IsEqual(parser.FLOAT) || ldt.IsEqual(parser.FLOAT) {
		kind = "float"
		type_ = parser.FLOAT
	}
	if type_ == parser.INT && (ldt == parser.STRING || rdt == parser.STRING) {
		if ldt == parser.STRING {
			left = fmt.Sprintf("stoi(%s, &strsz)", left)
		}
		if rdt == parser.STRING {
			right = fmt.Sprintf("stoi(%s, &strsz)", right)
		}
	}
	return fmt.Sprintf("(%s)%s %s (%s)%s", kind, left, opMap[val.Operation], kind, right), type_, nil
}

var opMap = map[parser.Operator]string{
	parser.ADDITION:       "+",
	parser.SUBTRACTION:    "-",
	parser.MULTIPLICATION: "*",
	parser.DIVISION:       "/",
}

func CompileFloor(val *parser.FloorStmt) (string, parser.DataType, error) {
	dat, _, err := compileStmtRaw(val.Val)
	if err != nil {
		return "", parser.NULL, err
	}
	return fmt.Sprintf("floor((float)%s)", dat), parser.STRING, nil
}

func CompileCeil(val *parser.CeilStmt) (string, parser.DataType, error) {
	dat, _, err := compileStmtRaw(val.Val)
	if err != nil {
		return "", parser.NULL, err
	}
	return fmt.Sprintf("ceil((float)%s)", dat), parser.STRING, nil
}

func CompileRound(val *parser.RoundStmt) (string, parser.DataType, error) {
	dat, _, err := compileStmtRaw(val.Val)
	if err != nil {
		return "", parser.NULL, err
	}
	return fmt.Sprintf("round(float)%s)", dat), parser.STRING, nil
}
