package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/t3h2mas/advent-2020/04/passport"

	"github.com/t3h2mas/advent-2020/input"
)

func splitEntries(entryLog string) []string {
	return regexp.MustCompile(`(?m)^\s*$`).Split(entryLog, -1)
}

func normalizeEntry(entry string) string {
	return strings.TrimSpace(regexp.MustCompile(`\r?\n`).ReplaceAllString(entry, " "))
}

func main() {
	passportBatchEntryLog, err := input.ReadAll("./input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day four, part two")

	validEntries := 0
	entries := splitEntries(passportBatchEntryLog)
	for _, entry := range entries {
		normalized := normalizeEntry(entry)
		_, err = passport.FromBatchEntry(normalized)
		if err == nil {
			validEntries++
		}
	}

	fmt.Printf("solution: %d\n", validEntries)
}
