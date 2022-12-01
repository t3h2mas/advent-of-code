package docking

import (
	"fmt"
	"strconv"
)

type Program struct {
	memory map[int]int
}

func NewProgram() Program {
	return Program{
		make(map[int]int),
	}
}

func bstr(n int) []byte {
	return []byte(fmt.Sprintf("%036b", n))
}

func (p Program) Memory() map[int]int {
	return p.memory
}

func (p Program) WriteTo(address int, value int) {
	p.memory[address] = value
}

func (p Program) Write(address int, value int, mask string) {
	bValue := bstr(value)
	for i := 0; i < 36; i++ {
		if mask[i] == 'X' {
			continue
		}

		if mask[i] == '1' {
			bValue[i] = '1'
		}

		if mask[i] == '0' {
			bValue[i] = '0'
		}
	}
	parsed, err := strconv.ParseInt(string(bValue), 2, 64)
	if err != nil {
		panic(err)
	}

	p.WriteTo(address, int(parsed))
}
