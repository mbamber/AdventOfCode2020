package main

import (
	"context"
	"sort"
	"strconv"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	lines := strings.Split(input, "\n")
	numbers := []int{}
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, i)
	}

	// Add 0 for the charing outlet
	numbers = append(numbers, 0)

	sort.Ints(numbers)

	// Add the maximum + 3 for the device
	numbers = append(numbers, numbers[len(numbers)-1]+3)

	// Find all of subslices
	// These are found by breaking the slice of numbers wherever there is a gap of exactly 3. This works
	// because these numbers MUST be included in each permutation
	subslices := [][]int{}
	currMinIndex := 0
	for i := 1; i < len(numbers); i = i + 1 {
		if numbers[i]-numbers[i-1] == 3 {
			subslice := numbers[currMinIndex:i]
			subslices = append(subslices, subslice)
			currMinIndex = i
		}
	}
	// Don't bother adding the last subslice (which is just the last number), because it doesn't change the number of overall permutations

	// Find the total by multiplying together all the combinations of all the subslices
	total := 1
	for _, slice := range subslices {
		total = total * countRoutes(slice)
	}

	return total, nil
}

// countRoutes by first considering the fact that the first number has exactly one route.
// Then the number of routes to each subsequent number can be found by adding the number of
// routes to reach the previous number (each number may be able to reach 1, 2, or 3 other numbers)
func countRoutes(s []int) int {
	waysToReach := map[int]int{0: 1}
	for i := range s {
		for j := i + 1; j < i+4; j = j + 1 {
			if j >= len(s) {
				continue
			}
			if s[j]-s[i] <= 3 {
				waysToReach[j] = waysToReach[j] + waysToReach[i]
			}
		}
	}
	return waysToReach[len(waysToReach)-1]
}
