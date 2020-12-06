package main

import (
	"context"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	entries := strings.Split(input, "\n")

	var valid int
	for _, password := range entries {
		ok, err := passwordIsValidP2(password)
		if err != nil {
			return nil, err
		}
		if ok {
			valid++
		}
	}

	return valid, nil
}

func passwordIsValidP2(p string) (bool, error) {
	i, j, char, password, err := parsePassword(p)
	if err != nil {
		return false, err
	}

	return (string(password[i-1]) == char) != (string(password[j-1]) == char), nil // A XOR B <-> A != B
}
