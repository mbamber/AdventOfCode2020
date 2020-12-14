package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
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
			adr, err := strconv.Atoi(id)
			if err != nil {
				return nil, err
			}

			bits, err := convertToBits(parts[1])
			if err != nil {
				return nil, err
			}
			memory[adr] = applyMask(mask, bits)
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

func convertToBits(s string) (string, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return "", err
	}
	newS := strconv.FormatUint(i, 2)
	return fmt.Sprintf("%036s", newS), nil
}

func applyMask(mask, to string) string {
	newVal := ""
	for i := 0; i < len(mask); i = i + 1 {
		switch string(mask[i]) {
		case "X":
			newVal = newVal + string(to[i])
		case "1":
			newVal = newVal + "1"
		case "0":
			newVal = newVal + "0"
		}
	}
	return newVal
}
