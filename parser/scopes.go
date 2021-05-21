package parser

type empty struct{}

type Scope struct {
	Block      Block
	Statements []Statement

	keywords map[string]empty
}

func (s *Scope) HasKeyword(keyword string) bool {
	_, exists := s.keywords[keyword]
	return exists
}

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

type ScopeStack struct {
	scopes []*Scope
}

func (s *ScopeStack) GetScope() *Scope {
	return s.scopes[0]
}

func (s *ScopeStack) AddScope(scope *Scope) {
	s.scopes = append([]*Scope{scope}, s.scopes...)
}

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

func (s *ScopeStack) AddStatement(stmt Statement) {
	s.scopes[0].Statements = append(s.scopes[0].Statements, stmt)
}

func NewScopeStack() *ScopeStack {
	return &ScopeStack{
		scopes: make([]*Scope, 0),
	}
}
