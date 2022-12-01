package bus

import (
	"reflect"
	"testing"
)

func TestBusesFrom(t *testing.T) {
	input := "7,13,x,x,59,x,31,19"
	want := []Bus{
		{7},
		{13},
		{59},
		{31},
		{19},
	}
	if got := BusesFrom(input); !reflect.DeepEqual(got, want) {
		t.Errorf("BusesFrom() = %v, want %v", got, want)
	}
}

func TestIndexedBusesFrom(t *testing.T) {
	input := "7,13,x,x,59,x,31,19"
	want := map[int]Bus{
		0: {7},
		1: {13},
		4: {59},
		6: {31},
		7: {19},
	}
	if got := IndexedBusesFrom(input); !reflect.DeepEqual(got, want) {
		t.Errorf("BusesFrom() = %v, want %v", got, want)
	}
}
