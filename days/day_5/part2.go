package main

import (
	"context"
	"errors"
	"sort"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	tickets := strings.Fields(input)

	ids := []int{}
	for _, ticket := range tickets {
		ids = append(ids, GetID(ticket))
	}

	sort.Ints(ids)

	for i, id := range ids {
		if id+1 != ids[i+1] {
			return id + 1, nil
		}
	}
	return nil, errors.New("ticket not found")
}
