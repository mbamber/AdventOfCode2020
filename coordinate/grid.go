package coordinate

// StringGrid is a grid of strings
type StringGrid map[Coordinate]string

// NewStringGrid creates an empty `StringGrid` with all values set to `null`
func NewStringGrid(rows, columns int, null string) StringGrid {
	grid := StringGrid{}

	for y := 0; y < rows; y = y + 1 {
		for x := 0; x < columns; x = x + 1 {
			grid[NewCartesian2D(x, y)] = null
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

	grid := NewStringGrid(len(rows), maxColumns, null)

	for y := 0; y < maxRows; y = y + 1 {
		for x := 0; x < len(rows[y]); x = x + 1 {
			grid.SetAt(x, y, string(rows[y][x]))
		}
	}
	return grid
}

// Get retrieves the value at the given coordinate
func (sg StringGrid) Get(c Coordinate) string {
	return sg[c]
}

// GetAt retrieves the value at the coordinate defined by `(x, y)`
func (sg StringGrid) GetAt(x, y int) string {
	return sg.Get(NewCartesian2D(x, y))
}

// Set updates the value at the given coordinate
func (sg StringGrid) Set(c Coordinate, v string) {
	sg[c] = v
}

// SetAt updates the value at the coordinate defined by `(x, y)`
func (sg StringGrid) SetAt(x, y int, v string) {
	sg.Set(NewCartesian2D(x, y), v)
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
func (sg StringGrid) Iterate(f func(x, y int, v string) (stop bool, err error)) error {
	for c, v := range sg {
		stop, err := f(c.AsCartesian().X, c.AsCartesian().Y, v)
		if err != nil {
			return err
		}

		if stop {
			break
		}
	}
	return nil
}

// Equal checks that two `StringGrid`s are exactly the same; that is, they have exactly the same dimensions
// and each cell contains the same value
func (sg StringGrid) Equal(other StringGrid) bool {
	sgRows, sgCols := sg.Dimensions()
	otherRows, otherCols := other.Dimensions()
	if sgRows != otherRows || sgCols != otherCols {
		return false
	}

	areEqual := true
	sg.Iterate(func(x, y int, v string) (stop bool, err error) {
		if sg.GetAt(x, y) != other.GetAt(x, y) {
			areEqual = false
			return true, nil
		}
		return false, nil
	})

	return areEqual
}

// Dimensions returns the numbers of rows and columns of the grid
func (sg StringGrid) Dimensions() (rows, columns int) {
	for c := range sg {
		if c.AsCartesian().X > columns {
			columns = c.AsCartesian().X
		}

		if c.AsCartesian().Y > rows {
			rows = c.AsCartesian().Y
		}
	}
	return rows + 1, columns + 1 // Add 1 because of 0 indexing
}

// String returns a printable version of the grid
func (sg StringGrid) String() string {
	rows, columns := sg.Dimensions()

	grid := ""
	for y := 0; y < rows; y = y + 1 {
		for x := 0; x < columns; x = x + 1 {
			grid = grid + sg.GetAt(x, y)
		}
		grid = grid + "\n"
	}

	return grid
}
