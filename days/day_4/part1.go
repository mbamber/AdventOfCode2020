package main

import (
	"context"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	records := parse(strings.Split(strings.TrimSpace(input), "\n\n"))

	valid := 0
	for _, record := range records {
		recordIsValid := true
		for _, key := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
			if _, ok := record[key]; !ok {
				recordIsValid = false
			}
		}
		if recordIsValid {
			valid = valid + 1
		}
	}

	return valid, nil
}

func parse(records []string) []map[string]string {
	result := []map[string]string{}

	for _, record := range records {
		recordMap := map[string]string{}
		for _, field := range strings.Fields(record) {
			parts := strings.Split(field, ":")
			recordMap[parts[0]] = parts[1]
		}
		result = append(result, recordMap)
	}
	return result
}
