package main

import (
	"context"
	"strings"

	"github.com/mbamber/aoc/coordinate"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	lines := strings.Split(input, "\n")
	old := coordinate.NewStringGridFromStrings(lines, ".")
	new := coordinate.NewStringGridFromStrings(lines, ".")

	// Build a list of all adjacent seats
	seatsToCheck := map[coordinate.Cartesian][]coordinate.Cartesian{}
	old.Iterate(func(x, y int, _ string) (stop bool, err error) {
		c := coordinate.NewCartesian2D(x, y)
		s := c.Diagonal()
		s = append(s, c.Adjacent()...)

		uniqueSeats := []coordinate.Cartesian{}
		for _, coord := range s {
			exists := false
			for _, seat := range uniqueSeats {
				if seat == coord {
					exists = true
					break
				}
			}
			if !exists {
				uniqueSeats = append(uniqueSeats, coord)
			}
		}

		seatsToCheck[c] = uniqueSeats
		return false, nil
	})

	rows, cols := old.Dimensions()
	for {
		old.Iterate(func(x, y int, _ string) (stop bool, err error) {
			c := coordinate.NewCartesian2D(x, y)
			if old.Get(c) != "." {
				updateSeat(old, new, c, seatsToCheck[c], rows, cols, 4)
			}
			return false, nil
		})

		if old.Equal(new) {
			break
		}
		old, new = new, old
	}

	return old.Count("#"), nil
}

func updateSeat(old, new coordinate.StringGrid, c coordinate.Cartesian, seatsToSearch []coordinate.Cartesian, maxRows, maxColumns, maxOccupied int) {
	seatsUsed := 0

	for _, adj := range seatsToSearch {
		if adj.X < 0 || adj.X >= maxColumns || adj.Y < 0 || adj.Y >= maxRows {
			continue
		}

		if old.Get(adj) == "#" {
			seatsUsed = seatsUsed + 1
		}

		if seatsUsed >= maxOccupied {
			new.Set(c, "L")
			return
		}
	}

	if seatsUsed == 0 {
		new.Set(c, "#")
	} else {
		new.Set(c, old.Get(c))
	}
}
