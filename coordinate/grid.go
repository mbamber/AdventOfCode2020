package coordinate

import "fmt"

// StringGrid is a grid of strings
type StringGrid map[Coordinate]string

// NewStringGrid creates an empty `StringGrid` with all values set to `null`
func NewStringGrid(rows, columns, depth int, null string) StringGrid {
	grid := StringGrid{}

	for z := 0; z < depth; z = z + 1 {
		for y := 0; y < rows; y = y + 1 {
			for x := 0; x < columns; x = x + 1 {
				grid[NewCartesian(x, y, z)] = null
			}
		}
	}

	return grid
}

// NewStringGridFromStrings loads a slice of strings as rows in a `StringGrid`
func NewStringGridFromStrings(rows []string, null string) StringGrid {
	maxColumns := 0
	maxRows := len(rows)
	for _, row := range rows {
		if len(row) > maxColumns {
			maxColumns = len(row)
		}
	}

	grid := NewStringGrid(len(rows), maxColumns, 1, null)

	for y := 0; y < maxRows; y = y + 1 {
		for x := 0; x < len(rows[y]); x = x + 1 {
			grid.SetAt(x, y, 0, string(rows[y][x]))
		}
	}
	return grid
}

// Get retrieves the value at the given coordinate
func (sg StringGrid) Get(c Coordinate) string {
	return sg[c]
}

// GetAt retrieves the value at the coordinate defined by `(x, y)`
func (sg StringGrid) GetAt(x, y, z int) string {
	return sg.Get(NewCartesian(x, y, z))
}

// Set updates the value at the given coordinate
func (sg StringGrid) Set(c Coordinate, v string) {
	sg[c] = v
}

// SetAt updates the value at the coordinate defined by `(x, y)`
func (sg StringGrid) SetAt(x, y, z int, v string) {
	sg.Set(NewCartesian(x, y, z), v)
}

// Count counts the number of times `v` appears in the grid
func (sg StringGrid) Count(v string) int {
	c := 0
	for _, val := range sg {
		if v == val {
			c = c + 1
		}
	}
	return c
}

// Iterate calls `f` for each cell in the grid. Return `true` as the first argument from `f` to
// halt iteration
func (sg StringGrid) Iterate(f func(x, y, z int, v string) (stop bool, err error)) error {
	for c, v := range sg {
		stop, err := f(c.AsCartesian().X, c.AsCartesian().Y, c.AsCartesian().Z, v)
		if err != nil {
			return err
		}

		if stop {
			break
		}
	}
	return nil
}

// IterateOutside calls `f` for each cell in the grid, and each cell surrounding the grid. Return `true`
// as the first argument from `f` to halt iteration
func (sg StringGrid) IterateOutside(f func(x, y, z int, v string) (stop bool, err error), null string) error {
	maxY, maxX, maxZ := sg.Dimensions()
	minY, minX, minZ := sg.MinDimensions()

	for x := minX - 1; x <= maxX; x++ {
		for y := minY - 1; y <= maxY; y++ {
			for z := minZ - 1; z <= maxZ; z++ {
				c := NewCartesian(x, y, z)
				v := sg.Get(c)
				if v == "" {
					v = null
				}

				stop, err := f(x, y, z, v)
				if err != nil {
					return err
				}
				if stop {
					return nil
				}
			}
		}
	}

	return nil
}

// Equal checks that two `StringGrid`s are exactly the same; that is, they have exactly the same dimensions
// and each cell contains the same value
func (sg StringGrid) Equal(other StringGrid) bool {
	sgRows, sgCols, sgDepth := sg.Dimensions()
	otherRows, otherCols, otherDepth := other.Dimensions()
	if sgRows != otherRows || sgCols != otherCols || sgDepth != otherDepth {
		return false
	}

	areEqual := true
	sg.Iterate(func(x, y, z int, v string) (stop bool, err error) {
		if sg.GetAt(x, y, z) != other.GetAt(x, y, z) {
			areEqual = false
			return true, nil
		}
		return false, nil
	})

	return areEqual
}

// MinDimensions returns the lowest row, column and depth of the grid
func (sg StringGrid) MinDimensions() (rows, columns, depth int) {
	for c := range sg {
		if c.AsCartesian().X < columns {
			columns = c.AsCartesian().X
		}

		if c.AsCartesian().Y < rows {
			rows = c.AsCartesian().Y
		}

		if c.AsCartesian().Z < depth {
			depth = c.AsCartesian().Z
		}
	}
	return rows, columns, depth
}

// Dimensions returns the highest row, column and depth of the grid
func (sg StringGrid) Dimensions() (rows, columns, depth int) {
	for c := range sg {
		if c.AsCartesian().X > columns {
			columns = c.AsCartesian().X
		}

		if c.AsCartesian().Y > rows {
			rows = c.AsCartesian().Y
		}

		if c.AsCartesian().Z > depth {
			depth = c.AsCartesian().Z
		}
	}
	return rows + 1, columns + 1, depth + 1 // Add 1 because of 0 indexing
}

// String returns a printable version of the grid with z coordinates equal to 0
func (sg StringGrid) String() string {
	minRows, minColumns, minDepth := sg.MinDimensions()
	rows, columns, depth := sg.Dimensions()

	grid := ""
	for z := minDepth; z < depth; z++ {
		grid += fmt.Sprintf("z: %d\n", z)
		for i := minRows; i < 0; i++ {
			grid += " "
		}
		grid += " |\n"
		for y := minRows; y < rows; y++ {
			if y == 0 {
				grid += "-"
			} else {
				grid += " "
			}

			for x := minColumns; x < columns; x++ {
				grid = grid + sg.GetAt(x, y, z)
			}
			grid += "\n"
		}
		grid += "\n"
	}

	return grid
}
