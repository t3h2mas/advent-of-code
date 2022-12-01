package gameconsole

import (
	"testing"
)

func TestMachine_Load(t *testing.T) {
	t.Run("it should handle the example", func(t *testing.T) {
		input := []string{
			"nop +0",
			"acc +1",
			"jmp +4",
			"acc +3",
			"jmp -3",
			"acc -99",
			"acc +1",
			"jmp -4",
			"acc +6",
		}

		expected := 5

		m := NewMachine()
		_ = m.Load(input)

		got := m.Accumulator()

		if got != expected {
			t.Errorf("m.Accumulator() got = %d, want = %d", got, expected)
		}
	})
}

func TestMachine_LoadAndPatch(t *testing.T) {
	t.Run("it should find the accumulator after running patched instructions", func(t *testing.T) {
		input := []string{
			"nop +0",
			"acc +1",
			"jmp +4",
			"acc +3",
			"jmp -3",
			"acc -99",
			"acc +1",
			"jmp -4",
			"acc +6",
		}

		expected := 8

		m := LoadAndPatch(input)

		if m == nil {
			t.Error("failed to find correct instruction set")
		}

		got := m.Accumulator()

		if got != expected {
			t.Errorf("m.Accumulator() got = %d, want = %d", got, expected)
		}

	})
}
