package main

import (
	"context"
	"strconv"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	rules, _, nearbyTickets, err := parseInput(input)
	if err != nil {
		return nil, err
	}

	sum := 0
	for _, ticket := range nearbyTickets {
		_, invalid := ticketIsValid(ticket, rules)
		for _, i := range invalid {
			sum = sum + i
		}
	}

	return sum, nil
}

type intRange struct {
	min int
	max int
}

type rule struct {
	name   string
	ranges []intRange
}

func (r *rule) contains(i int) bool {
	for _, r := range r.ranges {
		if i >= r.min && i <= r.max {
			return true
		}
	}

	return false
}

func ticketIsValid(ticket []int, rules []rule) (valid, invalid []int) {
	for _, i := range ticket {
		ok := false
		for _, r := range rules {
			if r.contains(i) {
				ok = true
				valid = append(valid, i)
				break
			}
		}

		if !ok {
			invalid = append(invalid, i)
		}
	}
	return valid, invalid
}

func parseRule(s string) (rule, error) {
	r := rule{}
	parts := strings.Split(s, ": ")

	r.name = parts[0]

	ranges := strings.Split(parts[1], " or ")
	for _, rng := range ranges {
		minmax := strings.Split(rng, "-")
		min, err := strconv.Atoi(minmax[0])
		if err != nil {
			return r, err
		}
		max, err := strconv.Atoi(minmax[1])
		if err != nil {
			return r, err
		}

		r.ranges = append(r.ranges, intRange{min: min, max: max})
	}

	return r, nil
}

func parseInput(input string) (rules []rule, myTicket []int, nearbyTickets [][]int, err error) {
	sections := strings.Split(input, "\n\n")

	// Parse the rules
	ruleStrings := strings.Split(sections[0], "\n")
	for _, rule := range ruleStrings {
		r, err := parseRule(rule)
		if err != nil {
			return nil, nil, nil, err
		}
		rules = append(rules, r)
	}

	// Our ticket
	ticket := strings.Split(sections[1], "\n")[1]
	myTicket, err = parseTicket(ticket)
	if err != nil {
		return nil, nil, nil, err
	}

	// Other tickets
	for i, ticket := range strings.Split(sections[2], "\n") {
		if i == 0 {
			continue
		}
		t, err := parseTicket(ticket)
		if err != nil {
			return nil, nil, nil, err
		}
		nearbyTickets = append(nearbyTickets, t)
	}

	return rules, myTicket, nearbyTickets, nil
}

func parseTicket(t string) ([]int, error) {
	ticket := []int{}
	for _, i := range strings.Split(t, ",") {
		j, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		ticket = append(ticket, j)
	}
	return ticket, nil
}
