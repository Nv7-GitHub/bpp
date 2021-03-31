package parser

import "fmt"

// Run is a basic runner based on the original B++ interpreter
func (p *Program) Run() (string, error) {
	out := ""
	for _, val := range p.Program {
		ret, err := val(p)
		if err != nil {
			return out, err
		}
		if ret.Type.IsEqual(STRING) {
			if len(ret.Data.(string)) == 0 {
				ret.Type = NULL
			} else {
				ret.Data = `"` + ret.Data.(string) + `"`
			}
		}
		if !ret.Type.IsEqual(NULL) {
			if ret.Type.IsEqual(ARRAY) {
				out += "[ARRAY"
				for _, val := range ret.Data.([]Variable) {
					if val.Type.IsEqual(STRING) {
						val.Data = `"` + val.Data.(string) + `"`
					}
					out += " " + fmt.Sprintf("%v", val.Data)
				}
				out += "]\n"
				continue
			}
			out += fmt.Sprintf("%v\n", ret.Data)
		}
	}
	return out, nil
}
