package main

import (
	"fmt"

	"github.com/t3h2mas/advent-2020/17/hyperlife"

	"github.com/t3h2mas/advent-2020/17/cubelife"

	"github.com/t3h2mas/advent-2020/input"
)

func main() {
	fmt.Println("Day seventeen, part one")
	seed, err := input.ReadAll("./input")

	if err != nil {
		panic(err)
	}

	game := cubelife.GameFromSeed(seed)

	for i := 0; i < 6; i++ {
		game.Step()
	}

	fmt.Printf("solution: %d\n", game.Active())

	fmt.Println("Day seventeen, part two")

	hyperGame := hyperlife.GameFromSeed(seed)

	for i := 0; i < 6; i++ {
		hyperGame = hyperGame.Step()
	}

	fmt.Printf("solution: %d\n", hyperGame.Active())
}
