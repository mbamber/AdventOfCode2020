package coordinate

import (
	"fmt"
	"log"
	"math"

	"github.com/golang/geo/s1"
)

// Cartesian represents a cartesian (x,y,z) coordinate pair
type Cartesian struct {
	X, Y, Z int
}

// NewCartesian returns a new 3 dimensional cartesian coordinate
func NewCartesian(x, y, z int) Cartesian {
	return Cartesian{
		X: x,
		Y: y,
		Z: z,
	}
}

// NewCartesian2D returns a new 2 dimensional cartesian coordinate
func NewCartesian2D(x, y int) Cartesian {
	return Cartesian{
		X: x,
		Y: y,
		Z: 0,
	}
}

// Adjacent returns all the coordinates that are exactly adjacent to the current coordinate
// i.e. those who have a manhatten distance of 1 to the current coordinate
func (c Cartesian) Adjacent() []Cartesian {
	adj := []Cartesian{
		NewCartesian(c.X+1, c.Y, c.Z),
		NewCartesian(c.X-1, c.Y, c.Z),

		NewCartesian(c.X, c.Y+1, c.Z),
		NewCartesian(c.X, c.Y-1, c.Z),

		NewCartesian(c.X, c.Y, c.Z+1),
		NewCartesian(c.X, c.Y, c.Z-1),
	}

	return adj
}

// Diagonal returns all the coordinates that are exactly diagonal to the current coordinate
// i.e. those who share a coordinate in one dimension, but differ by exactly one in the other
// two dimensions
func (c Cartesian) Diagonal() []Cartesian {
	diag := []Cartesian{
		NewCartesian(c.X+1, c.Y+1, c.Z),
		NewCartesian(c.X-1, c.Y+1, c.Z),
		NewCartesian(c.X+1, c.Y-1, c.Z),
		NewCartesian(c.X-1, c.Y-1, c.Z),

		NewCartesian(c.X+1, c.Y, c.Z+1),
		NewCartesian(c.X-1, c.Y, c.Z+1),
		NewCartesian(c.X+1, c.Y, c.Z-1),
		NewCartesian(c.X-1, c.Y, c.Z-1),

		NewCartesian(c.X, c.Y+1, c.Z+1),
		NewCartesian(c.X, c.Y-1, c.Z+1),
		NewCartesian(c.X, c.Y+1, c.Z-1),
		NewCartesian(c.X, c.Y-1, c.Z-1),
	}

	return diag
}

// AsCartesian returns the coordinate unchanged
func (c Cartesian) AsCartesian() Cartesian {
	return c
}

// EuclideanDistance returns the euclidean distance to `coord`
func (c Cartesian) EuclideanDistance(coord Coordinate) float64 {
	x := float64(c.X - coord.AsCartesian().X)
	y := float64(c.Y - coord.AsCartesian().Y)
	z := float64(c.Z - coord.AsCartesian().Z)

	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2) + math.Pow(z, 2))
}

// ManhattanDistance returns the Manhattan distance to `coord`
func (c Cartesian) ManhattanDistance(coord Coordinate) int {
	x := math.Abs(float64(c.X - coord.AsCartesian().X))
	y := math.Abs(float64(c.Y - coord.AsCartesian().Y))
	z := math.Abs(float64(c.Z - coord.AsCartesian().Z))

	return int(x + y + z)
}

// Rotate the coordinate about the origin by `angle` in a 2D (x,y) plane
func (c Cartesian) Rotate(angle s1.Angle) (Coordinate, error) {
	epsilon := 1e-13

	// Matrix multiplication
	sin, cos := math.Sincos(angle.Radians())
	x, y := float64(c.X), float64(c.Y)
	newX := cos*x + sin*y
	newY := -1*sin*x + cos*y // Note the -1 here to rotate clockwise

	// Due to floating point arithmetic, we need to check if the new values are within
	// an epsilon of their truncated values. Use `math.Round()` in case the error is negative
	if math.Abs(newX-math.Round(newX)) <= epsilon {
		newX = math.Round(newX)
	}
	if math.Abs(newY-math.Round(newY)) <= epsilon {
		newY = math.Round(newY)
	}

	if math.Trunc(newX) != newX || math.Trunc(newY) != newY {
		log.Printf("[ERROR] rotation leads to %.10f, %.10f\n", newX, newY)
		return Origin, fmt.Errorf("cannot rotate by %f degrees to a new integer coordinate", angle.Degrees())
	}

	return NewCartesian2D(int(newX), int(newY)), nil
}

// RotateAbout the given coordinate by `angle` in a 2D (x,y) plane
func (c Cartesian) RotateAbout(angle s1.Angle, about Coordinate) (Coordinate, error) {
	c.X, c.Y = c.X-about.AsCartesian().X, c.Y-about.AsCartesian().Y

	new, err := c.Rotate(angle)
	if err != nil {
		return Origin, err
	}

	newC := new.AsCartesian()
	newC.X, newC.Y = newC.X+about.AsCartesian().X, newC.Y+about.AsCartesian().Y

	return newC, nil
}
