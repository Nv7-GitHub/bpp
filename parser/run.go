package parser

import "fmt"

// Run is a basic runner based on the original B++ interpreter
func (p *Program) Run() (string, error) {
	out := ""
	for _, val := range p.Program {
		ret, err := val.Exec(p)
		if ret.Type == STRING {
			ret.Data = `"` + ret.Data.(string) + `"`
		}
		if err != nil {
			return out, err
		}
		if ret.Type != NULL {
			if ret.Type == ARRAY {
				out += "[ARRAY"
				for _, val := range ret.Data.([]Variable) {
					if val.Type == STRING {
						val.Data = `"` + val.Data.(string) + `"`
					}
					out += " " + fmt.Sprintf("%v", val.Data)
				}
				out += "]\n"
				continue
			}
			out += fmt.Sprintf("%v", ret.Data)
		}
	}
	return out, nil
}
