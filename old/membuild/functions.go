package membuild

import (
	"github.com/Nv7-Github/bpp/old/parser"
)

// Function stores the data for a compiled function.
type Function struct {
	Body   []Instruction
	Name   string
	Type   parser.FunctionType
	Return Instruction
}

// FunctionBlock compiles a FUNCTION block.
func FunctionBlock(p *Program, stm *parser.FunctionBlock) (Instruction, error) {
	body := make([]Instruction, len(stm.Body))
	for i, stm := range stm.Body {
		stmt, err := BuildStmt(p, stm)
		if err != nil {
			return nil, err
		}
		body[i] = stmt
	}
	ret, err := BuildStmt(p, stm.Return)
	if err != nil {
		return nil, err
	}

	return func(p *Program) (Data, error) {
		fn := Function{
			Name:   stm.Name,
			Body:   body,
			Type:   stm.Signature,
			Return: ret,
		}
		p.Functions[fn.Name] = fn

		return NewBlankData(), nil
	}, nil
}

// FunctionCallStmt compiles a function call
func FunctionCallStmt(p *Program, stm *parser.FunctionCallStmt) (Instruction, error) {
	args := make([]Instruction, len(stm.Args))
	for i, stm := range stm.Args {
		stmt, err := BuildStmt(p, stm)
		if err != nil {
			return nil, err
		}
		args[i] = stmt
	}

	return func(p *Program) (Data, error) {
		fn := p.Functions[stm.Name]

		// Create stack and add parameters to memory
		stack := createStack(p)
		var err error
		for i, arg := range args {
			stack.Memory[fn.Type.Names[i]], err = arg(p)
			if err != nil {
				return NewBlankData(), err
			}
		}

		// Execute function
		for _, inst := range fn.Body {
			out, err := inst(stack)
			if err != nil {
				return NewBlankData(), err
			}

			err = stack.Runner(out)
			if err != nil {
				return NewBlankData(), err
			}
		}

		return fn.Return(stack)
	}, nil
}

func createStack(p *Program) *Program {
	out := &Program{
		Instructions: p.Instructions,
		Memory:       make(map[string]Data), // Don't copy memory, functions aren't allowed to access global memory
		Functions:    make(map[string]Function),
		Args:         p.Args,
		Runner:       p.Runner,
	}

	for k, v := range p.Functions {
		out.Functions[k] = v
	}

	return out
}
