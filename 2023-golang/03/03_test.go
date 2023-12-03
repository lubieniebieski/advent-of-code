package main

import (
	"reflect"
	"slices"
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
		`467..114..
		...*......
		..35..633.
		......#...
		617*......
		.....+.58.
		..592.....
		......755.
		...$.*....
		.664.598..`,
		4361,
	},
	{
		`467..114..
		...*......
		..35..633.
		......#...
		617*......
		.....+.58.
		..592.....
		......755.
		...$.*....
		.664.598..`,
		467835,
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

func TestFindPart(t *testing.T) {
	want := []Part{
		{PartNumber: 3, x: 4, y: 0},
		{PartNumber: 313, x: 13, y: 0},
	}
	got := FindParts("..te3st..work313", 0)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindParts() = %v, want %v", got, want)
	}
	if len(got) != len(want) {
		t.Errorf("FindParts() = %v, want %v", got, want)
	}
}

func TestPart_AdjecentToSymbol(t *testing.T) {
	testCases := []struct {
		name  string
		input []string
		want  bool
	}{
		{
			name:  "there is no symbol",
			input: []string{"testte", "t111st", "testte"},
			want:  false,
		},
		{
			name:  "there is symbol on the right",
			input: []string{"testte", "t111*t", "testte"},
			want:  true,
		},
		{
			name:  "there is symbol on the left",
			input: []string{"testte", "*111st", "testte"},
			want:  true,
		},
		{
			name:  "there is symbol on the top",
			input: []string{"tes*te", "t111st", "testte"},
			want:  true,
		},
		{
			name:  "there is symbol on the bottom",
			input: []string{"testte", "t111st", "tes*te"},
			want:  true,
		},
		{
			name:  "there is symbol on the top left",
			input: []string{"*estte", "t111st", "testte"},
			want:  true,
		},
		{
			name:  "there is symbol on the top right",
			input: []string{"test*e", "t111st", "testte"},
			want:  true,
		},
		{
			name:  "there is symbol on the bottom left",
			input: []string{"testte", "t111st", "*estte"},
			want:  true,
		},
		{
			name:  "there is symbol on the bottom right",
			input: []string{"testte", "t111st", "test*e"},
			want:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			part := Part{PartNumber: 111, x: 1, y: 1}
			symbols := []string{"*"}
			got := part.AdjecentToSymbol(tc.input, symbols)
			if got != tc.want {
				t.Errorf("AdjecentToSymbol() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFindSymbols(t *testing.T) {
	t.Run("should return empty array if there is no symbol", func(t *testing.T) {
		input := []string{"...12", "122...."}
		got := FindSymbols(input)
		want := []string{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("FindSymbols() = %v, want %v", got, want)
		}
	})
	t.Run("should return array with one symbol", func(t *testing.T) {
		input := []string{"1212...", "4..4.....*"}
		got := FindSymbols(input)
		want := []string{"*"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("FindSymbols() = %v, want %v", got, want)
		}
	})
	t.Run("should return symbol only once", func(t *testing.T) {
		input := []string{"..*", "...*", "....*"}
		got := FindSymbols(input)
		want := []string{"*"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("FindSymbols() = %v, want %v", got, want)
		}
	})
	t.Run("should return array with two symbols", func(t *testing.T) {
		input := []string{"3..*", "2..*", "22.&"}
		got := FindSymbols(input)
		want := []string{"*", "&"}
		slices.Contains(got, "*")
		if !slices.Contains(got, "*") || !slices.Contains(got, "&") {
			t.Errorf("FindSymbols() = %v, want %v", got, want)
		}
	})
}

func TestIsSymbol(t *testing.T) {
	t.Run("should return false if there is no symbol", func(t *testing.T) {
		input := "test"
		symbols := []string{"*"}
		got := IsSymbol(input, symbols)
		want := false
		if got != want {
			t.Errorf("IsSymbol() = %v, want %v", got, want)
		}
	})
	t.Run("should return true if there is symbol", func(t *testing.T) {
		input := "te*st"
		symbols := []string{"*"}
		got := IsSymbol(input, symbols)
		want := true
		if got != want {
			t.Errorf("IsSymbol() = %v, want %v", got, want)
		}
	})

}
