package main

import (
	"context"
	"strconv"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	lines := strings.Split(input, "\n")
	estimate, err := strconv.Atoi(lines[0])
	if err != nil {
		return nil, err
	}

	busses := strings.Split(lines[1], ",")
	waitTimes := map[int]int{}
	for _, bus := range busses {
		if bus == "x" {
			continue
		}

		busID, err := strconv.Atoi(bus)
		if err != nil {
			return nil, err
		}

		waitTime := (((estimate / busID) + 1) * busID) % estimate
		waitTimes[busID] = waitTime
	}

	bestBus := 0
	for id, waitTime := range waitTimes {
		if bestBus == 0 {
			bestBus = id
			continue
		}

		if waitTime < waitTimes[bestBus] {
			bestBus = id
		}
	}

	return bestBus * waitTimes[bestBus], nil
}
