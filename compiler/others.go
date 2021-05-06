package compiler

import (
	"fmt"

	"github.com/Nv7-Github/Bpp/parser"
)

func CompileGoto(val *parser.GotoStmt) (string, parser.DataType, error) {
	label, _, err := compileStmtRaw(val.Label)
	if err != nil {
		return "", parser.NULL, err
	}
	if label[0] == '"' && label[len(label)-1] == '"' {
		label = label[1 : len(label)-1]
	}
	_, exists := variableTypes[label]
	if exists {
		tmp := tmpUsed
		tmpUsed++
		return fmt.Sprintf("void* tmp%d = &&%s; goto *tmp%d;", tmp, label, tmp), parser.NULL, nil
	}
	return fmt.Sprintf("goto %s;", label), parser.NULL, nil
}

func CompileSection(val *parser.SectionStmt) (string, parser.DataType, error) {
	return fmt.Sprintf("%s: ;", val.Label), parser.NULL, nil
}

func CompileArgs(val *parser.ArgsStmt) (string, parser.DataType, error) {
	ind, _, err := compileStmtRaw(val.Index)
	if err != nil {
		return "", parser.NULL, err
	}
	return fmt.Sprintf("argv[((int)%s)+1]", ind), parser.STRING, nil
}

func BuildVarMap() string {
	out := ""
	for k, v := range variableTypes {
		out += fmt.Sprintf("%s%s;", typeMap[v], k) + "\n"
	}
	return out
}

func CompileRepeat(val *parser.RepeatStmt) (string, parser.DataType, error) {
	ind, _, err := compileStmtRaw(val.Count)
	if err != nil {
		return "", parser.NULL, err
	}
	v, _, err := compileStmtRaw(val.Val)
	if err != nil {
		return "", parser.NULL, err
	}
	return fmt.Sprintf("repeat(%s, (int)%s)", v, ind), parser.STRING, nil
}

func CompileChoose(val *parser.ChooseStmt) (string, parser.DataType, error) {
	v, _, err := compileStmtRaw(val.Data)
	if err != nil {
		return "", parser.NULL, err
	}
	return fmt.Sprintf("choose(%s)", v), parser.STRING, nil
}
