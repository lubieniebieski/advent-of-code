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
	{
		``,
		8410,
	},
}

func TestPartOne(t *testing.T) {
	t.Run("Works for provided example", func(t *testing.T) {
		want := testCases[1].want
		got := PartOne(testCases[1].input)
		toolstest.CompareWithExample(t, want, got)
	})

}

func TestPartTwo(t *testing.T) {
	t.Run("Works for provided example x100", func(t *testing.T) {
		want := 8410
		got := Calculate(testCases[1].input, 100)
		toolstest.CompareWithExample(t, want, got)
	})

	t.Run("Works for provided example x10", func(t *testing.T) {
		want := 1030
		got := Calculate(testCases[1].input, 10)
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
					{value: "#"},
					{value: "."},
				},
				{
					{value: "."},
					{value: "."},
					{value: "."},
				},

				{
					{value: "."},
					{value: "#"},
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

func TestGrid_ShortestPathBetweenGalaxies(t *testing.T) {
	t.Run("Works for simple example", func(t *testing.T) {
		input := `#..
		.#.
		..#`
		grid := GridFromString(input)
		want := 4
		got := grid.ShortestPathBetweenGalaxies(Point{0, 0}, Point{2, 2}, 1)
		if got != want {
			t.Errorf("That's not what I wanted! Got: %v, want: %v", got, want)
		}
	})

	t.Run("Works for example from readme", func(t *testing.T) {
		grid := GridFromString(testCases[1].input)
		tCases := []struct {
			from Point
			to   Point
			want int
		}{
			{Point{0, 9}, Point{4, 9}, 5},
		}
		for _, tCase := range tCases {
			got := grid.ShortestPathBetweenGalaxies(tCase.from, tCase.to, 1)
			if got != tCase.want {
				t.Errorf("That's not what I wanted! Got: %v, want: %v", got, tCase.want)
			}
		}
	})
}
