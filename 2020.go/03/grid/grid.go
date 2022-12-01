package grid

import "errors"

// ScrollingGrid a grid that duplicates itself width-wise
// meaning that the only y axis bounds can return an OutOfBoundsError
type ScrollingGrid []string

var OutOfBoundsError = errors.New("OutOfBounds")

func (g ScrollingGrid) CellAt(x, y int) (rune, error) {
	var cell rune
	if y < 0 || y > len(g)-1 {
		return cell, OutOfBoundsError
	}

	cell = rune(g[y][x%len(g[y])])

	return cell, nil
}
