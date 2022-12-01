package solution

import (
	"errors"
	"strconv"
	"strings"

	"github.com/t3h2mas/advent-2020/07/graph"
)

// LineParts transforms a line like "it parses a line with two contains clauses"
// into a nested array of strings like [[light, red, bags], [1, bright white bag], ...]
func LineParts(line string) [][]string {
	var parts [][]string
	for idx, chunk := range strings.Split(line, ", ") {
		if idx == 0 {
			rootParts := strings.Split(chunk, " contain ")
			for _, rootPart := range rootParts {
				parts = append(parts, strings.Split(rootPart, " "))
			}
		} else {
			parts = append(parts, strings.Split(chunk, " "))
		}
	}

	return parts
}

type Bag struct {
	Color    string
	Contains []string
}

func BagClauses(parts [][]string) *Bag {
	// parse root
	if len(parts) == 0 {
		panic(errors.New("cannot parse empty parts"))
	}

	result := &Bag{}

	root := parts[0]
	result.Color = strings.Join([]string{root[0], root[1]}, " ")

	// handle additional clauses
	for i := 1; i < len(parts); i++ {
		result.Contains = append(
			result.Contains,
			strings.Join([]string{parts[i][1], parts[i][2]}, " "),
		)
	}

	return result
}

type BagContains struct {
	Color string
	Count int
}

type WeightedBag struct {
	Color    string
	Contains []*BagContains
}

func WeightedBagClauses(parts [][]string) *WeightedBag {
	if len(parts) == 0 {
		panic(errors.New("cannot parse empty parts"))
	}

	result := &WeightedBag{}

	root := parts[0]

	result.Color = strings.Join([]string{root[0], root[1]}, " ")

	for i := 1; i < len(parts); i++ {
		if parts[i][0] == "no" {
			continue
		}
		containsCount, err := strconv.Atoi(parts[i][0])
		if err != nil {
			panic(err)
		}
		containsColor := strings.Join([]string{parts[i][1], parts[i][2]}, " ")

		result.Contains = append(result.Contains, &BagContains{
			Color: containsColor,
			Count: containsCount,
		})
	}

	return result
}

func GraphFromInput(inputLines []string) *graph.Digraph {
	weightedDigraph := graph.NewDigraph()

	for _, ruleLine := range inputLines {
		parts := LineParts(ruleLine)

		bag := WeightedBagClauses(parts)

		for _, contains := range bag.Contains {
			weightedDigraph.AddEdge(bag.Color, contains.Color, contains.Count)
		}
	}
	return weightedDigraph
}
