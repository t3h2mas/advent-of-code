package main

import (
	"fmt"

	"github.com/t3h2mas/advent-2020/05/boarding"
	"github.com/t3h2mas/advent-2020/input"
)

type RowColSet map[[2]int]bool

func NewRowColSet() RowColSet {
	var rcs RowColSet
	rcs = make(map[[2]int]bool)

	return rcs
}

func (s RowColSet) Add(row, col int) {
	s[[2]int{row, col}] = true
}

func (s RowColSet) Has(row, col int) bool {
	_, hasRowCol := s[[2]int{row, col}]

	return hasRowCol
}

func main() {
	passes, err := input.ReadLines("./input")

	if err != nil {
		panic(err)
	}

	fmt.Println("Day five, part one")
	maxID := 0
	for _, p := range passes {
		pass, err := boarding.PassFrom(p)

		if err != nil {
			panic(err)
		}

		if pass.SeatID() > maxID {
			maxID = pass.SeatID()
		}
	}
	fmt.Printf("solution: %d\n", maxID)

	fmt.Println("Day five, part two")
	maxRow := 0
	minRow := 100000
	setIDs := NewRowColSet()

	for _, p := range passes {
		pass, err := boarding.PassFrom(p)

		if err != nil {
			panic(err)
		}

		setIDs.Add(pass.Row(), pass.Col())
		if pass.Row() > maxRow {
			maxRow = pass.Row()
		}

		if pass.Row() < minRow {
			minRow = pass.Row()
		}
	}

	for row := minRow + 1; row <= maxRow; row++ {
		for col := 0; col < 8; col++ {
			if !setIDs.Has(row, col) {
				// make a new pass to generate the seat ID
				myPass := boarding.NewPass(row, col)
				fmt.Printf("%+v\n", myPass)

				fmt.Printf("solution: %d\n", myPass.SeatID())
				return
			}
		}
	}
}
