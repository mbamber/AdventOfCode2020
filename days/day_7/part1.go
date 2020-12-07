package main

import (
	"context"
	"regexp"
	"strconv"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	bags, err := buildTree(input)
	if err != nil {
		return nil, err
	}

	shinyGold := bagSliceContains(bags, "shiny gold")
	rootBags := shinyGold.containedBy(bags, []*bag{})

	return len(rootBags), nil
}

func buildTree(input string) ([]*bag, error) {
	innerRegexp := regexp.MustCompile(`(?:(\d+.*?) bags?[,\.])`)

	bags := []*bag{}

	// Do an initial pass to create all the initial bags
	for _, rule := range strings.Split(input, "\n") {
		ruleParts := strings.Split(rule, " bags contain ")
		color := ruleParts[0]

		b := &bag{
			color: color,
		}
		bags = append(bags, b)
	}

	// Do a second pass to add all the sub bags
	for _, rule := range strings.Split(input, "\n") {
		ruleParts := strings.Split(rule, " bags contain ")
		b := bagSliceContains(bags, ruleParts[0])

		contents := innerRegexp.FindAllStringSubmatch(ruleParts[1], -1)
		for _, content := range contents {
			details := strings.SplitN(content[1], " ", 2)
			count, err := strconv.Atoi(details[0])
			if err != nil {
				return nil, err
			}
			color := details[1]

			subBag := bagSliceContains(bags, color)
			if subBag == nil {
				subBag = &bag{
					color: color,
				}
				bags = append(bags, subBag)
			}

			for i := 0; i < count; i = i + 1 {
				b.contains = append(b.contains, subBag)
			}
		}
	}

	return bags, nil
}

type bag struct {
	color    string
	contains []*bag
}

func (b *bag) containedBy(bags []*bag, currBags []*bag) []*bag {
	// 1.   Find all of the bags that contain b as a sub bag
	// 2.   For each of those bags:
	// 2.1 		Add the bag to the list of bags (if it doesn't already exist)
	// 2.2 		Call recursively for that bag

	for _, subBag := range bags {
		contains := bagSliceContains(subBag.contains, b.color)
		if contains != nil {
			alreadyTracked := bagSliceContains(currBags, subBag.color)
			if alreadyTracked == nil {
				currBags = append(currBags, subBag)
			}

			bagsToAdd := subBag.containedBy(bags, currBags)
			for _, toAdd := range bagsToAdd {
				alreadyTracked := bagSliceContains(currBags, toAdd.color)
				if alreadyTracked == nil {
					currBags = append(currBags, toAdd)
				}
			}
		}
	}

	return currBags
}

func bagSliceContains(s []*bag, color string) *bag {
	for _, b := range s {
		if b.color == color {
			return b
		}
	}
	return nil
}
