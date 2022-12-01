package gameconsole

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Operation struct {
	value string
}

func NewOperation(value string) Operation {
	validOperations := []string{
		"acc",
		"jmp",
		"nop",
	}

	for _, validOp := range validOperations {
		if value == validOp {
			return Operation{value: value}
		}
	}

	panic(fmt.Errorf("could not create operation with invalid type: %s", value))
}

type Argument struct {
	value int
}

func ArgumentFromString(s string) Argument {
	parsed, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return Argument{value: parsed}
}

type Instruction struct {
	Op  Operation
	Arg Argument
}

func InstructionFromString(s string) Instruction {
	parts := strings.Split(s, " ")

	if len(parts) != 2 {
		panic(errors.New("failed to parse parts"))
	}

	return Instruction{
		Op:  NewOperation(parts[0]),
		Arg: ArgumentFromString(parts[1]),
	}
}
