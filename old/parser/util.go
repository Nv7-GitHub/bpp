package parser

import "fmt"

// ParseArgs parses the source code of arguments to a function
func ParseArgs(args []string, pos *Pos) ([]Statement, error) {
	out := make([]Statement, len(args))
	var err error
	for i, arg := range args {
		out[i], err = ParseStmt(arg, pos)
		if err != nil {
			return []Statement{}, err
		}
		if out[i] == nil {
			out[i] = &BasicStatement{
				pos: pos,
			}
		}
	}
	return out, nil
}

// MatchTypes compares 2 signatures, and is used in type-checking for function and block parsing. It supports variadic arguments.
func MatchTypes(data []Statement, pos *Pos, types []DataType) error {
	if len(types) > 1 && types[len(types)-1] == VARIADIC {
		for i, arg := range data {
			if i < len(types)-2 {
				if !arg.Type().IsEqual(types[i]) {
					return fmt.Errorf("%v: argument %d is of wrong type", pos, i+1)
				}
			} else {
				if !arg.Type().IsEqual(types[len(types)-2]) {
					return fmt.Errorf("%v: argument %d is of wrong type", pos, i+1)
				}
			}
		}
		return nil
	}

	if len(data) != len(types) {
		return fmt.Errorf("%v: argument count doesn't match expected (expected %d, got %d)", pos, len(types), len(data))
	}
	for i, arg := range data {
		if !arg.Type().IsEqual(types[i]) {
			return fmt.Errorf("%v: argument %d is of wrong type", pos, i+1)
		}
	}
	return nil
}
