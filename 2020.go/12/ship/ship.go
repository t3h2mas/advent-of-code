package ship

type Ship struct {
	facing   CardinalDirection
	position Position
}

func (s *Ship) Position() Position {
	return s.position
}

func (s *Ship) faceTowards(d CardinalDirection) {
	s.facing = d
}

func (s *Ship) turnRight() {
	switch s.facing {
	case North:
		s.faceTowards(East)
	case East:
		s.faceTowards(South)
	case South:
		s.faceTowards(West)
	case West:
		s.faceTowards(North)
	}
}

func (s *Ship) turnLeft() {
	switch s.facing {
	case North:
		s.faceTowards(West)
	case West:
		s.faceTowards(South)
	case South:
		s.faceTowards(East)
	case East:
		s.faceTowards(North)
	}
}

func (s *Ship) turn(direction Direction, degrees int) {
	numberOfTurns := degrees / 90

	switch direction {
	case Left:
		for i := 0; i < numberOfTurns; i++ {
			s.turnLeft()
		}
	case Right:
		for i := 0; i < numberOfTurns; i++ {
			s.turnRight()
		}
	}
}

func NewShip() *Ship {
	ship := &Ship{}

	// The ship starts by facing east
	ship.faceTowards(East)

	return ship
}

func (s *Ship) Act(action rune, units int) {
	switch action {
	case 'F':
		// move forward by the given value in the direction the ship is currently facing
		s.move(s.facing, units)

	case 'N':
		// move north by the given value
		s.move(North, units)

	case 'S':
		// move south by the given value
		s.move(South, units)

	case 'E':
		// move east by the given value
		s.move(East, units)

	case 'W':
		// move west by the given value
		s.move(West, units)

	case 'L':
		// turn left by the given number of degrees
		s.turn(Left, units)

	case 'R':
		// turn right by the given number of degrees
		s.turn(Right, units)
	}
}

func (s *Ship) move(direction CardinalDirection, units int) {
	switch direction {
	case North:
		s.position.y += units
	case South:
		s.position.y -= units
	case East:
		s.position.x += units
	case West:
		s.position.x -= units
	}
}
