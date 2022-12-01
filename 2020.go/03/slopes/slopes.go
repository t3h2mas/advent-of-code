package slopes

import "github.com/t3h2mas/advent-2020/03/grid"

type TobogganDirection struct {
	forward int
	down    int
}

func NewTobogganDirection(forwardRate, downwardRate int) *TobogganDirection {
	return &TobogganDirection{
		forwardRate,
		downwardRate,
	}
}

func PathEncounters(grid grid.ScrollingGrid, direction *TobogganDirection) map[rune]int {
	x := 0
	y := 0

	results := make(map[rune]int)

	for {
		cell, err := grid.CellAt(x, y)

		if err != nil {
			break
		}

		cellCount := results[cell]

		results[cell] = cellCount + 1

		x += direction.forward
		y += direction.down
	}

	return results
}
