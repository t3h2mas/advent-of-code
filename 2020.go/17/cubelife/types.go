package cubelife

type cubeSet map[cube]bool

func newCubeSet() cubeSet {
	return make(map[cube]bool)
}

func (s cubeSet) Add(c cube) {
	s[c] = true
}

func (s cubeSet) Has(c cube) bool {
	_, exists := s[c]
	return exists
}

func (s cubeSet) Remove(c cube) {
	delete(s, c)
}

type cubeTracker struct {
	cubes      cubeSet
	boundaries boundaries
}

func newCubeTracker() *cubeTracker {
	return &cubeTracker{
		cubes: newCubeSet(),
		boundaries: boundaries{
			x: &limit{0, 0},
			y: &limit{0, 0},
			z: &limit{0, 0},
		},
	}
}

func (ct cubeTracker) Add(c cube) {
	ct.cubes.Add(c)

	if c.x > ct.boundaries.x.high {
		ct.boundaries.x.high = c.x
	}

	if c.x < ct.boundaries.x.low {
		ct.boundaries.x.low = c.x
	}

	if c.y > ct.boundaries.y.high {
		ct.boundaries.y.high = c.y
	}

	if c.y < ct.boundaries.y.low {
		ct.boundaries.y.low = c.y
	}

	if c.z > ct.boundaries.z.high {
		ct.boundaries.z.high = c.z
	}

	if c.z < ct.boundaries.z.low {
		ct.boundaries.z.low = c.z
	}
}

type limit struct {
	high int
	low  int
}

type boundaries struct {
	x *limit
	y *limit
	z *limit
}
