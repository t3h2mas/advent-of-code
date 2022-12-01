package graph

import "testing"

func TestComputeAggregateWeight(t *testing.T) {
	type args struct {
		node *digraphNode
	}
	type testCase struct {
		name string
		args args
		want int
	}

	t.Run("it should compute the first example", func(t *testing.T) {
		graph := NewDigraph()
		graph.AddEdge("shiny gold", "dark olive", 1)
		graph.AddEdge("shiny gold", "vibrant plum", 2)
		graph.AddEdge("vibrant plum", "faded blue", 5)
		graph.AddEdge("vibrant plum", "dotted black", 6)
		graph.AddEdge("dark olive", "faded blue", 3)
		graph.AddEdge("dark olive", "dotted black", 4)

		target := graph.Get("shiny gold")

		test := testCase{
			args: args{
				target,
			},
			want: 32,
		}
		if got := ComputeAggregateWeight(test.args.node); got != test.want {
			t.Errorf("ComputeAggregateWeight() = %v, want %v", got, test.want)
		}
	})

	t.Run("it should compute the second example", func(t *testing.T) {
		graph := NewDigraph()
		graph.AddEdge("shiny gold", "dark red", 2)
		graph.AddEdge("dark red", "dark orange", 2)
		graph.AddEdge("dark orange", "dark yellow", 2)
		graph.AddEdge("dark yellow", "dark green", 2)
		graph.AddEdge("dark green", "dark blue", 2)
		graph.AddEdge("dark blue", "dark violet", 2)

		target := graph.Get("shiny gold")

		test := testCase{
			args: args{
				target,
			},
			want: 126,
		}
		if got := ComputeAggregateWeight(test.args.node); got != test.want {
			t.Errorf("ComputeAggregateWeight() = %v, want %v", got, test.want)
		}
	})
}
