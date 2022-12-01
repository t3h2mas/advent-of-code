package seats

import (
	"reflect"
	"testing"
)

func TestLayoutFrom(t *testing.T) {
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want Layout
	}{
		{
			name: "should build a layout from a multiline grid",
			args: args{
				grid: []string{
					"L.L",
					"LLL",
				},
			},
			want: Layout{
				height:   2,
				width:    3,
				occupied: map[Point]bool{},
				floor: map[Point]bool{
					{1, 0}: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LayoutFrom(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LayoutFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
