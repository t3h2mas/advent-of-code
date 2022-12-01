package cubelife

import (
	"fmt"
	"strings"
)

type Game struct {
	active *cubeTracker
}

func parseSeed(seed string) []cube {
	var results []cube

	lines := strings.Split(seed, "\n")
	reverseStrings(lines)

	for y, line := range lines {
		for x, ch := range line {
			switch ch {
			case '.':
				break
			case '#':
				results = append(results, newCube(x, y, 0))
			}
		}
	}
	return results
}

func (g *Game) activeNeighbors(c cube) int {
	activeCount := 0

	for _, neighbor := range c.Neighbors() {
		if g.active.cubes.Has(neighbor) {
			activeCount++
		}
	}

	return activeCount
}

func (g *Game) Step() {
	nextActive := newCubeTracker()

	for c := range g.active.cubes {
		activeNeighbors := g.activeNeighbors(c)

		// If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains active.
		// Otherwise the cube becomes inactive
		if activeNeighbors == 2 || activeNeighbors == 3 {
			nextActive.Add(c)
		}
	}

	for x := g.active.boundaries.x.low - 1; x <= g.active.boundaries.x.high+1; x++ {
		for y := g.active.boundaries.y.low - 1; y <= g.active.boundaries.y.high+1; y++ {
			for z := g.active.boundaries.z.low - 1; z <= g.active.boundaries.z.high+1; z++ {
				c := newCube(x, y, z)
				if g.active.cubes.Has(c) {
					continue
				}

				// If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active. Otherwise, the cube remains inactive.
				if g.activeNeighbors(c) == 3 {
					nextActive.Add(c)
				}
			}
		}
	}

	g.active = nextActive
}

func (g *Game) cubes() cubeSet {
	return g.active.cubes
}

func GameFromSeed(seed string) *Game {
	active := newCubeTracker()

	for _, c := range parseSeed(seed) {
		active.Add(c)
	}

	return &Game{
		active: active,
	}
}

func reverseStrings(xs []string) {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
}

func (g *Game) String() string {
	var b strings.Builder

	for z := g.active.boundaries.z.low; z <= g.active.boundaries.z.high; z++ {
		b.WriteString(fmt.Sprintf("z=%d\n", z))
		for y := g.active.boundaries.y.high; y >= g.active.boundaries.y.low; y-- {
			for x := g.active.boundaries.x.low; x <= g.active.boundaries.x.high; x++ {
				if g.active.cubes.Has(cube{x, y, z}) {
					b.WriteString("#")
				} else {
					b.WriteString(".")
				}
			}

			b.WriteString(fmt.Sprintf("\n"))
		}

		if z != g.active.boundaries.z.high {
			b.WriteString(fmt.Sprintf("\n"))
		}
	}

	return b.String()
}

func (g *Game) Active() int {
	return len(g.active.cubes)
}
