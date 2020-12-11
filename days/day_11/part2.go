package main

import (
	"context"
	"errors"
	"strings"

	"github.com/mbamber/aoc/coordinate"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	lines := strings.Split(input, "\n")
	old := coordinate.NewStringGridFromStrings(lines, ".")
	new := coordinate.NewStringGridFromStrings(lines, ".")

	rows, cols := old.Dimensions()

	// Build a list of all adjacent seats
	seatsToCheck := map[coordinate.Cartesian][]coordinate.Cartesian{}
	old.Iterate(func(x, y int, _ string) (stop bool, err error) {
		c := coordinate.NewCartesian2D(x, y)
		dxdy := []int{-1, 0, 1}
		seats := []coordinate.Cartesian{}

		for _, dy := range dxdy {
			for _, dx := range dxdy {
				if dx == 0 && dy == 0 {
					continue
				}

				s, err := getFirstSeat(old, c, dx, dy, rows, cols)
				if err == nil {
					seats = append(seats, s)
				}
			}
		}

		seatsToCheck[c] = seats
		return false, nil
	})

	for {
		old.Iterate(func(x, y int, _ string) (stop bool, err error) {
			c := coordinate.NewCartesian2D(x, y)
			if old.Get(c) != "." {
				updateSeat(old, new, c, seatsToCheck[c], rows, cols, 5)
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

func getFirstSeat(g coordinate.StringGrid, c coordinate.Cartesian, dx, dy, maxRows, maxCols int) (coordinate.Cartesian, error) {
	m := 1

	for {
		y := c.Y + m*dy
		x := c.X + m*dx
		if x < 0 || x >= maxCols || y < 0 || y >= maxRows {
			return coordinate.Origin.AsCartesian(), errors.New("out of range")
		}

		if g.GetAt(x, y) != "." {
			return coordinate.NewCartesian2D(x, y), nil
		}

		m = m + 1
	}
}
