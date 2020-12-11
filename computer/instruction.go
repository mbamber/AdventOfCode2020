package computer

import (
	"errors"
	"strconv"
	"strings"
)

// ParseInstruction parses a string into an Instruction
func ParseInstruction(s string) (Instruction, error) {
	parts := strings.Split(s, " ")
	arg, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	switch parts[0] {
	case "acc":
		return &InstructionAcc{Arg: arg}, nil
	case "jmp":
		return &InstructionJmp{Arg: arg}, nil
	case "nop":
		return &InstructionNoOp{Arg: arg}, nil
	}
	return nil, errors.New("unknown instruction")
}

// Instruction represents a single Instruction that can be performed by a Computer
type Instruction interface {
	Execute(c *Computer)
}

// InstructionAcc represents the Acc instruction that adjusts the computer's accumalator value
type InstructionAcc struct {
	Arg int
}

// Execute the instruction in the context of the given Computer
func (i *InstructionAcc) Execute(c *Computer) {
	c.Acc = c.Acc + i.Arg
	c.Ipr = c.Ipr + 1
}

// InstructionJmp represents the Jmp instruction that adjusts the commputer's instruction pointer
type InstructionJmp struct {
	Arg int
}

// Execute the instruction in the context of the given Computer
func (i *InstructionJmp) Execute(c *Computer) {
	c.Ipr = c.Ipr + i.Arg
}

// InstructionNoOp represents the NoOp instruction that does nothing except increment the
// computer's instruction pointer
type InstructionNoOp struct {
	Arg int
}

// Execute the instruction in the context of the given Computer
func (i *InstructionNoOp) Execute(c *Computer) {
	c.Ipr = c.Ipr + 1
}
