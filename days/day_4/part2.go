package main

import (
	"context"
	"regexp"
	"strconv"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	records := parse(strings.Split(input, "\n\n"))

	hclRegexp := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	pidRegexp := regexp.MustCompile(`^[0-9]{9}$`)

	valid := 0
	for _, record := range records {
		recordIsValid := true
		for _, key := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
			val, ok := record[key]
			if !ok {
				recordIsValid = false
				break
			}

			switch key {
			case "byr":
				i, err := strconv.Atoi(val)
				if err != nil {
					recordIsValid = false
					break
				}
				if i < 1920 || i > 2002 {
					recordIsValid = false
					break
				}
			case "iyr":
				i, err := strconv.Atoi(val)
				if err != nil {
					recordIsValid = false
					break
				}
				if i < 2010 || i > 2020 {
					recordIsValid = false
					break
				}
			case "eyr":
				i, err := strconv.Atoi(val)
				if err != nil {
					recordIsValid = false
					break
				}
				if i < 2020 || i > 2030 {
					recordIsValid = false
					break
				}
			case "hgt":
				suff := "cm"
				n := strings.TrimSuffix(val, "cm")
				if n == val {
					suff = "in"
					n = strings.TrimSuffix(val, "in")
				}
				if n == val { // Doesn't end with cm or in
					recordIsValid = false
					break
				}

				i, err := strconv.Atoi(n)
				if err != nil {
					recordIsValid = false
					break
				}
				switch suff {
				case "cm":
					if i < 150 || i > 193 {
						recordIsValid = false
						break
					}
				case "in":
					if i < 59 || i > 76 {
						recordIsValid = false
						break
					}
				}
			case "hcl":
				if !hclRegexp.MatchString(val) {
					recordIsValid = false
					break
				}
			case "ecl":
				if !(val == "amb" || val == "blu" || val == "brn" || val == "gry" || val == "grn" || val == "hzl" || val == "oth") {
					recordIsValid = false
					break
				}
			case "pid":
				if !pidRegexp.MatchString(val) {
					recordIsValid = false
					break
				}
			}
		}
		if recordIsValid {
			valid = valid + 1
		}
	}

	return valid, nil
}
