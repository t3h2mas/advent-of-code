package bus

import (
	"strconv"
	"strings"
)

type Bus struct {
	id int
}

func (b Bus) ID() int {
	return b.id
}

func BusesFrom(input string) []Bus {
	var result []Bus

	for _, bid := range strings.Split(input, ",") {
		if bid == "x" {
			continue
		}

		parsed, err := strconv.Atoi(bid)
		if err != nil {
			panic(err)
		}

		result = append(result, Bus{id: parsed})
	}

	return result
}

func IndexedBusesFrom(input string) map[int]Bus {
	result := make(map[int]Bus)

	for idx, bid := range strings.Split(input, ",") {
		if bid == "x" {
			continue
		}

		parsed, err := strconv.Atoi(bid)
		if err != nil {
			panic(err)
		}

		result[idx] = Bus{id: parsed}
	}

	return result
}
