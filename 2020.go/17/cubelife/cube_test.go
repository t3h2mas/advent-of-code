package cubelife

import (
	"testing"
)

func TestCube_Neighbors_returns_26_cubes(t *testing.T) {
	c := cube{
		0, 0, 0,
	}

	neighbors := c.Neighbors()
	want := 26
	got := len(neighbors)

	if got != want {
		t.Errorf("len(c.Neighbors) = %d, want %d", got, want)
	}
}
