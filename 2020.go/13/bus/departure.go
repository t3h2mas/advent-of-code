package bus

import (
	"math/big"

	"github.com/t3h2mas/advent-2020/13/remainder"
)

type Departure struct {
	busID int
	time  int
}

func (d Departure) BusID() int {
	return d.busID
}

func (d Departure) Time() int {
	return d.time
}

func ClosestDeparture(buses []Bus, targetDeparture int) Departure {
	var closestDeparture int
	var busID int

	for idx, bus := range buses {
		busDeparture := bus.ID()
		multiple := 1
		for busDeparture <= targetDeparture {
			busDeparture = bus.ID() * multiple
			multiple++
		}

		if idx == 0 {
			closestDeparture = busDeparture
			busID = bus.ID()
			continue
		}

		if busDeparture < closestDeparture {
			closestDeparture = busDeparture
			busID = bus.ID()
		}
	}
	return Departure{
		busID,
		closestDeparture,
	}
}

func EarliestOrderedDeparture(buses map[int]Bus) int {
	var n []*big.Int
	var a []*big.Int

	for idx, bus := range buses {
		a = append(a, big.NewInt(int64(bus.ID())))
		n = append(n, big.NewInt(int64((-1*idx)%bus.ID())))
	}

	r, err := remainder.Crt(n, a)

	if err != nil {
		panic(err)
	}

	return int(r.Int64())
}
