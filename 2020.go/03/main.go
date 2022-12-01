package main

import (
	"fmt"

	"github.com/t3h2mas/advent-2020/03/slopes"

	"github.com/t3h2mas/advent-2020/03/grid"

	"github.com/t3h2mas/advent-2020/input"
)

func main() {
	var slope grid.ScrollingGrid
	slope, err := input.ReadLines("./input")
	if err != nil {
		panic(err)
	}

	tree := '#'

	fmt.Printf("Day three, part one")

	direction := slopes.NewTobogganDirection(3, 1)

	encounters := slopes.PathEncounters(slope, direction)

	treeCount := encounters[tree]
	fmt.Printf("solution: %d\n", treeCount)

	fmt.Println("Day three, part two")
	directions := []*slopes.TobogganDirection{
		slopes.NewTobogganDirection(1, 1),
		slopes.NewTobogganDirection(3, 1),
		slopes.NewTobogganDirection(5, 1),
		slopes.NewTobogganDirection(7, 1),
		slopes.NewTobogganDirection(1, 2),
	}

	treeEncounters := 1
	for _, d := range directions {
		encounters := slopes.PathEncounters(slope, d)
		treeEncounters = treeEncounters * encounters[tree]
	}
	fmt.Printf("solution: %d\n", treeEncounters)
}
