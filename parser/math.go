package parser

import "fmt"

var mathMap = map[string]Operator{
	"+": ADDITION,
	"-": SUBTRACTION,
	"*": MULTIPLICATION,
	"/": DIVISION,
	"^": POWER,
}

// MathStmt is the equivalent of [MATH stmt.Left stmt.Operation stmt.Right]
type MathStmt struct {
	*BasicStatement
	Operation Operator
	Left      Statement
	Right     Statement
}

func (m *MathStmt) Type() DataType {
	return INT | FLOAT
}

// RoundStmt is the equivalent of [ROUND stmt.Val]
type RoundStmt struct {
	*BasicStatement
	Val Statement
}

func (r *RoundStmt) Type() DataType {
	return INT
}

// FloorStmt is the equivalent of [FLOOR stmt.Val]
type FloorStmt struct {
	*BasicStatement
	Val Statement
}

func (f *FloorStmt) Type() DataType {
	return INT
}

// CeilStmt is the equivalent of [CEIL stmt.Val]
type CeilStmt struct {
	*BasicStatement
	Val Statement
}

func (c *CeilStmt) Type() DataType {
	return INT
}

func SetupMath() {
	parsers["MATH"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			dat, ok := args[1].(*Data)
			if !ok {
				return nil, fmt.Errorf("line %d: argument 2 to COMPARE must be an operator", line)
			}
			opTxt, ok := dat.Data.(string)
			if !ok {
				return nil, fmt.Errorf("line %d: argument 2 to COMPARE must be an operator", line)
			}
			op, exists := mathMap[opTxt]
			if !exists {
				return nil, fmt.Errorf("line %d: unknown comparison operator '%s'", line, opTxt)
			}
			return &MathStmt{
				Operation:      op,
				Left:           args[0],
				Right:          args[2],
				BasicStatement: &BasicStatement{line: line},
			}, nil
		},
		Signature: []DataType{ANY | NULL, IDENTIFIER, ANY | NULL},
	}

	parsers["ROUND"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &RoundStmt{
				BasicStatement: &BasicStatement{line: line},
				Val:            args[0],
			}, nil
		},
		Signature: []DataType{FLOAT},
	}

	parsers["FLOOR"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &FloorStmt{
				BasicStatement: &BasicStatement{line: line},
				Val:            args[0],
			}, nil
		},
		Signature: []DataType{FLOAT},
	}

	parsers["CEIL"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &CeilStmt{
				BasicStatement: &BasicStatement{line: line},
				Val:            args[0],
			}, nil
		},
		Signature: []DataType{FLOAT},
	}
}
