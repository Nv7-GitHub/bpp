package compiler

import (
	"fmt"
	"reflect"

	"github.com/Nv7-Github/Bpp/parser"
)

var variableTypes map[string]parser.DataType

func CompileData(val *parser.Data) (string, parser.DataType, error) {
	d := processData(val)
	if d.Type.IsEqual(parser.STRING) {
		return fmt.Sprintf(`"%s"`, d.Value), parser.STRING, nil
	}
	if d.Type.IsEqual(parser.INT) {
		return fmt.Sprintf(`%d`, d.Value), parser.INT, nil
	}
	if d.Type.IsEqual(parser.FLOAT) {
		return fmt.Sprintf(`%f`, d.Value), parser.FLOAT, nil
	}
	return "", parser.NULL, fmt.Errorf("line %d: unknown type %s", val.Line(), reflect.TypeOf(val))
}

func CompileDefine(val *parser.DefineStmt) (string, parser.DataType, error) {
	label := val.Label.(*parser.Data).Data.(string)
	v, dt, err := compileStmtRaw(val.Value)
	if err != nil {
		return "", parser.NULL, err
	}
	variableTypes[label] = dt
	return fmt.Sprintf("%s = %s;", label, v), parser.NULL, nil
}

func CompileVar(val *parser.VarStmt) (string, parser.DataType, error) {
	label := val.Label.(*parser.Data).Data.(string)
	return label, variableTypes[label], nil
}

func CompileRandint(val *parser.RandintStmt) (string, parser.DataType, error) {
	lower, _, err := compileStmtRaw(val.Lower)
	if err != nil {
		return "", parser.NULL, err
	}
	upper, _, err := compileStmtRaw(val.Upper)
	if err != nil {
		return "", parser.NULL, err
	}
	return fmt.Sprintf("(rand() %% ((int)%s - (int)%s)) + (int)%s", upper, lower, lower), parser.INT, nil
}

func CompileIndex(val *parser.IndexStmt) (string, parser.DataType, error) {
	ind, _, err := compileStmtRaw(val.Index)
	if err != nil {
		return "", parser.NULL, err
	}
	dat, _, err := compileStmtRaw(val.Value)
	if err != nil {
		return "", parser.NULL, err
	}
	return fmt.Sprintf("std::string(1, %s[(int)%s])", dat, ind), parser.STRING, nil
}

func CompileRandom(val *parser.RandomStmt) (string, parser.DataType, error) {
	lower, _, err := compileStmtRaw(val.Lower)
	if err != nil {
		return "", parser.NULL, err
	}
	upper, _, err := compileStmtRaw(val.Upper)
	if err != nil {
		return "", parser.NULL, err
	}
	return fmt.Sprintf("((float)rand()/(float)(RAND_MAX)) * ((float)%s - (float)%s)", lower, upper), parser.FLOAT, nil
}

func CompileConcat(val *parser.ConcatStmt) (string, parser.DataType, error) {
	if len(val.Strings) == 2 {
		one, _, err := compileStmtRaw(val.Strings[0])
		if err != nil {
			return "", parser.NULL, err
		}
		two, _, err := compileStmtRaw(val.Strings[1])
		if err != nil {
			return "", parser.NULL, err
		}
		return fmt.Sprintf("std::string(%s) + std::string(%s)", one, two), parser.STRING, nil
	}
	first, _, err := compileStmtRaw(val.Strings[0])
	if err != nil {
		return "", parser.NULL, err
	}
	val.Strings = val.Strings[1:]
	rest, _, err := CompileConcat(val)
	if err != nil {
		return "", parser.NULL, err
	}
	return fmt.Sprintf("std::string(%s) + %s", first, rest), parser.STRING, nil
}
