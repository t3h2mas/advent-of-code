package main

import (
	"fmt"
	"math"

	"github.com/t3h2mas/advent-2020/12/action"

	"github.com/t3h2mas/advent-2020/12/ship"

	"github.com/t3h2mas/advent-2020/input"
)

func main() {
	actions, err := input.ReadLines("./input")

	if err != nil {
		panic(err)
	}

	fmt.Println("Day twelve, part one")
	s := ship.NewShip()
	for _, actionInput := range actions {
		a := action.NewActionFrom(actionInput)

		s.Act(a.Type(), a.Units())
	}
	pos := s.Position()

	result := math.Abs(float64(pos.X())) + math.Abs(float64(pos.Y()))

	fmt.Printf("solution %d\n", int(result))

	fmt.Println("Day twelve, part two")
	ws := ship.NewWaypointShip()
	for _, actionInput := range actions {
		a := action.NewActionFrom(actionInput)

		ws.Act(a.Type(), a.Units())
	}

	pos = ws.ShipPosition()

	result = math.Abs(float64(pos.X())) + math.Abs(float64(pos.Y()))
	fmt.Printf("solution %d\n", int(result))
}
