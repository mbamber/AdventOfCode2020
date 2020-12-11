package computer

// Computer represents a programmable computer
type Computer struct {
	usedIprs []int

	Acc     int
	Ipr     int
	Program []Instruction
}

// New parses a set of raw instructions and returns a new Computer
func New(rawProgram []string) (*Computer, error) {
	program := []Instruction{}
	for _, instruction := range rawProgram {
		i, err := ParseInstruction(instruction)
		if err != nil {
			return nil, err
		}
		program = append(program, i)
	}

	return &Computer{
		usedIprs: []int{},

		Acc:     0,
		Ipr:     0,
		Program: program,
	}, nil
}

// Copy returns a copy of a computer, with the same state and program
func (c *Computer) Copy() *Computer {
	var newProgram = make([]Instruction, len(c.Program))
	copy(newProgram, c.Program)

	var newUsedIprs = make([]int, len(c.usedIprs))
	copy(newUsedIprs, c.usedIprs)

	return &Computer{
		usedIprs: newUsedIprs,

		Acc:     c.Acc,
		Ipr:     c.Ipr,
		Program: newProgram,
	}
}

// Run starts execution and continues until the program terminates.
// An ErrInfiniteLoop error is returned if an infinite loop is detected
func (c *Computer) Run(ld LoopDetector) error {
	for {
		if err := ld.Detect(c); err != nil {
			return err
		}
		ld.Update(c)

		// Successful execution
		if c.Ipr >= len(c.Program) {
			return nil
		}

		c.Step()
	}
}

// Step runs the next instruction
func (c *Computer) Step() {
	c.Program[c.Ipr].Execute(c)
}
