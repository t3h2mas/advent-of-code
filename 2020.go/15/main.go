package main

import (
	"fmt"

	"github.com/t3h2mas/advent-2020/15/memgame"
)

func main() {
	startingNumbers := []int{1, 0, 16, 5, 17, 4}

	fmt.Println("Day fifteen, part one")
	fmt.Printf("solution: %d\n", memgame.SpokenAt(startingNumbers, 2020))

	fmt.Println("Day fifteen, part two")
	fmt.Printf("solution: %d\n", memgame.SpokenAt(startingNumbers, 30000000))
}
