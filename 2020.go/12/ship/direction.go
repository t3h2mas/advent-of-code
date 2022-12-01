package ship

type CardinalDirection int

const (
	North CardinalDirection = iota
	East
	South
	West
)

func (d CardinalDirection) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

type Direction int

const (
	Forward Direction = iota
	Left
	Right
)

func (d Direction) String() string {
	return [...]string{"Forward", "Left", "Right"}[d]
}
