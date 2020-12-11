package main

import (
	"context"
	"errors"
	"strconv"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	numbers := []int{}
	for _, line := range strings.Split(input, "\n") {
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, i)
	}

	for i := 25; i < len(numbers); i = i + 1 {
		if !isSumOfPrevious(i, numbers) {
			return numbers[i], nil
		}
	}

	return nil, errors.New("no solution found")
}

func isSumOfPrevious(i int, numbers []int) bool {
	for a := i - 25; a < i-1; a = a + 1 {
		for b := a + 1; b < i; b = b + 1 {
			if numbers[a]+numbers[b] == numbers[i] {
				return true
			}
		}
	}
	return false
}
