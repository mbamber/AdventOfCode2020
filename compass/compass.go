package compass

import (
	"fmt"
	"math"

	"github.com/golang/geo/s1"
)

// Bearing is a compass bearing
type Bearing float64

const (
	// North represents due north
	North Bearing = 0.0
	// East represents due East
	East Bearing = 90.0
	// South represents due South
	South Bearing = 180.0
	// West represents due west
	West Bearing = 270.0
)

// Normalized returns a normalized version of the bearing, 0 <= b < 360
func (b Bearing) Normalized() Bearing {
	normalizedDegrees := math.Mod(float64(b), 360)
	if normalizedDegrees < 0 {
		normalizedDegrees = normalizedDegrees + 360
	}
	return Bearing(normalizedDegrees)
}

// Turn returns a new direction by turning through `degrees` clockwise
func (b Bearing) Turn(angle s1.Angle) Bearing {
	newBearing := Bearing(float64(b) + angle.Degrees())
	return newBearing.Normalized()
}

func (b Bearing) String() string {
	switch b {
	case North:
		return "N"
	case East:
		return "E"
	case South:
		return "S"
	case West:
		return "W"
	default:
		return fmt.Sprintf("%f", float64(b))
	}
}
