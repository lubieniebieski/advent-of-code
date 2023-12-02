package main

import (
	"testing"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools/toolstest"
)

// Input for tests from Readme
var testCases = []struct {
	input string
	want  int
}{
	{
		"First is for fun!",
		1337,
	},
	{
		`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
		Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
		Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
		Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
		8,
	},
	{
		`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
		Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
		Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
		Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
		2286,
	},
}

func TestPartOne(t *testing.T) {
	want := testCases[1].want
	got := PartOne(testCases[1].input)
	toolstest.CompareWithExample(t, want, got)
}

func TestPartTwo(t *testing.T) {
	want := testCases[2].want
	got := PartTwo(testCases[2].input)
	toolstest.CompareWithExample(t, want, got)
}

// Main tests below

func TestNewGame(t *testing.T) {
	input := "Game 123: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	g := NewGame(input)
	t.Run("sets the Id correctly", func(t *testing.T) {

		want := 123
		if g.Id != want {
			t.Errorf("Parse() = %d, want %d", g.Id, want)
		}
	})
}

func TestIsValid(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		g := NewGame("Game 123: 223 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
		want := false
		if g.IsValid() != want {
			t.Errorf("IsValid() = %v, want %v", g.IsValid(), want)
		}
	})
	t.Run("valid", func(t *testing.T) {
		g := NewGame("Game 123: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
		want := true
		if g.IsValid() != want {
			t.Errorf("IsValid() = %v, want %v", g.IsValid(), want)
		}
	})
}

func TestValidAmount(t *testing.T) {
	cases := []struct {
		amount int
		color  string
		want   bool
	}{
		{1, "red", true},
		{13, "red", false},
		{1, "green", true},
		{14, "green", false},
		{1, "blue", true},
		{15, "blue", false},
	}
	for _, c := range cases {
		t.Run("For "+c.color, func(t *testing.T) {
			if ValidAmount(c.amount, c.color) != c.want {
				t.Errorf("ValidAmount() = %v, want %v", ValidAmount(c.amount, c.color), c.want)
			}
		})
	}
}
