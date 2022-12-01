package hyperlife

import "strings"

type Game struct {
	cubes      hyperCubeSet
	boundaries *boundaries
}

func (g *Game) activate(c hyperCube) {
	g.cubes.Add(c)

	if c.x > g.boundaries.x.high {
		g.boundaries.x.high = c.x
	}

	if c.x < g.boundaries.x.low {
		g.boundaries.x.low = c.x
	}

	if c.y > g.boundaries.y.high {
		g.boundaries.y.high = c.y
	}

	if c.y < g.boundaries.y.low {
		g.boundaries.y.low = c.y
	}

	if c.z > g.boundaries.z.high {
		g.boundaries.z.high = c.z
	}

	if c.z < g.boundaries.z.low {
		g.boundaries.z.low = c.z
	}

	if c.w > g.boundaries.w.high {
		g.boundaries.w.high = c.w
	}

	if c.w < g.boundaries.w.low {
		g.boundaries.w.low = c.w
	}
}

func (g *Game) Step() *Game {
	nextState := newGame()

	for c := range g.cubes {
		activeNeighbors := g.activeNeighbors(c)

		if activeNeighbors == 2 || activeNeighbors == 3 {
			nextState.activate(c)
		}
	}

	for x := g.boundaries.x.low - 1; x <= g.boundaries.x.high+1; x++ {
		for y := g.boundaries.y.low - 1; y <= g.boundaries.y.high+1; y++ {
			for z := g.boundaries.z.low - 1; z <= g.boundaries.z.high+1; z++ {
				for w := g.boundaries.w.low - 1; w <= g.boundaries.w.high+1; w++ {
					c := newHyperCube(x, y, z, w)

					if g.cubes.Has(c) {
						continue
					}

					if g.activeNeighbors(c) == 3 {
						nextState.activate(c)
					}
				}
			}
		}
	}

	return nextState
}

func (g *Game) activeNeighbors(c hyperCube) int {
	activeCount := 0

	for _, neighbor := range c.Neighbors() {
		if g.cubes.Has(neighbor) {
			activeCount++
		}
	}

	return activeCount
}

func (g *Game) Active() int {
	return len(g.cubes)
}

func newGame() *Game {
	return &Game{
		cubes: newHyperCubeSet(),
		boundaries: &boundaries{
			x: &limit{
				high: 0,
				low:  0,
			},
			y: &limit{
				high: 0,
				low:  0,
			},
			z: &limit{
				high: 0,
				low:  0,
			},
			w: &limit{
				high: 0,
				low:  0,
			},
		},
	}
}

func GameFromSeed(seed string) *Game {
	game := newGame()

	for _, c := range parseSeed(seed) {
		game.activate(c)
	}

	return game
}

func parseSeed(seed string) []hyperCube {
	var results []hyperCube

	lines := strings.Split(seed, "\n")
	reverseStrings(lines)

	for y, line := range lines {
		for x, ch := range line {
			switch ch {
			case '.':
				break
			case '#':
				results = append(results, newHyperCube(x, y, 0, 0))
			}
		}
	}

	return results
}

type boundaries struct {
	x, y, z, w *limit
}

type limit struct {
	high int
	low  int
}

func reverseStrings(xs []string) {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
}
