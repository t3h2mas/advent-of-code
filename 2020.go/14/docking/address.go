package docking

import (
	"fmt"
	"math"
	"strconv"
)

func parseBin(b []byte) int {
	parsed, err := strconv.ParseInt(string(b), 2, 64)
	if err != nil {
		panic(err)
	}
	return int(parsed)
}

func DecodeAddress(address int, mask string) []int {
	var results []int

	bValue := bstr(address)
	var floatingBits []int
	for i := 0; i < 36; i++ {
		bit := mask[i]
		if bit == '0' {
			continue
		}

		if bit == '1' {
			bValue[i] = '1'
		}

		if bit == 'X' {
			bValue[i] = 'X'
			floatingBits = append(floatingBits, i)
		}
	}

	if len(floatingBits) == 0 {
		results = append(results, parseBin(bValue))
		return results
	}

	var floatingOptions []string

	binRange := int(math.Pow(2, float64(len(floatingBits))))
	for i := 0; i < binRange; i++ {
		fstr := "%0" + fmt.Sprintf("%d", len(floatingBits)) + "b"
		opt := fmt.Sprintf(fstr, i)
		if len(opt) > len(floatingBits) {
			break
		}
		floatingOptions = append(floatingOptions, opt)
	}

	for _, opt := range floatingOptions {
		masked := append([]byte{}, bValue...)
		for bidx, optBit := range opt {
			masked[floatingBits[bidx]] = byte(optBit)
		}

		results = append(results, parseBin(masked))
	}

	return results
}
