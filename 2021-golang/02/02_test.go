package main

import (
	"testing"

	"github.com/lubieniebieski/advent-of-code/2022/tools"
)

func TestPartOne(t *testing.T) {
	input := `forward 5
	down 5
	forward 8
	up 3
	down 8
	forward 2
`
	want := 150
	strings := tools.ExtractStringsFromString(input)
	commands := ConvertStringsToCommand(strings)
	got, _ := PartOne(commands)
	compareWithExample(t, want, got)
}

func TestConvertToCommand(t *testing.T) {
	var convertTests = []struct {
		input    string
		expected Command
	}{
		{"forward 1", Command{"forward", 1}},
		{"down 5", Command{"down", 5}},
		{"up 15", Command{"up", 15}},
		{"back 2", Command{"back", 2}},
	}
	for _, test := range convertTests {
		got := ConvertToCommand(test.input)

		if got != test.expected {
			t.Errorf("It should be converted to %v, but got %v instead", test.expected, got)
		}
	}

}

func TestHorizontalValue(t *testing.T) {
	t.Run("works for 'down' value", func(t *testing.T) {
		command := Command{"down", 4}
		want := 0
		got := command.Horizontal()
		if got != want {
			t.Errorf("Value should be %v, but got %v instead", want, got)
		}
	})
	t.Run("works for 'up' value", func(t *testing.T) {
		command := Command{"up", 4}
		want := 0
		got := command.Horizontal()
		if got != want {
			t.Errorf("Value should be %v, but got %v instead", want, got)
		}
	})
	t.Run("works for 'forward' value", func(t *testing.T) {
		command := Command{"forward", 4}
		want := 4
		got := command.Horizontal()
		if got != want {
			t.Errorf("Value should be %v, but got %v instead", want, got)
		}
	})
	t.Run("works for 'back' value", func(t *testing.T) {
		command := Command{"back", 4}
		want := -4
		got := command.Horizontal()
		if got != want {
			t.Errorf("Value should be %v, but got %v instead", want, got)
		}
	})

}

func TestVerticalValue(t *testing.T) {
	t.Run("works for 'up' value", func(t *testing.T) {
		command := Command{"up", 4}
		want := -4
		got := command.Vertical()
		if got != want {
			t.Errorf("Value should be %v, but got %v instead", want, got)
		}
	})

	t.Run("works for 'down' value", func(t *testing.T) {
		command := Command{"down", 4}
		want := 4
		got := command.Vertical()
		if got != want {
			t.Errorf("Value should be %v, but got %v instead", want, got)
		}
	})
	t.Run("works for 'forward' value", func(t *testing.T) {
		command := Command{"forward", 4}
		want := 0
		got := command.Vertical()
		if got != want {
			t.Errorf("Value should be %v, but got %v instead", want, got)
		}
	})
	t.Run("works for 'back' value", func(t *testing.T) {
		command := Command{"back", 4}
		want := 0
		got := command.Vertical()
		if got != want {
			t.Errorf("Value should be %v, but got %v instead", want, got)
		}
	})
}

func compareWithExample(t *testing.T, want, got int) {
	t.Helper()

	if got != want {
		t.Errorf("According to the example, we should have %d, but we've got %d", want, got)
	}
}
