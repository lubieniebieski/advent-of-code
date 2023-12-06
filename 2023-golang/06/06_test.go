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
		`Time:      7  15   30
		Distance:  9  40  200`,
		288,
	},
	{
		``,
		46,
	},
}

func TestPartOne(t *testing.T) {
	want := testCases[1].want
	got := PartOne(testCases[1].input)
	toolstest.CompareWithExample(t, want, got)
}
func TestPartTwo(t *testing.T) {
	t.SkipNow()
	want := testCases[2].want
	got := PartTwo(testCases[1].input)
	toolstest.CompareWithExample(t, want, got)
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

func TestParseRaces(t *testing.T) {
	want := []Race{
		{Time: 7, Distance: 9},
		{Time: 15, Distance: 40},
		{Time: 30, Distance: 200},
	}
	if !reflect.DeepEqual(ParseRaces(testCases[1].input), want) {
		t.Errorf("ParseRaces() = %v, want %v", ParseRaces(testCases[1].input), want)
	}

}

func TestWinningOptionsCount(t *testing.T) {
	testCases := []struct {
		race Race
		want int
	}{
		{race: Race{Time: 7, Distance: 9}, want: 4},
		{race: Race{Time: 15, Distance: 40}, want: 8},
		{race: Race{Time: 30, Distance: 200}, want: 9},
	}
	for _, tc := range testCases {
		got := tc.race.WinningOptionsCount()
		if tc.want != got {
			t.Errorf("WinningOptionsCount() = %v, want %v", got, tc.want)
		}
	}

}
