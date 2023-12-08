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
		`RL

		AAA = (BBB, CCC)
		BBB = (DDD, EEE)
		CCC = (ZZZ, GGG)
		DDD = (DDD, DDD)
		EEE = (EEE, EEE)
		GGG = (GGG, GGG)
		ZZZ = (ZZZ, ZZZ`,
		2,
	},
	{
		``,
		5905,
	},
	{
		`LLR

		AAA = (BBB, BBB)
		BBB = (AAA, ZZZ)
		ZZZ = (ZZZ, ZZZ)`,
		6,
	},
}

func TestPartOne(t *testing.T) {
	t.Run("Original input", func(t *testing.T) {
		want := testCases[1].want
		got := PartOne(testCases[1].input)
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("Original input #2", func(t *testing.T) {
		want := testCases[3].want
		got := PartOne(testCases[3].input)
		fmt.Println(got, want)
		// toolstest.CompareWithExample(t, want, got)
	})

}

func TestPartTwo(t *testing.T) {
	t.SkipNow()
	t.Run("Original input", func(t *testing.T) {
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

func TestNodeFromString(t *testing.T) {
	testCases := []struct {
		input string
		want  Node
	}{
		{
			"AAA = (BBB, CCC)",
			Node{
				ID: "AAA",
				Left: &Node{
					ID: "BBB",
				},
				Right: &Node{
					ID: "CCC",
				},
			},
		},
	}

	for _, tc := range testCases {
		got := NodeFromString(tc.input)
		if got.ID != tc.want.ID {
			t.Errorf("ID is wrong, got: %s, want: %s.", got.ID, tc.want.ID)
		}
		if got.Left.ID != tc.want.Left.ID {
			t.Errorf("Left is wrong, got: %s, want: %s.", got.Left.ID, tc.want.Left.ID)
		}
		if got.Right.ID != tc.want.Right.ID {
			t.Errorf("Right is wrong, got: %s, want: %s.", got.Right.ID, tc.want.Right.ID)
		}
	}
}
