package tickets

type intRange struct {
	start int
	stop  int
}

func (ir intRange) contains(v int) bool {
	return v >= ir.start && v <= ir.stop
}

type Rule struct {
	field  string
	ranges []intRange
}

func (r Rule) Valid(value int) bool {
	for _, r := range r.ranges {
		if r.contains(value) {
			return true
		}
	}

	return false
}

func (r Rule) Field() string {
	return r.field
}
