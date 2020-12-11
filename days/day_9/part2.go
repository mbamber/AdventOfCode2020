package main

import (
	"context"
	"errors"
	"strconv"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	numbers := []int{}
	for _, line := range strings.Split(input, "\n") {
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, i)
	}

	var n int
	for i := 25; i < len(numbers); i = i + 1 {
		if !isSumOfPrevious(i, numbers) {
			n = numbers[i]
		}
	}

	for i := 0; i < n; i = i + 1 {
		sum, min, max := 0, numbers[i], 0
		for j := i; j < n; j = j + 1 {
			sum = sum + numbers[j]
			if numbers[j] < min {
				min = numbers[j]
			}
			if numbers[j] > max {
				max = numbers[j]
			}

			if sum > n {
				break
			} else if sum == n {
				return min + max, nil
			}
		}
	}

	return nil, errors.New("no solution found")
}
