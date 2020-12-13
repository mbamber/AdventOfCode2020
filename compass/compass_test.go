package compass_test

import (
	"math"
	"testing"

	"github.com/golang/geo/s1"
	"github.com/mbamber/aoc/compass"
	"github.com/stretchr/testify/assert"
)

func TestTurn(t *testing.T) {
	cases := map[string]struct {
		start    compass.Bearing
		turn     float64
		expected compass.Bearing
	}{
		"turn -630 degrees": {
			start:    compass.North,
			turn:     -630,
			expected: compass.East,
		},
		"turn -90 degrees": {
			start:    compass.North,
			turn:     -90,
			expected: compass.West,
		},
		"turn 0 degrees": {
			start:    compass.North,
			turn:     0,
			expected: compass.North,
		},
		"turn 90 degrees": {
			start:    compass.North,
			turn:     90,
			expected: compass.East,
		},
		"turn 180 degrees": {
			start:    compass.North,
			turn:     180,
			expected: compass.South,
		},
		"turn 270 degrees": {
			start:    compass.North,
			turn:     270,
			expected: compass.West,
		},
		"turn 360 degrees": {
			start:    compass.North,
			turn:     360,
			expected: compass.North,
		},
		"turn 810 degrees": {
			start:    compass.North,
			turn:     810,
			expected: compass.East,
		},
		"turn 450 degrees": {
			start:    compass.North,
			turn:     450,
			expected: compass.East,
		},
		"turn 20 degrees": {
			start:    compass.North,
			turn:     20,
			expected: compass.Bearing(20),
		},
		"turn 180 degrees from west": {
			start:    compass.West,
			turn:     180,
			expected: compass.East,
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			angle := s1.Angle(data.turn / 360 * 2 * math.Pi)
			out := data.start.Turn(angle)
			assert.Equal(t, data.expected, out)
		})
	}
}
