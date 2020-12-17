package main

import (
	"context"
	"strings"

	"github.com/mbamber/aoc/coordinate"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	g1 := coordinate.NewStringGridFromStrings(strings.Split(input, "\n"), ".")
	g2 := coordinate.NewStringGridFromStrings(strings.Split(input, "\n"), ".")

	var curr, next coordinate.StringGrid
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			curr = g1
			next = g2
		} else {
			curr = g2
			next = g1
		}

		err := curr.IterateOutside(func(x, y, z int, v string) (stop bool, err error) {
			c := coordinate.NewCartesian(x, y, z)
			activeCount := 0
			for _, s := range c.Surrounding() {
				switch curr.Get(s) {
				case "#":
					activeCount++
				}
			}

			if v == "#" && (activeCount != 2 && activeCount != 3) {
				next.Set(c, ".")
			} else if v == "." && activeCount == 3 {
				next.Set(c, "#")
			} else {
				next.Set(c, v)
			}

			return false, nil
		}, ".")

		if err != nil {
			return nil, err
		}
	}

	return next.Count("#"), nil
}
