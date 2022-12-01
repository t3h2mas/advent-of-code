package graph

type digraphNode struct {
	value string
	edges []*digraphEdge
}

func newDigraphNode(val string) *digraphNode {
	return &digraphNode{
		value: val,
	}
}

func (n *digraphNode) addEdge(e *digraphEdge) {
	n.edges = append(n.edges, e)
}

func (n *digraphNode) Edges() []*digraphEdge {
	return n.edges
}

func (n *digraphNode) Value() string {
	return n.value
}

type digraphEdge struct {
	dest   *digraphNode
	weight int
}

func newDigraphEdge(dest *digraphNode, weight int) *digraphEdge {
	return &digraphEdge{
		dest,
		weight,
	}
}

func (e *digraphEdge) Destination() *digraphNode {
	return e.dest
}

func (e *digraphEdge) Weight() int {
	return e.weight
}

type Digraph struct {
	nodes map[string]*digraphNode
}

func NewDigraph() *Digraph {
	return &Digraph{
		make(map[string]*digraphNode),
	}
}

func (g *Digraph) Get(val string) *digraphNode {
	if node, exists := g.nodes[val]; exists {
		return node
	}

	node := newDigraphNode(val)

	g.nodes[val] = node

	return node
}

func (g *Digraph) AddEdge(source string, destination string, weight int) {
	g.Get(source).addEdge(
		newDigraphEdge(
			g.Get(destination),
			weight,
		),
	)
}

func (g *Digraph) Nodes() map[string]*digraphNode {
	return g.nodes
}

func ComputeAggregateWeight(node *digraphNode) int {
	memo := make(map[string]int)

	return computeAggregateWeightForNode(node, memo)
}

func computeAggregateWeightForNode(node *digraphNode, memo map[string]int) int {
	if len(node.Edges()) == 0 {
		return 0
	}

	if result, hasValue := memo[node.Value()]; hasValue {
		return result
	}

	total := 0

	for _, edge := range node.Edges() {
		total += edge.Weight() + edge.Weight()*computeAggregateWeightForNode(edge.Destination(), memo)
	}

	memo[node.Value()] = total

	return total
}
