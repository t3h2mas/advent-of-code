package tickets

type Ticket struct {
	values []int
}

func (t Ticket) Values() []int {
	return t.values
}
