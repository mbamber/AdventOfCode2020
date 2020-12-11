package coordinate_test

import (
	"testing"

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
