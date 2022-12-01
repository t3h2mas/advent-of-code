package expenses

import (
	"errors"
)

/**
 * Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle input); apparently, something isn't quite adding up.
 * Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.
 */

func FixReportPartOne(entries []int) (int, error) {
	components, err := sumsTo(2020, entries)
	if err != nil {
		return 0, err
	}

	return components[0] * components[1], nil
}

func FixReportPartTwo(entries []int) (int, error) {
	components, err := tripletSumsTo(2020, entries)
	if err != nil {
		return 0, err
	}

	return components[0] * components[1] * components[2], nil
}

type IntSet struct {
	data map[int]bool
}

func NewIntSet() *IntSet {
	return &IntSet{
		data: make(map[int]bool),
	}
}

func (s *IntSet) Has(val int) bool {
	_, hasValue := s.data[val]
	return hasValue
}

func (s *IntSet) Add(val int) {
	s.data[val] = true
}

func sumsTo(target int, search []int) ([2]int, error) {
	var result [2]int

	seen := NewIntSet()

	for _, n := range search {
		if seen.Has(target - n) {
			result[0] = n
			result[1] = target - n
			return result, nil
		}
		seen.Add(n)
	}

	return result, errors.New("unable to find two numbers that sum to target")
}

func tripletSumsTo(target int, search []int) ([3]int, error) {
	var result [3]int
	partialSums := make(map[int][2]int)

	for i := 0; i < len(search); i++ {
		for j := 0; j < len(search); j++ {
			if i == j {
				continue
			}
			partialSums[search[i]+search[j]] = [2]int{search[i], search[j]}
		}
	}

	for _, n := range search {
		complement := target - n
		if p, hasComplement := partialSums[complement]; hasComplement {
			result[0] = n
			result[1] = p[0]
			result[2] = p[1]

			return result, nil
		}
	}
	return result, errors.New("unable to find three numbers that sum to target")
}
