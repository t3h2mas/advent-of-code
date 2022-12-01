package main

import (
	"fmt"

	"github.com/t3h2mas/advent-2020/01/expenses"
	"github.com/t3h2mas/advent-2020/input"
)

func main() {
	expenseEntries, err := input.ReadIntLines("./input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day one, part one:")
	solution, err := expenses.FixReportPartOne(expenseEntries)

	if err != nil {
		panic(err)
	}

	fmt.Printf("solution: %d\n", solution)

	fmt.Println("Day one, part two:")
	solution, err = expenses.FixReportPartTwo(expenseEntries)

	if err != nil {
		panic(err)
	}

	fmt.Printf("solution: %d\n", solution)
}
