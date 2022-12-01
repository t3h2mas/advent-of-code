package main

import (
	"errors"
	"fmt"

	"github.com/t3h2mas/advent-2020/08/gameconsole"

	"github.com/t3h2mas/advent-2020/input"
)

func main() {
	instructionLines, err := input.ReadLines("./input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day eight, part one")
	machine := gameconsole.NewMachine()
	err = machine.Load(instructionLines)

	fmt.Printf("%#v\n", err)
	fmt.Printf("solution: %d\n", machine.Accumulator())

	fmt.Println("Day eight, part two")
	m := gameconsole.LoadAndPatch(instructionLines)

	if m == nil {
		panic(errors.New("failed to patch instructions"))
	}

	fmt.Printf("solution: %d\n", m.Accumulator())
}
