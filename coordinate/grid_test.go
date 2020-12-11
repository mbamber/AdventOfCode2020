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
		null         string
		expectedSize int
	}{
		"6x5 grid": {
			columns:      6,
			rows:         5,
			null:         "x",
			expectedSize: 30,
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			grid := coordinate.NewStringGrid(data.rows, data.columns, data.null)

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
			expected: coordinate.NewStringGrid(5, 5, "#"),
		},
		"rectangle grid is created": {
			rows:     []string{"#####", "#####", "#####", "#####"},
			null:     ".",
			expected: coordinate.NewStringGrid(4, 5, "#"),
		},
		"incomplete grid is created and padded": {
			rows: []string{"####", "#####", "#####", "#####"},
			null: ".",
			expected: func() coordinate.StringGrid {
				g := coordinate.NewStringGrid(4, 5, "#")
				g.SetAt(4, 0, ".")
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

func TestDimensions(t *testing.T) {
	cases := map[string]struct {
		columns int
		rows    int
		null    string
	}{
		"6x5 grid": {
			columns: 6,
			rows:    5,
			null:    "x",
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			grid := coordinate.NewStringGrid(data.rows, data.columns, data.null)

			rows, columns := grid.Dimensions()

			assert.Equal(t, data.rows, rows)
			assert.Equal(t, data.columns, columns)
		})
	}
}

func TestString(t *testing.T) {
	cases := map[string]struct {
		columns  int
		rows     int
		null     string
		expected string
	}{
		"6x5 grid": {
			columns:  6,
			rows:     5,
			null:     "x",
			expected: "xxxxxx\nxxxxxx\nxxxxxx\nxxxxxx\nxxxxxx\n",
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			grid := coordinate.NewStringGrid(data.rows, data.columns, data.null)
			assert.Equal(t, data.expected, grid.String())
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
			a:        coordinate.NewStringGrid(10, 10, "."),
			b:        coordinate.NewStringGrid(10, 10, "."),
			expected: true,
		},
		"different grids are not equal": {
			a:        coordinate.NewStringGrid(10, 10, "."),
			b:        coordinate.NewStringGrid(10, 10, "#"),
			expected: false,
		},
		"different dimenstions are not equal": {
			a:        coordinate.NewStringGrid(10, 5, "."),
			b:        coordinate.NewStringGrid(5, 10, "."),
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
			g:        coordinate.NewStringGrid(5, 5, "#"),
			toCount:  "#",
			expected: 25,
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, data.expected, data.g.Count(data.toCount))
		})
	}
}
