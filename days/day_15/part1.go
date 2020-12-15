package main

import (
	"context"
	"strconv"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	starting := strings.Split(input, ",")
	numbers := []int{}
	cache := map[int][]int{} // Map of numbers to indexes

	for idx, start := range starting {
		i, err := strconv.Atoi(start)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, i)
		cache[i] = []int{idx + 1}
	}

	for i := len(numbers); i < 2020; i = i + 1 {
		next := genNextNumber(numbers, cache)
		numbers = append(numbers, next)
		cache[next] = append(cache[next], i+1)
	}

	return numbers[len(numbers)-1], nil
}

func genNextNumber(numbers []int, cache map[int][]int) int {
	lastNumber := numbers[len(numbers)-1]
	indexes := cache[lastNumber]
	if len(indexes) == 1 {
		return 0 // Never been said before
	}
	return indexes[len(indexes)-1] - indexes[len(indexes)-2]
}
