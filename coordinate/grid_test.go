package coordinate_test

import (
	"testing"

	"github.com/mbamber/aoc/coordinate"
	"github.com/stretchr/testify/assert"
)

func TestNewStringGrid(t *testing.T) {
	cases := map[string]struct {
		columns      int
		rows         int
		depth        int
		null         string
		expectedSize int
	}{
		"6x5 grid": {
			columns:      6,
			rows:         5,
			depth:        1,
			null:         "x",
			expectedSize: 30,
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			grid := coordinate.NewStringGrid(data.rows, data.columns, data.depth, data.null)

			assert.Len(t, grid, data.expectedSize)
			for _, v := range grid {
				assert.Equal(t, data.null, v)
			}
		})
	}
}

func TestNewStringGridFromStrings(t *testing.T) {
	cases := map[string]struct {
		rows     []string
		null     string
		expected coordinate.StringGrid
	}{
		"square grid is created": {
			rows:     []string{"#####", "#####", "#####", "#####", "#####"},
			null:     ".",
			expected: coordinate.NewStringGrid(5, 5, 1, "#"),
		},
		"rectangle grid is created": {
			rows:     []string{"#####", "#####", "#####", "#####"},
			null:     ".",
			expected: coordinate.NewStringGrid(4, 5, 1, "#"),
		},
		"incomplete grid is created and padded": {
			rows: []string{"####", "#####", "#####", "#####"},
			null: ".",
			expected: func() coordinate.StringGrid {
				g := coordinate.NewStringGrid(4, 5, 1, "#")
				g.SetAt(4, 0, 0, ".")
				return g
			}(),
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			assert.True(t, data.expected.Equal(coordinate.NewStringGridFromStrings(data.rows, data.null)))
		})
	}
}

func TestMinDimensions(t *testing.T) {
	cases := map[string]struct {
		columns int
		rows    int
		depth   int
		null    string
	}{
		"6x5x4 grid": {
			columns: 6,
			rows:    5,
			depth:   4,
			null:    "x",
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			grid := coordinate.NewStringGrid(data.rows, data.columns, data.depth, data.null)

			rows, columns, depth := grid.MinDimensions()

			assert.Equal(t, 0, rows)
			assert.Equal(t, 0, columns)
			assert.Equal(t, 0, depth)
		})
	}
}

func TestDimensions(t *testing.T) {
	cases := map[string]struct {
		columns int
		rows    int
		depth   int
		null    string
	}{
		"6x5x4 grid": {
			columns: 6,
			rows:    5,
			depth:   4,
			null:    "x",
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			grid := coordinate.NewStringGrid(data.rows, data.columns, data.depth, data.null)

			rows, columns, depth := grid.Dimensions()

			assert.Equal(t, data.rows, rows)
			assert.Equal(t, data.columns, columns)
			assert.Equal(t, data.depth, depth)
		})
	}
}

func TestString(t *testing.T) {
	cases := map[string]struct {
		grid     coordinate.StringGrid
		expected string
	}{
		"6x5x1 grid": {
			grid:     coordinate.NewStringGrid(5, 6, 1, "x"),
			expected: "z: 0\n |\n-xxxxxx\n xxxxxx\n xxxxxx\n xxxxxx\n xxxxxx\n\n",
		},
		"3x3x3 grid": {
			grid: coordinate.StringGrid{
				coordinate.NewCartesian(-1, -1, -1): "x",
				coordinate.NewCartesian(-1, -1, 0):  "x",
				coordinate.NewCartesian(-1, -1, 1):  "x",
				coordinate.NewCartesian(-1, 0, -1):  "x",
				coordinate.NewCartesian(-1, 0, 0):   "x",
				coordinate.NewCartesian(-1, 0, 1):   "x",
				coordinate.NewCartesian(-1, 1, -1):  "x",
				coordinate.NewCartesian(-1, 1, 0):   "x",
				coordinate.NewCartesian(-1, 1, 1):   "x",
				coordinate.NewCartesian(0, -1, -1):  "x",
				coordinate.NewCartesian(0, -1, 0):   "x",
				coordinate.NewCartesian(0, -1, 1):   "x",
				coordinate.NewCartesian(0, 0, -1):   "x",
				coordinate.NewCartesian(0, 0, 0):    "x",
				coordinate.NewCartesian(0, 0, 1):    "x",
				coordinate.NewCartesian(0, 1, -1):   "x",
				coordinate.NewCartesian(0, 1, 0):    "x",
				coordinate.NewCartesian(0, 1, 1):    "x",
				coordinate.NewCartesian(1, -1, -1):  "x",
				coordinate.NewCartesian(1, -1, 0):   "x",
				coordinate.NewCartesian(1, -1, 1):   "x",
				coordinate.NewCartesian(1, 0, -1):   "x",
				coordinate.NewCartesian(1, 0, 0):    "x",
				coordinate.NewCartesian(1, 0, 1):    "x",
				coordinate.NewCartesian(1, 1, -1):   "x",
				coordinate.NewCartesian(1, 1, 0):    "x",
				coordinate.NewCartesian(1, 1, 1):    "x",
			},
			expected: "z: -1\n  |\n xxx\n-xxx\n xxx\n\nz: 0\n  |\n xxx\n-xxx\n xxx\n\nz: 1\n  |\n xxx\n-xxx\n xxx\n\n",
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, data.grid.String())
		})
	}
}

func TestEqual(t *testing.T) {
	cases := map[string]struct {
		a        coordinate.StringGrid
		b        coordinate.StringGrid
		expected bool
	}{
		"empty grids are equal": {
			a:        coordinate.NewStringGrid(10, 10, 10, "."),
			b:        coordinate.NewStringGrid(10, 10, 10, "."),
			expected: true,
		},
		"different grids are not equal": {
			a:        coordinate.NewStringGrid(10, 10, 10, "."),
			b:        coordinate.NewStringGrid(10, 10, 10, "#"),
			expected: false,
		},
		"different dimenstions are not equal": {
			a:        coordinate.NewStringGrid(10, 5, 3, "."),
			b:        coordinate.NewStringGrid(5, 10, 3, "."),
			expected: false,
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, data.a.Equal(data.b))
			assert.Equal(t, data.expected, data.b.Equal(data.a))
		})
	}
}

func TestCount(t *testing.T) {
	cases := map[string]struct {
		g        coordinate.StringGrid
		toCount  string
		expected int
	}{
		"all cells are the same": {
			g:        coordinate.NewStringGrid(5, 5, 5, "#"),
			toCount:  "#",
			expected: 125,
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, data.g.Count(data.toCount))
		})
	}
}

func TestIterateOutside(t *testing.T) {
	sg := coordinate.NewStringGrid(3, 3, 3, "#")
	count := 0
	err := sg.IterateOutside(func(x, y, z int, v string) (bool, error) {
		count++
		return false, nil
	}, "#")

	assert.NoError(t, err)
	assert.Equal(t, 125, count)
}
