package tickets

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []Rule
	}{
		{
			name: "it should parse the rules string",
			args: args{input: strings.Join([]string{
				"class: 1-3 or 5-7",
				"row: 6-11 or 33-44",
				"seat: 13-40 or 45-50",
				"departure time: 42-43 or 68-72",
			}, "\n")},
			want: []Rule{
				{
					field: "class",
					ranges: []intRange{
						{1, 3},
						{5, 7},
					},
				},
				{
					field: "row",
					ranges: []intRange{
						{6, 11},
						{33, 44},
					},
				},
				{
					field: "seat",
					ranges: []intRange{
						{13, 40},
						{45, 50},
					},
				},
				{
					field: "departure time",
					ranges: []intRange{
						{42, 43},
						{68, 72},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseRules(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseTickets(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []Ticket
	}{
		{
			name: "it should parse the tickets string",
			args: args{
				input: strings.Join([]string{
					"nearby tickets:",
					"7,3,47",
					"40,4,50",
					"55,2,20",
					"38,6,12"}, "\n"),
			},
			want: []Ticket{
				{[]int{7, 3, 47}},
				{[]int{40, 4, 50}},
				{[]int{55, 2, 20}},
				{[]int{38, 6, 12}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseTickets(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTickets() = %v, want %v", got, tt.want)
			}
		})
	}
}
