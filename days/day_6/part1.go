package main

import (
	"context"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	groups := strings.Split(strings.TrimSpace(input), "\n\n")

	return Count(groups, any), nil
}

func Count(groups []string, groupCounter func(s []string) int) int {
	total := 0
	for _, group := range groups {
		people := strings.Split(group, "\n")
		total = total + groupCounter(people)
	}
	return total
}

func any(people []string) int {
	groupQuestions := []rune{}
	for _, person := range people {
		for _, q := range person {
			groupQuestions = append(groupQuestions, q)
		}
	}

	unique := []rune{}
	for _, r := range groupQuestions {
		if !runeSliceContains(unique, r) {
			unique = append(unique, r)
		}
	}
	return len(unique)
}

func runeSliceContains(s []rune, r rune) bool {
	for _, a := range s {
		if a == r {
			return true
		}
	}

	return false
}
