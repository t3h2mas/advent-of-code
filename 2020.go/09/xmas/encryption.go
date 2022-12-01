package xmas

type Message struct {
	packets        []int
	preambleLength int
}

func NewMessage(packets []int, preambleLength int) *Message {
	return &Message{
		packets,
		preambleLength,
	}
}

type intSet map[int]bool

func (s intSet) Has(val int) bool {
	_, exists := s[val]
	return exists
}

func (s intSet) Add(val int) {
	s[val] = true
}

func newIntSet() intSet {
	return make(map[int]bool)
}

// Find the first number in the list after the preamble that is not the sum of two
// numbers in the preamble
func (m *Message) Crack() int {
	for i := m.preambleLength; i < len(m.packets); i++ {
		// build preamble set
		inPreamble := newIntSet()

		check := m.packets[i]
		for j := i - m.preambleLength; j < i; j++ {
			inPreamble.Add(m.packets[j])
		}

		found := false
		for j := i - m.preambleLength; j < i; j++ {
			if inPreamble.Has(check - m.packets[j]) {
				found = true
				break
			}
		}

		if !found {
			return check
		}

	}
	return 0
}

func (m *Message) Exploit(cracked int) int {
	for i := 0; i < len(m.packets); i++ {
		min := m.packets[i]
		max := m.packets[i]
		sum := 0
		for j := i; j < len(m.packets); j++ {
			sum += m.packets[j]
			if m.packets[j] > max {
				max = m.packets[j]
			}

			if sum == cracked {
				return min + max
			}

			if sum > cracked {
				break
			}
		}
	}
	return 0
}
