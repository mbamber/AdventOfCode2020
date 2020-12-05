package main

import (
	"context"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	tickets := strings.Fields(input)

	maxID := 0
	for _, ticket := range tickets {
		id := GetID(ticket)
		if id > maxID {
			maxID = id
		}
	}

	return maxID, nil
}

func parse(s string) (row, column int) {
	minRow, maxRow := 0, 128
	minColumn, maxColumn := 0, 7

	for i := 0; i < 7; i++ {
		minRow, maxRow = Split(minRow, maxRow, GetUpper(s[i]))
	}

	for i := 7; i < 10; i++ {
		minColumn, maxColumn = Split(minColumn, maxColumn, GetUpper(s[i]))
	}

	return minRow, minColumn
}

func GetUpper(c byte) bool {
	return string(c) == "B" || string(c) == "R"
}

func Split(min, max int, upper bool) (newMin, newMax int) {
	middle := (max - min + 1) / 2
	if upper {
		return min + middle, max
	} else {
		return min, min + middle - 1
	}
}

func GetID(s string) int {
	row, column := parse(s)
	return getID(row, column)
}

func getID(row, column int) int {
	return (row * 8) + column
}
