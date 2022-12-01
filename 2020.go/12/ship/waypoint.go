package ship

type Waypoint struct {
	position Position
}

func NewWaypoint() *Waypoint {
	return &Waypoint{
		position: Position{
			x: 10,
			y: 1,
		},
	}
}

func (w *Waypoint) move(direction CardinalDirection, units int) {
	switch direction {
	case North:
		w.position.y += units
	case South:
		w.position.y -= units
	case East:
		w.position.x += units
	case West:
		w.position.x -= units
	}
}

func (w *Waypoint) rotate(direction Direction, degrees int) {
	switch direction {
	case Left:
		w.position = w.position.RotateCounterClockwise(degrees)

	case Right:
		w.position = w.position.RotateClockwise(degrees)
	}
}

func (w *Waypoint) Act(action rune, units int) {
	switch action {
	case 'N':
		w.move(North, units)
	case 'S':
		w.move(South, units)
	case 'E':
		w.move(East, units)
	case 'W':
		w.move(West, units)
	case 'L':
		w.rotate(Left, units)
	case 'R':
		w.rotate(Right, units)
	}
}
