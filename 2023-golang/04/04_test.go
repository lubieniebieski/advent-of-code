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
		`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
		Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
		Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
		Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
		Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`,
		13,
	},
	{
		`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
		Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
		Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
		Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
		Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`,
		13,
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
	got := PartTwo(testCases[2].input)
	toolstest.CompareWithExample(t, want, got)
}

// Main tests below

func TestAddCard(t *testing.T) {
	want := Card{
		Numbers:        []int{41, 48, 83, 86, 17},
		WinningNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53},
	}

	got := AddCard("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")

	t.Run("Numbers are correct", func(t *testing.T) {

		if reflect.DeepEqual(want.Numbers, got.Numbers) == false {
			t.Errorf("AddCard() = %v, want %v", got.Numbers, want.Numbers)
		}
	})
	t.Run("WinningNumbers are correct", func(t *testing.T) {
		if reflect.DeepEqual(want.WinningNumbers, got.WinningNumbers) == false {
			t.Errorf("AddCard() = %v, want %v", got.WinningNumbers, want.WinningNumbers)
		}
	})
}

func TestPoints(t *testing.T) {
	testCases := []struct {
		input Card
		want  int
	}{
		{
			Card{
				Numbers:        []int{41, 48},
				WinningNumbers: []int{83, 86},
			},
			0,
		},
		{
			Card{
				Numbers:        []int{41, 48},
				WinningNumbers: []int{83, 41},
			},
			1,
		},
		{
			Card{
				Numbers:        []int{41, 48},
				WinningNumbers: []int{41, 48},
			},
			2,
		},
		{
			Card{
				Numbers:        []int{41, 48, 47},
				WinningNumbers: []int{41, 48, 47},
			},
			4,
		},
		{
			Card{
				Numbers:        []int{41, 48, 47, 5},
				WinningNumbers: []int{41, 48, 47, 5},
			},
			8,
		},
	}

	for _, tc := range testCases {
		t.Run("Points", func(t *testing.T) {
			if tc.input.Points() != tc.want {
				t.Errorf("Points() = %v, want %v", tc.input.Points(), tc.want)
			}
		})
	}
}
