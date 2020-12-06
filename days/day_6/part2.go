package main

import (
	"context"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	groups := strings.Split(input, "\n\n")
	// groups = []string{groups[5], groups[6]}

	return Count(groups, all), nil
}

func all(people []string) int {
	questions := map[rune]int{}

	for _, p := range people {
		for _, q := range p {
			c := questions[q]
			questions[q] = c + 1
		}
	}

	common := []rune{}
	for q, c := range questions {
		if c == len(people) {
			common = append(common, q)
		}
	}

	return len(common)
}
