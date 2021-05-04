package parser

type Statement interface {
	Line() int
}

type StatementParser func(stmt string, line int) (Statement, error)
