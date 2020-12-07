package main

import (
	"context"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	bags, err := buildTree(input)
	if err != nil {
		return nil, err
	}

	shinyGold := bagSliceContains(bags, "shiny gold")

	contains := shinyGold.countAllContains(0)
	return contains, nil
}

func (b *bag) countAllContains(count int) int {
	if len(b.contains) == 0 {
		return 0
	}

	subCount := 0
	for _, subBag := range b.contains {
		subCount = subCount + subBag.countAllContains(0)
	}
	return len(b.contains) + subCount
}
