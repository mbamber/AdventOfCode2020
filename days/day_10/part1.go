package main

import (
	"context"
	"sort"
	"strconv"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
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

	diffs := map[int]int{}
	for i := 0; i < len(numbers)-1; i = i + 1 {
		diffs[numbers[i+1]-numbers[i]] = diffs[numbers[i+1]-numbers[i]] + 1
	}

	return diffs[1] * diffs[3], nil
}
