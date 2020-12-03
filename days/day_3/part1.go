package main

import (
	"context"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	return countTrees(3, 1, lines), nil
}

func cellIsTree(cell byte) bool {
	return string(cell) == "#"
}

func countTrees(dx, dy int, lines []string) int {
	x, trees := 0, 0
	for y := 0; y < len(lines); y = y + dy {
		if cellIsTree(lines[y][x]) {
			trees = trees + 1
		}

		x = (x + dx) % len(lines[y])
	}

	return trees
}
