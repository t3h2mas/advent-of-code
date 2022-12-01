package solution

import (
	"strings"

	"github.com/t3h2mas/advent-2020/types"

	"github.com/t3h2mas/advent-2020/16/tickets"
)

func TicketScanningErrorRate(input string) int {
	parts := strings.Split(input, "\n\n")

	rules := tickets.ParseRules(parts[0])
	tickets := tickets.ParseTickets(parts[2])
	errorRate := 0

	for _, ticket := range tickets {
		for _, value := range ticket.Values() {
			validAtLeastOnce := false
			for _, rule := range rules {
				if rule.Valid(value) {
					validAtLeastOnce = true
					break
				}
			}

			if !validAtLeastOnce {
				errorRate += value
			}
		}
	}

	return errorRate
}

func TicketDepartureProduct(input string) int {
	parts := strings.Split(input, "\n\n")

	rules := tickets.ParseRules(parts[0])
	targetTicket := tickets.ParseTickets(parts[1])[0]
	otherTickets := tickets.ParseTickets(parts[2])

	validTickets := []tickets.Ticket{targetTicket}

	// filter out invalid tickets
	for _, ticket := range otherTickets {
		if ticketIsValid(ticket, rules) {
			validTickets = append(validTickets, ticket)
		}
	}

	// Track which columns are valid for a particular rule
	rulesValidInColumns := make(map[string]types.IntSet)

	for _, rule := range rules {
		validInColumn := make(map[int]int)
		for _, ticket := range validTickets {
			for idx, value := range ticket.Values() {
				if rule.Valid(value) {
					validInColumn[idx] = validInColumn[idx] + 1
				}
			}
		}

		for column, count := range validInColumn {
			if count == len(validTickets) {
				if rulesValidInColumns[rule.Field()] == nil {
					rulesValidInColumns[rule.Field()] = types.NewIntSet()
				}
				rulesValidInColumns[rule.Field()].Add(column)
			}
		}
	}

	ruleColumns := make(map[string]int)

	// find the rule with only one valid column, assign the column,
	// remove the assigned column from the valid columns of all rules,
	// repeat until all rules have been assigned
	for len(ruleColumns) != len(rules) {
		// find the rule with only one option, and assign it
		for ruleField, validColumns := range rulesValidInColumns {
			if _, assigned := ruleColumns[ruleField]; assigned {
				continue
			}

			if len(validColumns) == 1 {
				for validColumn := range validColumns {
					ruleColumns[ruleField] = validColumn

					for _, vcs := range rulesValidInColumns {
						delete(vcs, validColumn)
					}
				}
			}
		}
	}

	product := 1

	for ruleField, column := range ruleColumns {
		if strings.HasPrefix(ruleField, "departure") {
			product *= targetTicket.Values()[column]
		}
	}

	return product
}

func ticketIsValid(ticket tickets.Ticket, rules []tickets.Rule) bool {
	for _, value := range ticket.Values() {
		validAtLeastOnce := false
		for _, rule := range rules {
			if rule.Valid(value) {
				validAtLeastOnce = true
				break
			}
		}
		if !validAtLeastOnce {
			return false
		}
	}
	return true
}
