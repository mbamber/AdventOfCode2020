package main

import (
	"context"
	"math"
	"strconv"
	"strings"

	"github.com/golang/geo/s1"

	"github.com/mbamber/aoc/compass"
	"github.com/mbamber/aoc/coordinate"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	instructions := strings.Split(input, "\n")
	c := coordinate.NewCartesian2D(0, 0)
	facing := compass.East
	for _, instruction := range instructions {
		action, value := parseInstruction(instruction)
		c, facing = doAction(action, value, c, facing)
	}

	return c.ManhattanDistance(coordinate.Origin), nil
}

func doAction(action string, value int, c coordinate.Cartesian, facing compass.Bearing) (coordinate.Cartesian, compass.Bearing) {
	switch action {
	case "N":
		c.Y = c.Y + value
	case "E":
		c.X = c.X + value
	case "S":
		c.Y = c.Y - value
	case "W":
		c.X = c.X - value
	case "F":
		c, facing = doAction(facing.String(), value, c, facing)
	case "L":
		facing = facing.Turn(degreesToAngle(-1 * value))
	case "R":
		facing = facing.Turn(degreesToAngle(value))
	}
	return c, facing
}

func degreesToAngle(d int) s1.Angle {
	return s1.Angle(2 * math.Pi * float64(d) / 360.0)
}

func parseInstruction(s string) (string, int) {
	action := string(s[0])
	value, _ := strconv.Atoi(s[1:])
	return action, value
}
