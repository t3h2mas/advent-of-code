package seats

func neighborsOf(seat Point) []Point {
	var results []Point
	for xAway := -1; xAway < 2; xAway++ {
		for yAway := -1; yAway < 2; yAway++ {
			if xAway == 0 && yAway == 0 {
				continue
			}
			results = append(
				results,
				NewPoint(
					seat.X()+xAway,
					seat.Y()+yAway,
				),
			)
		}
	}
	return results
}

func calculateOccupiedLineOfSight(layout Layout) PointSet {
	result := make(PointSet)

	directions := []struct {
		x int
		y int
	}{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	for y := 0; y < layout.height; y++ {
		for x := 0; x < layout.width; x++ {
			seat := NewPoint(x, y)

			if layout.floor.Has(seat) {
				continue
			}

			occupied := 0
			for _, direction := range directions {
				currentSeat := NewPoint(seat.X()+direction.x, seat.Y()+direction.y)

				for {
					if layout.SeatOutOfBounds(currentSeat) {
						break
					}

					if layout.floor.Has(currentSeat) {
						// update seat and continue looping
						currentSeat = NewPoint(currentSeat.X()+direction.x, currentSeat.Y()+direction.y)
						continue
					}

					// a seat has been found
					if layout.occupied.Has(currentSeat) {
						occupied++
					}
					break
				}
			}

			if layout.occupied.Has(seat) {
				if occupied < 5 {
					result.Add(seat)
				}
			} else {
				if occupied == 0 {
					result.Add(seat)
				}
			}
		}
	}

	return result
}

func calculateOccupied(layout Layout) PointSet {
	result := make(PointSet)

	for y := 0; y < layout.height; y++ {
		for x := 0; x < layout.width; x++ {
			seat := NewPoint(x, y)

			if layout.floor.Has(seat) {
				continue
			}

			occupiedNeighbors := 0

			for _, neighbor := range neighborsOf(seat) {
				if layout.SeatOutOfBounds(neighbor) {
					continue
				}
				if layout.occupied.Has(neighbor) {
					occupiedNeighbors++
				}
			}

			if layout.occupied.Has(seat) {
				// if a seat is occupied, and four or more seats adjacent are also occupied, the seat becomes empty
				if occupiedNeighbors < 4 {
					result.Add(seat)
				}
			} else {
				// if a seat is empty and there are no occupied neighbors the seat becomes occupied
				if occupiedNeighbors == 0 {
					result.Add(seat)
				}
			}
			if !layout.occupied.Has(seat) && occupiedNeighbors == 0 {
				result.Add(seat)
			}
		}
	}

	return result
}

func SimulateStableOccupiedSeatCount(layout Layout) int {
	for {
		nextOccupiedSeats := calculateOccupied(layout)

		if layout.occupied.Equals(nextOccupiedSeats) {
			break
		}

		layout.occupied = nextOccupiedSeats
	}

	return len(layout.occupied)
}

func SimulateStableOccupiedLineOfSightSeatCount(layout Layout) int {
	for {
		nextOccupiedSeats := calculateOccupiedLineOfSight(layout)

		if layout.occupied.Equals(nextOccupiedSeats) {
			break
		}

		layout.occupied = nextOccupiedSeats
	}

	return len(layout.occupied)

}
