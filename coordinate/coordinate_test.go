package coordinate_test

import (
	"testing"

	"github.com/mbamber/aoc/coordinate"
	"github.com/stretchr/testify/assert"
)

func TestAsCartesian(t *testing.T) {
	cases := map[string]struct {
		coordinate coordinate.Coordinate
		expected   coordinate.Coordinate
	}{
		"2D coordinate is unchanged": {
			coordinate: coordinate.NewCartesian2D(1, 2),
			expected:   coordinate.NewCartesian2D(1, 2),
		},
		"3D coordiate is unchanged": {
			coordinate: coordinate.NewCartesian(1, 2, 3),
			expected:   coordinate.NewCartesian(1, 2, 3),
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			out := data.coordinate.AsCartesian()
			assert.Equal(t, data.expected, out)
		})
	}
}

func TestEuclideanDistance(t *testing.T) {
	cases := map[string]struct {
		c1       coordinate.Coordinate
		c2       coordinate.Coordinate
		expected float64
	}{
		"3D distance to origin": {
			c1:       coordinate.NewCartesian(1, 2, 2),
			c2:       coordinate.Origin,
			expected: 3,
		},
		"2D distance to origin": {
			c1:       coordinate.NewCartesian2D(3, 4),
			c2:       coordinate.Origin,
			expected: 5,
		},
		"3D distance between points": {
			c1:       coordinate.NewCartesian(2, 3, 3),
			c2:       coordinate.NewCartesian(1, 1, 1),
			expected: 3,
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			out1 := data.c1.EuclideanDistance(data.c2)
			out2 := data.c2.EuclideanDistance(data.c1)

			assert.Equal(t, data.expected, out1)
			assert.Equal(t, data.expected, out2)
		})
	}
}

func TestManhattanDistance(t *testing.T) {
	cases := map[string]struct {
		c1       coordinate.Coordinate
		c2       coordinate.Coordinate
		expected int
	}{
		"3D distance to origin": {
			c1:       coordinate.NewCartesian(1, 2, 3),
			c2:       coordinate.Origin,
			expected: 6,
		},
		"2D distance to origin": {
			c1:       coordinate.NewCartesian2D(1, 2),
			c2:       coordinate.Origin,
			expected: 3,
		},
		"3D distance between points": {
			c1:       coordinate.NewCartesian(1, 2, 3),
			c2:       coordinate.NewCartesian(4, 5, 6),
			expected: 9,
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			out1 := data.c1.ManhattanDistance(data.c2)
			out2 := data.c2.ManhattanDistance(data.c1)

			assert.Equal(t, data.expected, out1)
			assert.Equal(t, data.expected, out2)
		})
	}
}
