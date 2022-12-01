package charging

import (
	"sort"

	"github.com/t3h2mas/advent-2020/types"
)

func ChainVoltageDifferences(adapters []int) int {
	sort.Ints(adapters)

	// handle built in voltage by adding 1 to the differences of 3 group
	differenceMap := map[int]int{3: 1}
	// handle the voltage difference from the initial device (0)
	differenceMap[adapters[0]]++
	differences := voltageDifferences(adapters, differenceMap)

	return differences[1] * differences[3]
}

func voltageDifferences(adapters []int, differences map[int]int) map[int]int {
	if len(adapters) <= 1 {
		return differences
	}

	adapterVoltage := adapters[0]
	diff := adapters[1] - adapterVoltage

	differences[diff]++

	return voltageDifferences(adapters[1:], differences)
}

// Dynamic programming approach to day ten, part two
// "What is the total number of distinct ways you can arrange the adapters to connect the charging outlet to your device?"
// This approach is a variant of the 'climb k stairs, skip red' problem set. Instead of a list of "red" spaces, we use
// the `available` set as a whitelist of voltages that have an adapter.
// Using this approach gives us a runtime of O(n * 3) or O(n)
func ChainPossibilities(adapters []int) int {
	maxVoltage := max(adapters)

	// voltages that have an adapter
	available := types.NewIntSet()

	for _, voltage := range adapters {
		available.Add(voltage)
	}

	dp := make([]int, maxVoltage+1)

	// base case, there is only one way to get to the first spot
	dp[0] = 1

	for i := 1; i <= maxVoltage; i++ {
		for j := 1; j < 4; j++ {
			if i-j < 0 {
				continue
			}
			// skip voltages that don't match an adapter
			if !available.Has(i) {
				dp[i] = 0
			} else {
				dp[i] += dp[(i - j)]
			}
		}
	}

	return dp[maxVoltage]
}

func max(nums []int) int {
	m := nums[0]

	for _, n := range nums {
		if n > m {
			m = n
		}
	}

	return m
}
