package hyperlife

type hyperCube struct {
	x int
	y int
	z int
	w int
}

func newHyperCube(x, y, z, w int) hyperCube {
	return hyperCube{x, y, z, w}
}

func (c hyperCube) Add(x, y, z, w int) hyperCube {
	return hyperCube{
		x: c.x + x,
		y: c.y + y,
		z: c.z + z,
		w: c.w + w,
	}
}

func (c hyperCube) Neighbors() []hyperCube {
	var neighbors []hyperCube

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					neighbors = append(neighbors, c.Add(x, y, z, w))
				}
			}
		}
	}

	return neighbors
}

type hyperCubeSet map[hyperCube]bool

func newHyperCubeSet() hyperCubeSet {
	return make(hyperCubeSet)
}

func (s hyperCubeSet) Add(c hyperCube) {
	s[c] = true
}

func (s hyperCubeSet) Has(c hyperCube) bool {
	_, exists := s[c]
	return exists
}
