package seats

type Layout struct {
	height   int
	width    int
	occupied PointSet
	floor    PointSet
}

func newLayout(height int, width int) Layout {
	return Layout{
		height:   height,
		width:    width,
		occupied: NewPointSet(),
		floor:    NewPointSet(),
	}
}

func (l Layout) setHeight(val int) {
	l.height = val
}

func (l Layout) setWidth(val int) {
	l.width = val
}

func LayoutFrom(grid []string) Layout {
	result := newLayout(len(grid), len(grid[0]))

	for y := range grid {
		for x := range grid[y] {
			switch grid[y][x] {
			case '.':
				p := NewPoint(x, y)
				result.floor.Add(p)
			case '#':
				p := NewPoint(x, y)
				result.occupied.Add(p)
			}
		}
	}

	return result
}

func (l Layout) SeatOutOfBounds(seat Point) bool {
	if seat.X() < 0 || seat.X() >= l.width {
		return true
	}

	if seat.Y() < 0 || seat.Y() >= l.height {
		return true
	}
	return false
}
