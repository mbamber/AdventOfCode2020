package main

import (
	"context"
	"errors"
	"strconv"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
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

			for k := j + 1; k < len(parts); k++ {
				c := parts[k]
				ci, err := strconv.Atoi(c)
				if err != nil {
					return nil, err
				}

				if ai+bi+ci == 2020 {
					return ai * bi * ci, nil
				}
			}
		}
	}

	return nil, errors.New("unexpected error")
}
