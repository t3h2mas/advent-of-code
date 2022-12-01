package main

import (
	"fmt"

	"github.com/t3h2mas/advent-2020/09/xmas"
	"github.com/t3h2mas/advent-2020/input"
)

func main() {
	packets, err := input.ReadIntLines("./input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day nine, part one")
	message := xmas.NewMessage(packets, 25)
	vulnerableNumber := message.Crack()
	fmt.Printf("solution: %d\n", vulnerableNumber)

	fmt.Println("Day nine, part two")
	fmt.Printf("solution: %d\n", message.Exploit(vulnerableNumber))

}
