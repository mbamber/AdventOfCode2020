package main

import (
	"context"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	rules, myTicket, nearbyTickets, err := parseInput(input)
	if err != nil {
		return nil, err
	}

	validTickets := [][]int{}
	for _, ticket := range nearbyTickets {
		_, invalid := ticketIsValid(ticket, rules)
		if len(invalid) == 0 {
			validTickets = append(validTickets, ticket)
		}
	}

	// For each rule, try to identify the possible indexes
	indexes := map[int][]rule{}
	for _, r := range rules {
		for i := 0; i < len(validTickets[0]); i = i + 1 {
			indexFound := true
			for _, ticket := range validTickets {
				if !r.contains(ticket[i]) {
					indexFound = false
					break
				}
			}

			if indexFound {
				if indexes[i] == nil {
					indexes[i] = []rule{r}
				} else {
					indexes[i] = append(indexes[i], r)
				}
			}
		}
	}

	actualIndexes := reduceIndexes(indexes)

	total := 1
	for index, r := range actualIndexes {
		if strings.HasPrefix(r.name, "departure") {
			total = total * myTicket[index]
		}
	}

	return total, nil
}

func reduceIndexes(indexes map[int][]rule) map[int]rule {
	actualIndexes := map[int]rule{}

	for !indexesAreReduced(indexes) {
		for idx, rules := range indexes {
			if len(rules) == 1 {
				if _, ok := actualIndexes[idx]; !ok {
					actualIndexes[idx] = rules[0]
					indexes = removeRuleFromMap(indexes, rules[0], idx)
				}
			}
		}
	}

	return actualIndexes
}

func removeRuleFromMap(indexes map[int][]rule, r rule, exceptAt int) map[int][]rule {
	newIndexes := map[int][]rule{}

	for idx, rules := range indexes {
		if idx == exceptAt {
			continue
		}

		newIndexes[idx] = removeRuleFromSlice(rules, r)
	}

	return newIndexes
}

func removeRuleFromSlice(s []rule, r rule) []rule {
	newS := []rule{}

	for _, r1 := range s {
		if r1.name == r.name {
			continue
		}
		newS = append(newS, r1)
	}

	return newS
}

func indexesAreReduced(indexes map[int][]rule) bool {
	for _, rules := range indexes {
		if len(rules) != 1 {
			return false
		}
	}
	return true
}
