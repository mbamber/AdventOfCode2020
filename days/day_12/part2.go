package main

import (
	"context"
	"strings"

	"github.com/mbamber/aoc/coordinate"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	instructions := strings.Split(input, "\n")
	ship := coordinate.NewCartesian2D(0, 0)
	waypoint := coordinate.NewCartesian2D(10, 1)

	for _, instruction := range instructions {
		action, value := parseInstruction(instruction)

		switch action {
		case "N":
			waypoint.Y = waypoint.Y + value
		case "E":
			waypoint.X = waypoint.X + value
		case "S":
			waypoint.Y = waypoint.Y - value
		case "W":
			waypoint.X = waypoint.X - value
		case "L":
			w, err := waypoint.Rotate(degreesToAngle(-1 * value))
			if err != nil {
				return nil, err
			}
			waypoint = w.AsCartesian()
		case "R":
			w, err := waypoint.Rotate(degreesToAngle(value))
			if err != nil {
				return nil, err
			}
			waypoint = w.AsCartesian()
		case "F":
			ship.Y = ship.Y + (value * waypoint.Y)
			ship.X = ship.X + (value * waypoint.X)
		}

	}

	return ship.ManhattanDistance(coordinate.Origin), nil
}
