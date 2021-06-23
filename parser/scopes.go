package parser

type empty struct{}

// Scope represents the data of a B++ scope
type Scope struct {
	Block      Block
	Statements []Statement

	keywords map[string]empty
}

// HasKeyword checks if a B++ scope contains the keyword for a statement efficiently
func (s *Scope) HasKeyword(keyword string) bool {
	_, exists := s.keywords[keyword]
	return exists
}

// NewScope creates a new initialized scope from a block
func NewScope(block Block) *Scope {
	s := &Scope{
		Statements: make([]Statement, 0),
		Block:      block,
	}

	keys := block.Keywords()
	s.keywords = make(map[string]empty, len(keys))
	for _, val := range keys {
		s.keywords[val] = empty{}
	}

	return s
}

// ScopeStack represents the data for a program's scopes - a Stack of scopes
type ScopeStack struct {
	scopes []*Scope
}

// GetScope gets the scope at the top of the stack
func (s *ScopeStack) GetScope() *Scope {
	return s.scopes[0]
}

// AddScope adds a scope to the stack
func (s *ScopeStack) AddScope(scope *Scope) {
	s.scopes = append([]*Scope{scope}, s.scopes...)
}

// FinishScope pops a scope off of the end of the stack, after processing the ending of the block
func (s *ScopeStack) FinishScope(keyword string, arguments []Statement) {
	remove := s.scopes[0].Block.End(keyword, arguments, s.scopes[0].Statements)
	if remove {
		if len(s.scopes) > 1 {
			s.scopes[1].Statements = append(s.scopes[1].Statements, s.scopes[0].Block)
		}
		s.scopes = s.scopes[1:]
	} else {
		s.scopes[0].Statements = make([]Statement, 0)
	}
}

// AddStatement adds a statement to the scope's statements
func (s *ScopeStack) AddStatement(stmt Statement) {
	s.scopes[0].Statements = append(s.scopes[0].Statements, stmt)
}

// NewScopeStack creates a new initialized stack of scopes
func NewScopeStack() *ScopeStack {
	return &ScopeStack{
		scopes: make([]*Scope, 0),
	}
}
