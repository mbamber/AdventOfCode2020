package main

import (
	"context"
	"errors"
	"strings"

	"github.com/mbamber/aoc/computer"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	c, err := computer.New(strings.Split(input, "\n"))
	if err != nil {
		return nil, err
	}

	for i, instruction := range c.Program {
		cCopy := c.Copy()

		switch inst := instruction.(type) {
		case *computer.InstructionNoOp:
			cCopy.Program[i] = &computer.InstructionJmp{
				Arg: inst.Arg,
			}
		case *computer.InstructionJmp:
			cCopy.Program[i] = &computer.InstructionNoOp{
				Arg: inst.Arg,
			}
		}

		err := cCopy.Run(computer.NewIprLoopDetector())
		if err == nil {
			return cCopy.Acc, nil
		}
	}

	return nil, errors.New("no valid solution found")
}
