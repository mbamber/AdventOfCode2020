package main

import (
	"context"
	"fmt"
	"math"
	"math/bits"
	"strconv"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	lines := strings.Split(input, "\n")

	var mask string = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	memory := map[int]string{}

	for _, line := range lines {
		parts := strings.Split(line, " = ")
		switch parts[0] {
		case "mask":
			mask = parts[1]
		default:
			id := parts[0][4 : len(parts[0])-1]
			adr, err := convertToBits(id)
			if err != nil {
				return nil, err
			}

			val, err := convertToBits(parts[1])
			if err != nil {
				return nil, err
			}

			adrs, err := applyMaskV2(mask, adr)
			if err != nil {
				return nil, err
			}
			for _, a := range adrs {
				memory[a] = val
			}

		}
	}

	var total int64 = 0
	for _, value := range memory {
		i, err := strconv.ParseInt(value, 2, 64)
		if err != nil {
			return nil, err
		}
		total = total + i
	}

	return total, nil
}

func applyMaskV2(mask, to string) ([]int, error) {
	// Count the number of floating bits in the mask (X)
	floating := 0
	for _, v := range mask {
		if string(v) == "X" {
			floating = floating + 1
		}
	}

	newVals := make([]string, int(math.Exp2(float64(floating))))
	options := make([]string, int(math.Exp2(float64(floating))))
	perms := getPerms(len(options))

	xCount := 0
	for i := 0; i < len(mask); i = i + 1 {
		switch string(mask[i]) {
		case "X":
			for j := range options {
				options[j] = string(perms[j][xCount])
			}
			xCount = xCount + 1
		case "1":
			for j := range options {
				options[j] = "1"
			}
		case "0":
			for j := range options {
				options[j] = string(to[i])
			}
		}

		// Update each string
		for j := range newVals {
			newVals[j] = newVals[j] + options[j]
		}
	}

	ids := []int{}
	for _, val := range newVals {
		id, err := strconv.ParseUint(val, 2, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, int(id))
	}

	return ids, nil
}

// getPerms gets all the
func getPerms(l int) []string {
	perms := []string{}
	for i := 0; i < l; i = i + 1 {
		perms = append(perms, fmt.Sprintf("%0*s", bits.Len(uint(l-1)), strconv.FormatUint(uint64(i), 2)))
	}
	return perms
}
