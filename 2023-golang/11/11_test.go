package main

import (
	"reflect"
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
		`...#......
		.......#..
		#.........
		..........
		......#...
		.#........
		.........#
		..........
		.......#..
		#...#.....,`,
		374,
	},
}

func TestPartOne(t *testing.T) {
	t.SkipNow()
	t.Run("Works for provided example", func(t *testing.T) {
		want := testCases[1].want
		got := PartOne(testCases[1].input)
		toolstest.CompareWithExample(t, want, got)
	})

}

func TestPartTwo(t *testing.T) {
	t.SkipNow()
	t.Run("Works for provided example", func(t *testing.T) {
		want := testCases[3].want
		got := PartTwo(testCases[3].input)
		toolstest.CompareWithExample(t, want, got)
	})
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartOne(testCases[1].input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartTwo(testCases[1].input)
	}
}

// Main tests below

func TestGridFromString(t *testing.T) {
	t.Run("Works for example with galaxies", func(t *testing.T) {
		input := `#..
		.#.
		..#`
		want := Grid{
			tiles: [][]Tile{
				{
					{value: "#"},
					{value: "."},
					{value: "."},
				},
				{
					{value: "."},
					{value: "#"},
					{value: "."},
				},
				{
					{value: "."},
					{value: "."},
					{value: "#"},
				},
			},
		}
		got := GridFromString(input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("That's not what I wanted! Got: %v, want: %v", got, want)

		}
	})

	t.Run("Works for example without galaxies", func(t *testing.T) {
		input := `.#.
				  ...
		          .#.`
		want := Grid{
			tiles: [][]Tile{
				{
					{value: "."},
					{value: "."},
					{value: "#"},
					{value: "."},
					{value: "."},
				},
				{
					{value: "."},
					{value: "."},
					{value: "."},
					{value: "."},
					{value: "."},
				},
				{
					{value: "."},
					{value: "."},
					{value: "."},
					{value: "."},
					{value: "."},
				},
				{
					{value: "."},
					{value: "."},
					{value: "#"},
					{value: "."},
					{value: "."},
				},
			},
		}
		got := GridFromString(input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("That's not what I wanted! Got: %v, want: %v", got, want)

		}
	})
}
