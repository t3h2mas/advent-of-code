package main

import (
	"fmt"

	"github.com/t3h2mas/advent-2020/07/graph"

	"github.com/t3h2mas/advent-2020/07/solution"

	"github.com/t3h2mas/advent-2020/07/luggage"

	"github.com/t3h2mas/advent-2020/input"
)

type StringSet map[string]bool

func (s StringSet) Add(val string) {
	s[val] = true
}

func (s StringSet) Has(val string) bool {
	_, exists := s[val]
	return exists
}

func NewStringSet() StringSet {
	return make(map[string]bool)
}

func main() {
	fmt.Println("Day seven, part one")
	bagRuleLines, err := input.ReadLines("./input")
	if err != nil {
		panic(err)
	}

	rules := luggage.NewBagRules()

	for _, ruleLine := range bagRuleLines {
		parts := solution.LineParts(ruleLine)

		bag := solution.BagClauses(parts)

		for _, contains := range bag.Contains {
			rules.AddRule(contains, bag.Color)
		}
	}

	targetBag := rules.Get("shiny gold")

	containingTarget := 0
	var queue []string
	seen := NewStringSet()

	for containedByColor := range targetBag.ContainedBy {
		queue = append(queue, containedByColor)
		seen.Add(containedByColor)
	}

	for len(queue) > 0 {
		color := queue[0]
		queue = queue[1:]

		containingTarget++

		bag := rules.Get(color)

		for containedByColor := range bag.ContainedBy {
			if !seen.Has(containedByColor) {
				seen.Add(containedByColor)
				queue = append(queue, containedByColor)
			}
		}
	}

	fmt.Printf("solution: %d\n", containingTarget)

	fmt.Println("Day seven, part two")
	weightedDigraph := solution.GraphFromInput(bagRuleLines)

	target := weightedDigraph.Get("shiny gold")
	fmt.Printf("solution: %d\n", graph.ComputeAggregateWeight(target))

}
