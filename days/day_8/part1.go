package main

import (
	"context"
	"strings"

	"github.com/mbamber/aoc/computer"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	c, err := computer.New(strings.Split(input, "\n"))
	if err != nil {
		return nil, err
	}

	err = c.Run(computer.NewIprLoopDetector())
	if err != computer.ErrInfiniteLoop {
		return nil, err
	}

	return c.Acc, nil
}
