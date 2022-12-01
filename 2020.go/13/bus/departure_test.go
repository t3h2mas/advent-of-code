package bus

import (
	"reflect"
	"testing"
)

func TestClosestDeparture(t *testing.T) {
	buses := []Bus{
		{7},
		{13},
		{59},
		{31},
		{19},
	}
	time := 939
	want := Departure{
		busID: 59,
		time:  944,
	}
	if got := ClosestDeparture(buses, time); !reflect.DeepEqual(got, want) {
		t.Errorf("ClosestDeparture() = %v, want %v", got, want)
	}
}

func TestEarliestOrderedDeparture(t *testing.T) {
	type args struct {
		buses map[int]Bus
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example one",
			args: args{
				buses: map[int]Bus{
					0: {7},
					1: {13},
					4: {59},
					6: {31},
					7: {19},
				},
			},
			want: 1068781,
		},
		{
			name: "example two",
			args: args{
				buses: map[int]Bus{
					0: {17},
					2: {13},
					3: {19},
				},
			},
			want: 3417,
		},
		{
			name: "example three",
			args: args{
				buses: map[int]Bus{
					0: {67},
					1: {7},
					2: {59},
					3: {61},
				},
			},
			want: 754018,
		},
		{
			name: "example four (large)",
			args: args{
				buses: map[int]Bus{
					0: {1789},
					1: {37},
					2: {47},
					3: {1889},
				},
			},
			want: 1202161486,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EarliestOrderedDeparture(tt.args.buses); got != tt.want {
				t.Errorf("EarliestOrderedDeparture() = %v, want %v", got, tt.want)
			}
		})
	}
}
