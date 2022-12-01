package solution

import "testing"

func TestTicketScanningErrorRate(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "it should calculate the error rate for the example input",
			args: args{input: `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`},

			want: 71,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TicketScanningErrorRate(tt.args.input); got != tt.want {
				t.Errorf("TicketScanningErrorRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
