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
		`32T3K 765
		T55J5 684
		KK677 28
		KTJJT 220
		QQQJA 483`,
		6440,
	},
	{
		``,
		1,
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

func TestCardFromString(t *testing.T) {
	want := Hand{Rank: 765, StrValue: "32T3K", CalculatedStrength: 1}
	got := HandFromString("32T3K 765")
	toolstest.CompareWithExample(t, want, got)
}

func TestHand_Strength(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"AAAAA", FiveOfAKind},
		{"AA8AA", FourOfAKind},
		{"23332", FullHouse},
		{"TTT98", ThreeOfAKind},
		{"23432", TwoPairs},
		{"A23A4", OnePair},
		{"23456", HighCard},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := Hand{StrValue: tc.input}.Strength()
			toolstest.CompareWithExample(t, tc.want, got)
		})
	}
}

func TestSortHandsByStrength(t *testing.T) {
	input := []Hand{
		{StrValue: "32T3K", CalculatedStrength: 1},
		{StrValue: "T55J5", CalculatedStrength: 4},
		{StrValue: "KK677", CalculatedStrength: 3},
		{StrValue: "KTJJT", CalculatedStrength: 2},
		{StrValue: "QQQJA", CalculatedStrength: 5},
	}
	want := []Hand{
		{StrValue: "32T3K"},
		{StrValue: "KTJJT"},
		{StrValue: "KK677"},
		{StrValue: "T55J5"},
		{StrValue: "QQQJA"},
	}

	got := SortHandsByStrength(input)
	for i, hand := range got {
		if hand.StrValue != want[i].StrValue {

			t.Errorf("SortHandsByStrength() = %v, want %v", got, want)
		}
	}

}
