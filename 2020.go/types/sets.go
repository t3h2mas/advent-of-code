package types

type IntSet map[int]bool

func NewIntSet() IntSet {
	return make(map[int]bool)
}

func (s IntSet) Has(val int) bool {
	_, exists := s[val]
	return exists
}

func (s IntSet) Add(val int) {
	s[val] = true
}
