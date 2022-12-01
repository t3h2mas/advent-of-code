package cubelife

type cube struct {
	x int
	y int
	z int
}

func newCube(x, y, z int) cube {
	return cube{x, y, z}
}

func (c cube) Add(x, y, z int) cube {
	return cube{
		x: c.x + x,
		y: c.y + y,
		z: c.z + z,
	}
}

func (c cube) Neighbors() []cube {
	var neighbors []cube

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				neighbors = append(neighbors, c.Add(x, y, z))
			}
		}
	}

	return neighbors
}
