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
		`0 3 6 9 12 15
		1 3 6 10 15 21
		10 13 16 21 30 45`,
		114,
	},
	{
		``,
		2,
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
	t.Run("Works for provided example", func(t *testing.T) {
		want := testCases[2].want
		got := PartTwo(testCases[1].input)
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

func TestLineValuesFromString(t *testing.T) {
	input := "-1 0 1 2 3 4 5"
	want := LineValues{Values: []int{-1, 0, 1, 2, 3, 4, 5}}
	got := LineValuesFromString(input)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("LineValuesFromString(%v) = %v, want %v", input, got, want)
	}
}

func TestLineValues_FindNext(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"0 3 6 9 12 15", 18},
		{"1 3 6 10 15 21", 28},
		{"10 13 16 21 30 45", 68},
		{"-3 -1 1 3", 5},
		{"8 23 43 64 87 126 230 542 1426 3703 9058 20729 44688 91697 180910 346222 647728 1193456 2183213 4001351 7416521", 14007565},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			input := LineValuesFromString(tc.input)
			want := tc.want
			got := input.FindNext()
			if want != got {
				t.Errorf("LineValuesFromString(%v) = %v, want %v", input, got, want)
			}
		})
	}
}

func TestLineValues_FindPrevious(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"10 13 16 21 30 45", 5},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			input := LineValuesFromString(tc.input)
			want := tc.want
			got := input.FindPrevious()
			if want != got {
				t.Errorf("LineValuesFromString(%v) = %v, want %v", input, got, want)
			}
		})
	}
}
