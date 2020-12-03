package main

import (
	"context"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	a := countTrees(1, 1, lines)
	b := countTrees(3, 1, lines)
	c := countTrees(5, 1, lines)
	d := countTrees(7, 1, lines)
	e := countTrees(1, 2, lines)

	return a * b * c * d * e, nil
}
