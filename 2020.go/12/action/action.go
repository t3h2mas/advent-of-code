package action

import (
	"strconv"
	"strings"
)

type Action struct {
	actionType rune
	units      int
}

func (a *Action) Units() int {
	return a.units
}

func (a *Action) Type() rune {
	return a.actionType
}

func NewActionFrom(input string) *Action {
	actionType := input[0]

	var unitStr strings.Builder

	for i := 1; i < len(input); i++ {
		unitStr.WriteByte(input[i])
	}

	units, err := strconv.Atoi(unitStr.String())
	if err != nil {
		panic(err)
	}

	return &Action{
		rune(actionType),
		units,
	}
}
