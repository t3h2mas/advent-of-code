package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/t3h2mas/advent-2020/14/docking"

	"github.com/t3h2mas/advent-2020/input"
)

func maskFor(s string) string {
	parts := strings.Split(s, " ")

	return parts[2]
}

type writeCmd struct {
	address int
	value   int
}

func writeCmdFor(s string) writeCmd {
	re := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	parts := re.FindStringSubmatch(s)

	address, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	value, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}

	return writeCmd{
		address: address,
		value:   value,
	}
}

func main() {
	commands, err := input.ReadLines("./input")
	if err != nil {
		panic(err)
	}

	{
		fmt.Println("Day fourteen, part one")
		program := docking.NewProgram()
		var mask string
		for _, instruction := range commands {
			if strings.HasPrefix(instruction, "mask") {
				mask = maskFor(instruction)
				continue
			}

			cmd := writeCmdFor(instruction)

			program.Write(cmd.address, cmd.value, mask)
		}

		sum := 0
		for _, val := range program.Memory() {
			sum += val
		}

		fmt.Printf("solution: %d\n", sum)
	}

	{
		fmt.Println("Day fourteen, part two")
		sum := 0
		var mask string
		program := docking.NewProgram()
		for _, instruction := range commands {
			if strings.HasPrefix(instruction, "mask") {
				mask = maskFor(instruction)
				continue
			}

			cmd := writeCmdFor(instruction)

			addresses := docking.DecodeAddress(cmd.address, mask)
			for _, address := range addresses {
				program.WriteTo(address, cmd.value)
			}
		}
		for _, val := range program.Memory() {
			sum += val
		}

		fmt.Printf("solution: %d\n", sum)
	}
}
