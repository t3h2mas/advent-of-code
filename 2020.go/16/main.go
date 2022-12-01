package main

import (
	"fmt"

	"github.com/t3h2mas/advent-2020/16/solution"

	"github.com/t3h2mas/advent-2020/input"
)

func main() {
	problemInput, err := input.ReadAll("./input")

	if err != nil {
		panic(err)
	}

	fmt.Println("Day sixteen, part one")
	fmt.Printf("solution: %d\n", solution.TicketScanningErrorRate(problemInput))

	fmt.Println("Day sixteen, part two")
	fmt.Printf("solution: %d\n", solution.TicketDepartureProduct(problemInput))
}
