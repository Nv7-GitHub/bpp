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
	return fmt.Sprintf("%s%s = %s;", typeMap[dt], label, v), parser.NULL, nil
}

func CompileVar(val *parser.VarStmt) (string, parser.DataType, error) {
	label := val.Label.(*parser.Data).Data.(string)
	return label, variableTypes[label], nil
}
