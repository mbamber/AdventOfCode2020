package main

import (
	"context"
	"strings"

	"github.com/mbamber/aoc/coordinate"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	g1 := new4DSpace(strings.Split(input, "\n"), ".")
	g2 := new4DSpace(strings.Split(input, "\n"), ".")

	var curr, next map[coordinate4D]string
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			curr = g1
			next = g2
		} else {
			curr = g2
			next = g1
		}

		minX, maxX, minY, maxY, minZ, maxZ, minW, maxW := 0, 0, 0, 0, 0, 0, 0, 0
		for c := range curr {
			if c.X < minX {
				minX = c.X
			}
			if c.X > maxX {
				maxX = c.X
			}

			if c.Y < minY {
				minY = c.Y
			}
			if c.Y > maxY {
				maxY = c.Y
			}

			if c.Z < minZ {
				minZ = c.Z
			}
			if c.Z > maxZ {
				maxZ = c.Z
			}

			if c.W < minW {
				minW = c.W
			}
			if c.W > maxW {
				maxW = c.W
			}
		}
		maxX++
		maxY++
		maxZ++
		maxW++

		for x := minX - 1; x <= maxX; x++ {
			for y := minY - 1; y <= maxY; y++ {
				for z := minZ - 1; z <= maxZ; z++ {
					for w := minW - 1; w <= maxW; w++ {
						c := coordinate4D{X: x, Y: y, Z: z, W: w}
						v := curr[c]
						if v == "" {
							v = "."
						}

						activeCount := 0
						for _, s := range c.Surrounding() {
							switch curr[s] {
							case "#":
								activeCount++
							}
						}

						if v == "#" && (activeCount != 2 && activeCount != 3) {
							next[c] = "."
						} else if v == "." && activeCount == 3 {
							next[c] = "#"
						} else {
							next[c] = v
						}
					}
				}
			}
		}
	}

	// Count
	c := 0
	for _, v := range next {
		if v == "#" {
			c++
		}
	}

	return c, nil
}

type coordinate4D struct {
	X, Y, Z, W int
}

func (c coordinate4D) Surrounding() []coordinate4D {
	idx := []int{-1, 0, 1}
	surr := []coordinate4D{}

	for _, x := range idx {
		for _, y := range idx {
			for _, z := range idx {
				for _, w := range idx {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}

					surr = append(surr, coordinate4D{X: c.X + x, Y: c.Y + y, Z: c.Z + z, W: c.W + w})
				}
			}
		}
	}

	return surr
}

func new4DSpace(rows []string, null string) map[coordinate4D]string {
	space := map[coordinate4D]string{}
	g := coordinate.NewStringGridFromStrings(rows, null)
	for c, v := range g {
		c4D := coordinate4D{X: c.AsCartesian().X, Y: c.AsCartesian().Y, Z: c.AsCartesian().Z, W: 0}
		space[c4D] = v
	}
	return space
}
