package ship

import "math"

type Position struct {
	x int
	y int
}

func (p Position) X() int {
	return p.x
}

func (p Position) Y() int {
	return p.y
}

func (p Position) RotateClockwise(degrees int) Position {
	x := float64(p.x)
	y := float64(p.y)

	angle := float64(degrees) * math.Pi / 180

	s := math.Sin(angle)
	c := math.Cos(angle)

	xNew := x*c + y*s
	yNew := x*-1*s + y*c

	return Position{int(math.Round(xNew)), int(math.Round(yNew))}
}

func (p Position) RotateCounterClockwise(degrees int) Position {
	x := float64(p.x)
	y := float64(p.y)

	angle := float64(degrees) * math.Pi / 180

	s := math.Sin(angle)
	c := math.Cos(angle)

	xNew := x*c - y*s
	yNew := x*s + y*c

	return Position{int(math.Round(xNew)), int(math.Round(yNew))}
}
