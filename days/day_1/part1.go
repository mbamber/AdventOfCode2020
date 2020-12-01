package main

import (
	"context"
	"errors"
	"strconv"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	parts := strings.Fields(input)
	for i, a := range parts {
		ai, err := strconv.Atoi(a)
		if err != nil {
			return nil, err
		}

		for j := i + 1; j < len(parts); j++ {
			b := parts[j]

			bi, err := strconv.Atoi(b)
			if err != nil {
				return nil, err
			}

			if ai+bi == 2020 {
				return ai * bi, nil
			}
		}
	}

	return nil, errors.New("unexpected error")
}
