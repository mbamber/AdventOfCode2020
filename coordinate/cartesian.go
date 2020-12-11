package coordinate

import "math"

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
