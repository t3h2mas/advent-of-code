package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/t3h2mas/advent-2020/06/customs"

	"github.com/t3h2mas/advent-2020/input"
)

func splitEntries(entryLog string) []string {
	return regexp.MustCompile(`(?m)^\s*$`).Split(entryLog, -1)
}

func main() {
	declarationFormAnswers, err := input.ReadAll("./input")
	if err != nil {
		panic(err)
	}

	formGroups := splitEntries(declarationFormAnswers)

	fmt.Println("Day six, part one")
	// For each group, count the number of questions to which anyone answered "yes". What is the sum of those counts?
	answeredYes := 0
	for _, group := range formGroups {
		answersByPerson := strings.Split(group, "\n")

		groupAnswers := customs.ParseGroupAnswers(answersByPerson)

		answeredYes += len(groupAnswers)
	}

	fmt.Printf("Solution: %d\n", answeredYes)

	// For each group, count the number of questions to which everyone answered "yes". What is the sum of those counts?
	fmt.Println("Day six, part two")
	allAnsweredYes := 0
	for _, group := range formGroups {
		answersByPerson := strings.Split(strings.TrimSpace(group), "\n")

		var groupAnswerIntersection customs.FormAnswers

		for _, answers := range answersByPerson {
			formAnswers := customs.ParseIndividualAnswers(answers)

			if groupAnswerIntersection == nil {
				groupAnswerIntersection = formAnswers
				continue
			}

			// apply intersection of current form answers to group answers
			for question := range groupAnswerIntersection {
				if !formAnswers.Has(question) {
					// an individual did not answer yes to question, remove from group
					delete(groupAnswerIntersection, question)
				}
			}
		}

		allAnsweredYes += len(groupAnswerIntersection)
	}

	fmt.Printf("Solution: %d\n", allAnsweredYes)
}
