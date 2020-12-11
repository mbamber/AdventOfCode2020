package coordinate

// Origin represents the coordinate at (0, 0, 0)
var Origin Coordinate = NewCartesian(0, 0, 0)

// Coordinate is an interface that represents a coordinate in a variety of formats
type Coordinate interface {
	// AsCartesian returns the cartesian equivalent cartesian representation of the coordinate
	AsCartesian() Cartesian

	// EuclideanDistance returns the euclidean distance to `coord`
	EuclideanDistance(coord Coordinate) float64

	// ManhattanDistance returns the manhattan distance to `coord`
	ManhattanDistance(coord Coordinate) int
}
