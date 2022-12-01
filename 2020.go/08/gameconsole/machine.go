package gameconsole

import (
	"errors"
	"fmt"
	"strings"
)

type Machine struct {
	accumulator int
	cursor      int
	visited     map[int]bool
}

func NewMachine() *Machine {
	return &Machine{
		0,
		0,
		map[int]bool{0: true},
	}
}

func (m *Machine) markVisited(cursor int) {
	m.visited[cursor] = true
}

func (m *Machine) haveVisited(cursor int) bool {
	_, haveVisited := m.visited[cursor]
	return haveVisited
}

var MachineCycleError = errors.New("machine cycle detected")

func (m *Machine) addToCursor(amount int) error {
	nextCursor := m.cursor + amount

	if m.haveVisited(nextCursor) {
		return MachineCycleError
	}

	m.cursor = nextCursor
	m.markVisited(m.cursor)

	return nil
}

func (m *Machine) Handle(instruction Instruction) error {
	var err error = nil
	switch instruction.Op.value {
	case "nop":
		err = m.addToCursor(1)
	case "acc":
		m.accumulator += instruction.Arg.value
		err = m.addToCursor(1)
	case "jmp":
		err = m.addToCursor(instruction.Arg.value)
	}

	return err
}

func (m *Machine) Accumulator() int {
	return m.accumulator
}

func (m *Machine) Cursor() int {
	return m.cursor
}

func (m *Machine) Clone() *Machine {
	visitedCopy := make(map[int]bool)
	for k, v := range m.visited {
		visitedCopy[k] = v
	}
	return &Machine{
		accumulator: m.accumulator,
		cursor:      m.cursor,
		visited:     visitedCopy,
	}
}

func (m *Machine) Load(instructions []string) error {
	instruction := InstructionFromString(instructions[0])

	for {
		err := m.Handle(instruction)
		if err != nil {
			return err
		}

		if m.Cursor() == len(instructions) {
			break
		}

		instruction = InstructionFromString(instructions[m.Cursor()])
	}

	return nil
}

func LoadAndPatch(instructions []string) *Machine {
	var instructionPossibilities [][]string

	for idx, instruction := range instructions {
		if strings.HasPrefix(instruction, "jmp") {
			instructionCopy := make([]string, len(instructions))
			copy(instructionCopy, instructions)
			instructionCopy[idx] = strings.Replace(instruction, "jmp", "nop", 1)
			instructionPossibilities = append(instructionPossibilities, instructionCopy)
		}

		if strings.HasPrefix(instruction, "nop") {
			instructionCopy := make([]string, len(instructions))
			copy(instructionCopy, instructions)
			instructionCopy[idx] = strings.Replace(instruction, "nop", "jmp", 1)
			instructionPossibilities = append(instructionPossibilities, instructionCopy)
		}
	}
	fmt.Printf("Working with %d instruction sets\n", len(instructionPossibilities))

	for _, possibility := range instructionPossibilities {
		m := NewMachine()

		err := m.Load(possibility)

		if err == nil {
			return m
		}
	}

	return nil
}
