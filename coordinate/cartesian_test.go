package coordinate_test

import (
	"errors"
	"math"
	"testing"

	"github.com/golang/geo/s1"
	"github.com/mbamber/aoc/coordinate"
	"github.com/stretchr/testify/assert"
)

func TestAdjacent(t *testing.T) {
	cases := map[string]struct {
		c        coordinate.Cartesian
		expected []coordinate.Cartesian
	}{
		"adjacent to (0, 0, 0)": {
			c: coordinate.NewCartesian(0, 0, 0),
			expected: []coordinate.Cartesian{
				coordinate.NewCartesian(1, 0, 0),
				coordinate.NewCartesian(-1, 0, 0),
				coordinate.NewCartesian(0, 1, 0),
				coordinate.NewCartesian(0, -1, 0),
				coordinate.NewCartesian(0, 0, 1),
				coordinate.NewCartesian(0, 0, -1),
			},
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, data.c.Adjacent())
		})
	}
}

func TestDiagonal(t *testing.T) {
	cases := map[string]struct {
		c        coordinate.Cartesian
		expected []coordinate.Cartesian
	}{
		"diagonal to (0, 0, 0)": {
			c: coordinate.NewCartesian(0, 0, 0),
			expected: []coordinate.Cartesian{
				coordinate.NewCartesian(1, 1, 0),
				coordinate.NewCartesian(-1, 1, 0),
				coordinate.NewCartesian(1, -1, 0),
				coordinate.NewCartesian(-1, -1, 0),
				coordinate.NewCartesian(1, 0, 1),
				coordinate.NewCartesian(-1, 0, 1),
				coordinate.NewCartesian(1, 0, -1),
				coordinate.NewCartesian(-1, 0, -1),
				coordinate.NewCartesian(0, 1, 1),
				coordinate.NewCartesian(0, -1, 1),
				coordinate.NewCartesian(0, 1, -1),
				coordinate.NewCartesian(0, -1, -1),
			},
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, data.c.Diagonal())
		})
	}
}

func TestRotate(t *testing.T) {
	cases := map[string]struct {
		c           coordinate.Coordinate
		rotateBy    float64
		expected    coordinate.Coordinate
		expectedErr error
	}{
		"rotate 90 degrees": {
			c:           coordinate.NewCartesian2D(1, 0),
			rotateBy:    90,
			expected:    coordinate.NewCartesian2D(0, -1),
			expectedErr: nil,
		},
		"rotate 180 degrees": {
			c:           coordinate.NewCartesian2D(1, 0),
			rotateBy:    180,
			expected:    coordinate.NewCartesian2D(-1, 0),
			expectedErr: nil,
		},
		"rotate 270 degrees": {
			c:           coordinate.NewCartesian2D(1, 0),
			rotateBy:    270,
			expected:    coordinate.NewCartesian2D(0, 1),
			expectedErr: nil,
		},
		"rotate 10 degrees causes error": {
			c:           coordinate.NewCartesian2D(1, 0),
			rotateBy:    10,
			expected:    coordinate.NewCartesian2D(0, 0),
			expectedErr: errors.New("cannot rotate by 10.000000 degrees to a new integer coordinate"),
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			a := s1.Angle(2 * math.Pi * data.rotateBy / 360)
			out, err := data.c.Rotate(a)
			if data.expectedErr != nil {
				assert.Equal(t, data.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, data.expected, out)
			}
		})
	}
}

func TestRotateAbout(t *testing.T) {
	cases := map[string]struct {
		c           coordinate.Coordinate
		about       coordinate.Coordinate
		rotateBy    float64
		expected    coordinate.Coordinate
		expectedErr error
	}{
		"rotate 90 degrees": {
			c:           coordinate.NewCartesian2D(2, 1),
			about:       coordinate.NewCartesian2D(1, 1),
			rotateBy:    90,
			expected:    coordinate.NewCartesian2D(1, 0),
			expectedErr: nil,
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			a := s1.Angle(2 * math.Pi * data.rotateBy / 360)
			out, err := data.c.RotateAbout(a, data.about)
			if data.expectedErr != nil {
				assert.Equal(t, data.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, data.expected, out)
			}
		})
	}
}

func TestSurrounding(t *testing.T) {
	cases := map[string]struct {
		c        coordinate.Cartesian
		expected []coordinate.Cartesian
	}{
		"points surrounding origin": {
			c: coordinate.Origin.AsCartesian(),
			expected: []coordinate.Cartesian{
				coordinate.NewCartesian(-1, -1, -1),
				coordinate.NewCartesian(-1, -1, 0),
				coordinate.NewCartesian(-1, -1, 1),
				coordinate.NewCartesian(-1, 0, -1),
				coordinate.NewCartesian(-1, 0, 0),
				coordinate.NewCartesian(-1, 0, 1),
				coordinate.NewCartesian(-1, 1, -1),
				coordinate.NewCartesian(-1, 1, 0),
				coordinate.NewCartesian(-1, 1, 1),
				coordinate.NewCartesian(0, -1, -1),
				coordinate.NewCartesian(0, -1, 0),
				coordinate.NewCartesian(0, -1, 1),
				coordinate.NewCartesian(0, 0, -1),
				coordinate.NewCartesian(0, 0, 1),
				coordinate.NewCartesian(0, 1, -1),
				coordinate.NewCartesian(0, 1, 0),
				coordinate.NewCartesian(0, 1, 1),
				coordinate.NewCartesian(1, -1, -1),
				coordinate.NewCartesian(1, -1, 0),
				coordinate.NewCartesian(1, -1, 1),
				coordinate.NewCartesian(1, 0, -1),
				coordinate.NewCartesian(1, 0, 0),
				coordinate.NewCartesian(1, 0, 1),
				coordinate.NewCartesian(1, 1, -1),
				coordinate.NewCartesian(1, 1, 0),
				coordinate.NewCartesian(1, 1, 1),
			},
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			s := data.c.Surrounding()
			assert.Len(t, s, len(data.expected))
			assert.Equal(t, data.expected, s)
		})
	}
}
