package main

import (
	"fmt"

	"github.com/t3h2mas/advent-2020/10/charging"

	"github.com/t3h2mas/advent-2020/input"
)

func main() {
	adapters, err := input.ReadIntLines("./input")

	if err != nil {
		panic(err)
	}

	fmt.Println("Day ten, part one")
	fmt.Printf("solution: %d\n", charging.ChainVoltageDifferences(adapters))

	fmt.Println("Day ten, part two")
	fmt.Printf("solution: %d\n", charging.ChainPossibilities(adapters))
}
