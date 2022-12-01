package seats

import (
	"reflect"
	"testing"
)

func Test_calculateOccupied(t *testing.T) {
	type args struct {
		layout Layout
	}
	tests := []struct {
		name string
		args args
		want Layout
	}{
		{
			name: "should calculate the next occupied seat state starting empty",
			args: args{
				layout: LayoutFrom(
					[]string{
						"L.LL",
						"LLLL",
						"L.L.",
					},
				),
			},
			want: LayoutFrom(
				[]string{
					"#.##",
					"####",
					"#.#.",
				},
			),
		},
		{
			name: "should calculate the next occupied seat state from a given grid",
			args: args{
				layout: LayoutFrom(
					[]string{
						"#.##.##.##",
						"#######.##",
						"#.#.#..#..",
						"####.##.##",
						"#.##.##.##",
						"#.#####.##",
						"..#.#.....",
						"##########",
						"#.######.#",
						"#.#####.##",
					},
				),
			},
			want: LayoutFrom(
				[]string{
					"#.LL.L#.##",
					"#LLLLLL.L#",
					"L.L.L..L..",
					"#LLL.LL.L#",
					"#.LL.LL.LL",
					"#.LLLL#.##",
					"..L.L.....",
					"#LLLLLLLL#",
					"#.LLLLLL.L",
					"#.#LLLL.##",
				},
			),
		},
		{
			name: "should calculate the same state when no changes can be applied",
			args: args{
				LayoutFrom(
					[]string{
						"#.#L.L#.##",
						"#LLL#LL.L#",
						"L.#.L..#..",
						"#L##.##.L#",
						"#.#L.LL.LL",
						"#.#L#L#.##",
						"..L.L.....",
						"#L#L##L#L#",
						"#.LLLLLL.L",
						"#.#L#L#.##",
					},
				),
			},
			want: LayoutFrom(
				[]string{
					"#.#L.L#.##",
					"#LLL#LL.L#",
					"L.#.L..#..",
					"#L##.##.L#",
					"#.#L.LL.LL",
					"#.#L#L#.##",
					"..L.L.....",
					"#L#L##L#L#",
					"#.LLLLLL.L",
					"#.#L#L#.##",
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateOccupied(tt.args.layout); !reflect.DeepEqual(got, tt.want.occupied) {
				t.Errorf("calculateOccupied() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimulateStableOccupiedSeatCount(t *testing.T) {
	type args struct {
		layout Layout
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should calculate seat occupancy after the seats have stabalized",
			args: args{
				LayoutFrom(
					[]string{
						"L.LL.LL.LL",
						"LLLLLLL.LL",
						"L.L.L..L..",
						"LLLL.LL.LL",
						"L.LL.LL.LL",
						"L.LLLLL.LL",
						"..L.L.....",
						"LLLLLLLLLL",
						"L.LLLLLL.L",
						"L.LLLLL.LL",
					},
				),
			},
			want: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimulateStableOccupiedSeatCount(tt.args.layout); got != tt.want {
				t.Errorf("SimulateStableOccupiedSeatCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
