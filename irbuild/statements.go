package irbuild

import "github.com/Nv7-Github/bpp/parser"

func (i *IRBuilder) AddStmt(stmt parser.Statement) (int, error) {
	switch s := stmt.(type) {
	case *parser.Const:
		return i.addConst(s)

	case *parser.DefineStmt:
		return i.addDefine(s)

	default:
		return 0, stmt.Pos().NewError("unknown statement type: %T", s)
	}
}
