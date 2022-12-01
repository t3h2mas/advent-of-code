package memgame

type previouslySpoken map[int]int

func newPreviouslySpoken() previouslySpoken {
	return make(previouslySpoken)
}

func (p previouslySpoken) Add(number int, turn int) {
	p[number] = turn
}

func (p previouslySpoken) LastSpoken(number int) (int, bool) {
	turn, hasNumber := p[number]
	return turn, hasNumber
}

func SpokenAt(startingNumbers []int, turn int) int {
	history := newPreviouslySpoken()

	var lastSpoken int
	currentTurn := 1
	for i := 0; i < len(startingNumbers); i++ {
		if i == len(startingNumbers)-1 {
			lastSpoken = startingNumbers[i]
		} else {
			history.Add(startingNumbers[i], currentTurn)
		}
		currentTurn++
	}

	for currentTurn <= turn {
		lastTurn := currentTurn - 1

		// look at when the previous number was spoken
		if turnSpoken, hasBeenSpoken := history.LastSpoken(lastSpoken); hasBeenSpoken {
			// the last number has been spoken before, get the difference between turns
			difference := lastTurn - turnSpoken
			history.Add(lastSpoken, lastTurn)
			lastSpoken = difference
		} else {
			// the previous speaking of the lastSpoken spoken was it's first
			history.Add(lastSpoken, lastTurn)
			lastSpoken = 0
		}

		currentTurn++
	}

	return lastSpoken
}
