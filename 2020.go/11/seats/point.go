package seats

import "fmt"

type Point struct {
	x int
	y int
}

func (p Point) X() int {
	return p.x
}

func (p Point) Y() int {
	return p.y
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

type PointSet map[Point]bool

func NewPointSet() PointSet {
	return make(map[Point]bool)
}

func (s PointSet) Has(val Point) bool {
	_, exists := s[val]
	return exists
}

func (s PointSet) Add(val Point) {
	s[val] = true
}

func (s PointSet) Remove(val Point) {
	delete(s, val)
}

func (s PointSet) Equals(other PointSet) bool {
	return fmt.Sprint(s) == fmt.Sprint(other)
}
