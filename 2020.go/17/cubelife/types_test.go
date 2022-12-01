package cubelife

import (
	"reflect"
	"testing"
)

func Test_cubeTracker_Add(t *testing.T) {
	type args struct {
		c cube
	}
	tests := []struct {
		name string
		args args
		want cubeTracker
	}{
		{
			name: "it should add a cube, and update the boundaries",
			args: args{
				c: cube{1, 2, 3},
			},
			want: cubeTracker{
				cubes: cubeSet{
					{1, 2, 3}: true,
				},
				boundaries: boundaries{
					x: &limit{high: 1, low: 0},
					y: &limit{high: 2, low: 0},
					z: &limit{high: 3, low: 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct := newCubeTracker()

			ct.Add(tt.args.c)

			if cubes := ct.cubes; !reflect.DeepEqual(cubes, tt.want.cubes) {
				t.Errorf("got = %v, want %v", cubes, tt.want)
			}

			if boundaries := ct.boundaries; !reflect.DeepEqual(boundaries, tt.want.boundaries) {
				t.Errorf("got = %v, want %v", boundaries, tt.want)
			}
		})
	}
}
