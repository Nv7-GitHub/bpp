package compiler

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Nv7-Github/Bpp/parser"
)

//go:embed template.ctxt
var template string

func Compile(prog *parser.Program) (string, error) {
	variableTypes = make(map[string]parser.DataType)
	out := ""
	for _, statement := range prog.Statements {
		newV, _, err := compileStmt(statement)
		if err != nil {
			return "", err
		}
		out += newV + "\n"
	}
	return strings.Replace(template, "content: ;", out, 1), nil
}

func compileStmt(stmt parser.Statement) (string, parser.DataType, error) {
	raw, dt, err := compileStmtRaw(stmt)
	if err != nil {
		return "", 0, err
	}
	return formatVal(raw, dt), dt, nil
}

func formatVal(val string, kind parser.DataType) string {
	if kind != parser.NULL {
		formatter := printfMap[kind]
		formatter = strings.ReplaceAll(formatter, "val", val)
		return formatter
	}
	return val
}

type Data struct {
	Type  parser.DataType
	Value interface{}
}

func processData(d *parser.Data) Data {
	if d.Type().IsEqual(parser.STRING) {
		str, ok := d.Data.(string)
		if !ok {
			str = fmt.Sprintf("%v", d.Data)
		}
		return Data{
			Type:  parser.STRING,
			Value: str,
		}
	}
	if d.Type().IsEqual(parser.FLOAT) {
		f, ok := d.Data.(float64)
		if !ok {
			f = float64(d.Data.(int))
		}
		return Data{
			Type:  parser.FLOAT,
			Value: f,
		}
	}
	if d.Type().IsEqual(parser.INT) {
		return Data{
			Type:  parser.INT,
			Value: d.Data,
		}
	}
	return Data{
		Type:  parser.NULL,
		Value: d.Data,
	}
}

var printfMap = map[parser.DataType]string{
	parser.STRING: `printf("%s\n", val);`,
	parser.FLOAT:  `printf("%f\n", val);`,
	parser.INT:    `printf("%d\n", val);`,
}

var typeMap = map[parser.DataType]string{
	parser.STRING: "char *",
	parser.FLOAT:  "float ",
	parser.INT:    "int ",
}
