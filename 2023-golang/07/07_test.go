package main

import (
	"fmt"
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
		5905,
	},
	{
		`2345A 1
		Q2KJJ 13
		Q2Q2Q 19
		T3T3J 17
		T3Q33 11
		2345J 3
		J345A 2
		32T3K 5
		T55J5 29
		KK677 7
		KTJJT 34
		QQQJA 31
		JJJJJ 37
		JAAAA 43
		AAAAJ 59
		AAAAA 61
		2AAAA 23
		2JJJJ 53
		JJJJ2 41`,
		6592,
	},
	{
		`AAAAA 2
		22222 3
		AAAAK 5
		22223 7
		AAAKK 11
		22233 13
		AAAKQ 17
		22234 19
		AAKKQ 23
		22334 29
		AAKQJ 31
		22345 37
		AKQJT 41
		23456 43`,
		1343,
	},
}

func TestPartOne(t *testing.T) {
	t.Run("Original input", func(t *testing.T) {
		want := testCases[1].want
		got := PartOne(testCases[1].input)
		toolstest.CompareWithExample(t, want, got)
	})

	t.Run("More test cases #1", func(t *testing.T) {
		want := testCases[3].want
		got := PartOne(testCases[3].input)
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("More test cases #2", func(t *testing.T) {
		want := testCases[4].want
		got := PartOne(testCases[4].input)
		toolstest.CompareWithExample(t, want, got)
	})

	t.Run("More test cases #3", func(t *testing.T) {
		want := 250347426
		got := PartOne(parsedData())
		toolstest.CompareWithExample(t, want, got)
	})
}

func TestPartTwo(t *testing.T) {
	t.Run("Original input", func(t *testing.T) {
		want := testCases[2].want
		got := PartTwo(testCases[1].input)
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("More test cases #1", func(t *testing.T) {
		want := 6839
		got := PartTwo(testCases[3].input)
		toolstest.CompareWithExample(t, want, got)
	})

	t.Run("More test cases #3", func(t *testing.T) {
		want := 251224870
		got := PartTwo(parsedData())
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
		{"KKKAA", FullHouse},
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
func TestHand_StrengthV2(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"32T3K", OnePair},
		{"KK677", TwoPairs},
		{"T55J5", FourOfAKind},
		{"JJ234", ThreeOfAKind},
		{"JK234", OnePair},
		{"JJ222", FiveOfAKind},
		{"KTJJT", FourOfAKind},
		{"KKJJT", FourOfAKind},
		{"KKJJK", FiveOfAKind},
		{"QKJJT", ThreeOfAKind},
		{"QQQJA", FourOfAKind},
		{"JJJJJ", FiveOfAKind},
		{"JJJAJ", FiveOfAKind},
		{"JJAAJ", FiveOfAKind},
		{"JAAAJ", FiveOfAKind},
		{"AAAAJ", FiveOfAKind},
		{"JAAAA", FiveOfAKind},
		{"AAAAA", FiveOfAKind},
		{"AATTJ", FullHouse},
		{"23332", FullHouse},
		{"234JJ", ThreeOfAKind},
		{"2345J", OnePair},
		{"23444", ThreeOfAKind},
		{"23422", ThreeOfAKind},
		{"23433", ThreeOfAKind},
		{"AAA33", FullHouse},
		{"A3AA4", ThreeOfAKind},
		{"33AA3", FullHouse},
		{"13AA3", TwoPairs},
		{"12345", HighCard},
		{"1234J", OnePair},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := Hand{StrValue: tc.input}.StrengthV2()
			toolstest.CompareWithExample(t, tc.want, got)
		})
	}
}

func TestSortHandsByStrength(t *testing.T) {
	input := HandsFromString(testCases[1].input)
	want := []Hand{
		{StrValue: "32T3K"},
		{StrValue: "KTJJT"},
		{StrValue: "KK677"},
		{StrValue: "T55J5"},
		{StrValue: "QQQJA"},
	}

	got := SortHandsByStrengthAsc(input)
	for i, hand := range got {
		if hand.StrValue != want[i].StrValue {
			t.Errorf("SortHandsByStrength() = %v, want %v", got, want)
		}
	}

	t.Run("More test cases", func(t *testing.T) {
		input := []Hand{
			{StrValue: "52346"},
			{StrValue: "23456"},
			{StrValue: "54321"},
		}
		want := []Hand{
			{StrValue: "23456"},
			{StrValue: "52346"},
			{StrValue: "54321"},
		}
		got := SortHandsByStrengthAsc(input)
		for i, hand := range got {
			if hand.StrValue != want[i].StrValue {
				t.Errorf("SortHandsByStrength() = %v, want %v", got, want)
			}
		}
	})
}

func TestSortHandsByStrengthV2(t *testing.T) {
	t.Run("More test cases", func(t *testing.T) {
		input := []Hand{
			{StrValue: "52346"},
			{StrValue: "23456"},
			{StrValue: "54321"},
		}
		want := []Hand{
			{StrValue: "23456"},
			{StrValue: "52346"},
			{StrValue: "54321"},
		}
		got := SortHandsByStrengthAsc(input)
		for i, hand := range got {
			if hand.StrValue != want[i].StrValue {
				t.Errorf("SortHandsByStrength() = %v, want %v", got, want)
			}
		}
	})
}

func TestCompareEqualTypesOfHands(t *testing.T) {

	testCases := []struct {
		lower  Hand
		higher Hand
	}{
		{Hand{StrValue: "T55J5"}, Hand{StrValue: "QQQJA"}},
		{Hand{StrValue: "QJJQ2"}, Hand{StrValue: "QQQQ2"}},
		{Hand{StrValue: "2JJJJ"}, Hand{StrValue: "JJJJ2"}},
		{Hand{StrValue: "AAAAJ"}, Hand{StrValue: "AAAAA"}},
		{Hand{StrValue: "JJJJJ"}, Hand{StrValue: "QQQQQ"}},
	}

	for _, tc := range testCases {
		name := fmt.Sprint(tc.lower.StrValue + " is lower than " + tc.higher.StrValue)
		t.Run(name, func(t *testing.T) {
			if IsFirstHandHigher(tc.lower, tc.higher) {
				t.Errorf("%v is higher than %v, should be the other way around", tc.lower.StrValue, tc.higher.StrValue)
			}
		})
	}
}

func TestCompareEqualTypesOfHandsV2(t *testing.T) {
	testCases := []struct {
		lower  Hand
		higher Hand
	}{
		{Hand{StrValue: "11111"}, Hand{StrValue: "1111A"}},
		{Hand{StrValue: "1111Q"}, Hand{StrValue: "1111K"}},
		{Hand{StrValue: "1111K"}, Hand{StrValue: "Q1111"}},
		{Hand{StrValue: "1111J"}, Hand{StrValue: "1111T"}},
		{Hand{StrValue: "JJJJ2"}, Hand{StrValue: "2AAAA"}},
		{Hand{StrValue: "2222T"}, Hand{StrValue: "TT2TT"}},
		{Hand{StrValue: "JJJJ2"}, Hand{StrValue: "2JJJJ"}},
	}

	for _, tc := range testCases {
		name := fmt.Sprint(tc.lower.StrValue + " is lower than " + tc.higher.StrValue)
		t.Run(name, func(t *testing.T) {
			if IsFirstHandHigherV2(tc.lower, tc.higher) {
				t.Errorf("%v is higher than %v, should be the other way around", tc.lower.StrValue, tc.higher.StrValue)
			}
		})
	}
}
