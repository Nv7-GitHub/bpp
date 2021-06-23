package gobpp

import (
	"go/ast"
	"go/token"
	"strings"
)

// Program defines the program source code and some program data
type Program struct {
	*strings.Builder

	Fset     *token.FileSet
	Funcs    map[string]Function
	FuncName string
}

// Pos uses the Fset to convert a token.Pos to a string
func (p *Program) Pos(pos token.Pos) string {
	return p.Fset.Position(pos).String()
}

// NodePos gets the position of a ast.Node and converts that to a string with the same method as Pos
func (p *Program) NodePos(node ast.Node) string {
	return p.Fset.Position(node.Pos()).String()
}

// Convert converts a parsed Go source code file and returns B++ source code
func Convert(fset *token.FileSet, filename string, f *ast.File) (string, error) {
	p := &Program{
		Fset: fset,
	}
	p.Init()

	err := p.AddFile(f)
	if err != nil {
		return "", err
	}

	p.End()
	return p.String(), nil
}
