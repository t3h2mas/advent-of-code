package tickets

import (
	"regexp"
	"strconv"
	"strings"
)

/*
class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50
*/

func mustParseInt(s string) int {
	result, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return result
}

func ParseRules(input string) []Rule {
	re := regexp.MustCompile(`(\w+\s?\w+?): (\d+)-(\d+) or (\d+)-(\d+)`)

	var results []Rule

	for _, line := range strings.Split(input, "\n") {
		parts := re.FindStringSubmatch(line)

		results = append(
			results,
			Rule{
				field: parts[1],
				ranges: []intRange{
					{mustParseInt(parts[2]), mustParseInt(parts[3])},
					{mustParseInt(parts[4]), mustParseInt(parts[5])},
				},
			})
	}

	return results
}

func ParseTickets(input string) []Ticket {
	var results []Ticket
	for idx, line := range strings.Split(input, "\n") {
		// disregard the first line
		if idx == 0 {
			continue
		}

		var values []int

		for _, n := range strings.Split(line, ",") {
			values = append(values, mustParseInt(n))
		}

		results = append(results, Ticket{values})
	}

	return results
}
