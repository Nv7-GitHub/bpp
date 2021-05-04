package membuild

import (
	"fmt"
	"reflect"

	"github.com/Nv7-Github/Bpp/parser"
)

func BuildStmt(p *Program, stmt parser.Statement, instructionum ...int) (Instruction, error) {
	switch s := stmt.(type) {
	case *parser.SectionStmt:
		p.Sections[s.Label] = instructionum[0]
		return NewBlankInstruction(), nil

	case *parser.GotoStmt:
		return GotoStmt(p, s)

	case *parser.Data:
		d := ParserDataToData(s)
		return func(p *Program) (Data, error) { return d, nil }, nil

	case *parser.DefineStmt:
		return DefineStmt(p, s)

	case *parser.VarStmt:
		return VarStmt(p, s)

	case *parser.IfStmt:
		return IfStmt(p, s)

	case *parser.ComparisonStmt:
		return CompareStmt(p, s)

	default:
		return nil, fmt.Errorf("line %d: unknown type %s", stmt.Line(), reflect.TypeOf(stmt).String())
	}
}
