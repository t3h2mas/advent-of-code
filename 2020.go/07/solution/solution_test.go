package solution

import (
	"reflect"
	"testing"

	"github.com/t3h2mas/advent-2020/07/graph"
)

func TestLineParts(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "it parses a line with one contains clause",
			args: args{
				line: "bright white bags contain 1 shiny gold bag.",
			},
			want: [][]string{
				{"bright", "white", "bags"},
				{"1", "shiny", "gold", "bag."},
			},
		},
		{
			name: "it parses a line with two contains clauses",
			args: args{
				line: "light red bags contain 1 bright white bag, 2 muted yellow bags.",
			},
			want: [][]string{
				{"light", "red", "bags"},
				{"1", "bright", "white", "bag"},
				{"2", "muted", "yellow", "bags."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LineParts(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LineParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBagClauses(t *testing.T) {
	type args struct {
		parts [][]string
	}
	tests := []struct {
		name string
		args args
		want *Bag
	}{
		{
			name: "gets clauses for a bag",
			args: args{
				[][]string{
					{"light", "red", "bags"},
					{"1", "bright", "white", "bag"},
					{"2", "muted", "yellow", "bags."},
				},
			},
			want: &Bag{
				Color: "light red",
				Contains: []string{
					"bright white",
					"muted yellow",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BagClauses(tt.args.parts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BagClauses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraphFromInput(t *testing.T) {
	input := []string{
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}

	expected := graph.NewDigraph()
	expected.AddEdge("shiny gold", "dark olive", 1)
	expected.AddEdge("shiny gold", "vibrant plum", 2)
	expected.AddEdge("vibrant plum", "faded blue", 5)
	expected.AddEdge("vibrant plum", "dotted black", 6)
	expected.AddEdge("dark olive", "faded blue", 3)
	expected.AddEdge("dark olive", "dotted black", 4)

	got := GraphFromInput(input)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("GraphFromInput() = %#v\nwant %#v", got, expected)
	}
}
