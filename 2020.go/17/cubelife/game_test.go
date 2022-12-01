package cubelife

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseSeed(t *testing.T) {
	type args struct {
		seed string
	}
	tests := []struct {
		name string
		args args
		want []cube
	}{
		{
			name: "create a game with active cubes seeded from a string",
			args: args{
				seed: ".#.\n..#\n###",
			},
			want: []cube{
				{0, 0, 0},
				{1, 0, 0},
				{2, 0, 0},
				{2, 1, 0},
				{1, 2, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSeed(tt.args.seed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GameFromSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_Step(t *testing.T) {
	tests := []struct {
		name string
		game *Game
		want string
	}{
		{
			name: "should generate the next game state after a step",
			game: GameFromSeed(".#.\n..#\n###"),
			want: "z=-1\n#..\n..#\n.#.\n\nz=0\n#.#\n.##\n.#.\n\nz=1\n#..\n..#\n.#.\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.game.Step()

			got := tt.game.String()
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("cubes mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
