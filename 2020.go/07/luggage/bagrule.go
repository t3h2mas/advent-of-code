package luggage

type BagRule struct {
	Color       string
	ContainedBy map[string]*BagRule
}

func NewBagRule(color string) *BagRule {
	return &BagRule{
		Color:       color,
		ContainedBy: make(map[string]*BagRule),
	}
}

type BagRules struct {
	rules map[string]*BagRule
}

func NewBagRules() *BagRules {
	return &BagRules{
		make(map[string]*BagRule),
	}
}

func (rs *BagRules) Get(color string) *BagRule {
	// return if exists
	if bagRule, exists := rs.rules[color]; exists {
		return bagRule
	}

	// otherwise create and return
	bagRule := NewBagRule(color)
	rs.rules[color] = bagRule

	return bagRule
}

func (rs *BagRules) AddRule(color, containedBy string) {
	bagRule := rs.Get(color)

	bagRule.ContainedBy[containedBy] = rs.Get(containedBy)
}
