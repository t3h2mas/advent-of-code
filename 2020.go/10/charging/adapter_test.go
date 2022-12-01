package charging

import "testing"

func TestChainVoltageDifferences(t *testing.T) {
	type args struct {
		adapters []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "it should return the product of the differences of 1,3",
			args: args{adapters: []int{
				16,
				10,
				15,
				5,
				1,
				11,
				7,
				19,
				6,
				12,
				4,
			}},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChainVoltageDifferences(tt.args.adapters); got != tt.want {
				t.Errorf("ChainVoltageDifferences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChainPossibilities(t *testing.T) {
	type args struct {
		adapters []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "it should calculate how many possible routes there are",
			args: args{
				adapters: []int{
					16,
					10,
					15,
					5,
					1,
					11,
					7,
					19,
					6,
					12,
					4,
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChainPossibilities(tt.args.adapters); got != tt.want {
				t.Errorf("ChainPossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}
