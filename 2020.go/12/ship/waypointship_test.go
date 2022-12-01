package ship

import (
	"reflect"
	"testing"

	"github.com/t3h2mas/advent-2020/12/action"
)

func TestWaypointShip_Act(t *testing.T) {
	actions := []*action.Action{
		action.NewActionFrom("F10"),
		action.NewActionFrom("N3"),
		action.NewActionFrom("F7"),
		action.NewActionFrom("R90"),
		action.NewActionFrom("F11"),
	}

	ws := NewWaypointShip()

	for _, a := range actions {
		ws.Act(a.Type(), a.Units())
	}

	got := ws.ShipPosition()
	want := Position{-214, -72}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %#v, want %#v", got, want)
	}
}
