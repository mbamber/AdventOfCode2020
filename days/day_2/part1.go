package main

import (
	"context"
	"strconv"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	entries := strings.Split(input, "\n")

	var valid int
	for _, password := range entries {
		ok, err := passwordIsValidP1(password)
		if err != nil {
			return nil, err
		}
		if ok {
			valid++
		}
	}

	return valid, nil
}

func passwordIsValidP1(p string) (bool, error) {
	min, max, char, password, err := parsePassword(p)
	if err != nil {
		return false, err
	}

	numTimes := strings.Count(password, char)
	return numTimes >= min && numTimes <= max, nil
}

func parsePassword(p string) (i, j int, char, password string, err error) {
	parts := strings.Fields(p)
	minmax := strings.Split(parts[0], "-")

	i, err = strconv.Atoi(minmax[0])
	if err != nil {
		return 0, 0, "", "", err
	}

	j, err = strconv.Atoi(minmax[1])
	if err != nil {
		return 0, 0, "", "", err
	}

	char = strings.TrimSuffix(parts[1], ":")
	password = parts[2]

	return
}
