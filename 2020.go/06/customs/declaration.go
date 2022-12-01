package customs

type FormAnswers map[rune]bool

func NewFormAnswers() FormAnswers {
	return make(map[rune]bool)
}

func (f FormAnswers) MarkYes(question rune) {
	f[question] = true
}

func (f FormAnswers) Has(question rune) bool {
	_, hasQuestion := f[question]
	return hasQuestion
}

func ParseGroupAnswers(answers []string) FormAnswers {
	result := NewFormAnswers()

	for _, answer := range answers {
		for _, question := range answer {
			result.MarkYes(question)
		}
	}

	return result
}

func ParseIndividualAnswers(answers string) FormAnswers {
	result := NewFormAnswers()

	for _, question := range answers {
		result.MarkYes(question)
	}

	return result
}
