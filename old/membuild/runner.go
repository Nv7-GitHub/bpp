package membuild

// Run runs the program
func (p *Program) Run() error {
	for _, instruction := range p.Instructions {
		val, err := instruction(p)
		if err != nil {
			return err
		}
		err = p.Runner(val)
		if err != nil {
			return err
		}
	}
	return nil
}
