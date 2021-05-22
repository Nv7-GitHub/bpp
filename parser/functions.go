package parser

import "fmt"

type FunctionType struct {
	Signature  []DataType
	Names      []string
	ReturnType DataType
}

var functionTypes map[string]FunctionType

// ParamStmt is the equivalent of [PARAM stmt.Name stmt.Type]
type ParamStmt struct {
	*BasicStatement
	Name string
	Kind DataType
}

func (p *ParamStmt) Type() DataType {
	return PARAMETER
}

// FunctionBlock is the equivalent of [IFB stmt.Condition] stmt.Body [ELSE] stmt.Else [ENDIF], the stmt.Else may be nil
type FunctionBlock struct {
	*BasicStatement

	Name      string
	Signature FunctionType
	Return    Statement
	Body      []Statement
}

func (f *FunctionBlock) Type() DataType {
	return f.Return.Type()
}

func (f *FunctionBlock) Keywords() []string {
	return []string{"RETURN"}
}

func (f *FunctionBlock) EndSignature() []DataType {
	return []DataType{ANY}
}

func (f *FunctionBlock) End(_ string, args []Statement, statements []Statement) bool {
	f.Return = args[0]
	f.Body = statements
	f.Signature.ReturnType = f.Return.Type()
	functionTypes[f.Name] = f.Signature
	return true
}

var dataTypes = map[string]DataType{
	"INT":    INT,
	"FLOAT":  FLOAT,
	"STRING": STRING,
	"ARRAY":  ARRAY,
}

func SetupFunctions() {
	parsers["PARAM"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			name, ok := args[0].(*Data)
			if !ok {
				return nil, fmt.Errorf("line %d: parameter 1 to PARAM must be constant", line)
			}

			kind, ok := args[1].(*Data)
			if !ok {
				return nil, fmt.Errorf("line %d: parameter 2 to PARAM must be constant", line)
			}

			k, exists := dataTypes[kind.Data.(string)]
			if !exists {
				return nil, fmt.Errorf("line %d: parameter 2 to PARAM must be INT, FLOAT, STRING, or ARRAY", line)
			}

			return &ParamStmt{
				BasicStatement: &BasicStatement{line: line},
				Name:           name.Data.(string),
				Kind:           k,
			}, nil
		},
		Signature: []DataType{IDENTIFIER, IDENTIFIER},
	}

	blocks["FUNCTION"] = BlockParser{
		Parse: func(args []Statement, line int) (Block, error) {
			sig := FunctionType{
				Signature: make([]DataType, len(args)-1),
				Names:     make([]string, len(args)-1),
			}
			for i, arg := range args[1:] {
				par, ok := arg.(*ParamStmt)
				if !ok {
					return nil, fmt.Errorf("line %d: parameters must be a PARAM", line)
				}

				sig.Signature[i] = par.Kind
				sig.Names[i] = par.Name
			}

			fn := &FunctionBlock{
				BasicStatement: &BasicStatement{line: line},
				Signature:      sig,
				Name:           args[0].(*Data).Data.(string),
			}

			functionTypes[fn.Name] = fn.Signature

			return fn, nil
		},
		Signature: []DataType{IDENTIFIER, PARAMETER, VARIADIC},
	}
}

// FunctionCallStmt is the equivalent of [stmt.Name stmt.Args...]
type FunctionCallStmt struct {
	*BasicStatement
	Name       string
	Args       []Statement
	ReturnType DataType
}

func (f *FunctionCallStmt) Type() DataType {
	return f.ReturnType
}
