package parser

import "fmt"

func ParseArgs(args []string, line int) ([]Statement, error) {
	out := make([]Statement, len(args))
	var err error
	for i, arg := range args {
		out[i], err = ParseStmt(arg, line)
		if err != nil {
			return []Statement{}, err
		}
	}
	return out, nil
}

func MatchTypes(data []Statement, line int, types ...DataType) error {
	if len(data) != len(types) {
		return fmt.Errorf("line %d: argument count doesn't match expected (expected %d, got %d)", line, len(types), len(data))
	}
	for i, arg := range data {
		if !arg.Type().IsEqual(types[i]) {
			return fmt.Errorf("line %d: argument %d is of wrong type", line, i+1)
		}
	}
	return nil
}
