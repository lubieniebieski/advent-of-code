package main

import (
	"testing"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
	"github.com/lubieniebieski/advent-of-code/2022-golang/tools/toolstest"
)

func TestPartOne(t *testing.T) {
	want := 15
	got := PartOne(testData())
	toolstest.CompareWithExample(t, want, got)
}

func TestPartTwo(t *testing.T) {
	want := 12
	got := PartTwo(testData())
	toolstest.CompareWithExample(t, want, got)
}

func TestRoundScore(t *testing.T) {
	t.Run("for a case of 'A Y'", func(t *testing.T) {
		want := 8
		got := RoundScore("A Y")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'A X'", func(t *testing.T) {
		want := 4
		got := RoundScore("A X")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'A Z'", func(t *testing.T) {
		want := 3
		got := RoundScore("A Z")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'B X'", func(t *testing.T) {
		want := 1
		got := RoundScore("B X")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'B Y'", func(t *testing.T) {
		want := 5
		got := RoundScore("B Y")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'B Z'", func(t *testing.T) {
		want := 9
		got := RoundScore("B Z")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'C Z'", func(t *testing.T) {
		want := 6
		got := RoundScore("C Z")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'C X'", func(t *testing.T) {
		want := 7
		got := RoundScore("C X")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'C Y'", func(t *testing.T) {
		want := 2
		got := RoundScore("C Y")
		toolstest.CompareWithExample(t, want, got)
	})
}

func TestRoundScoreV2(t *testing.T) {
	t.Run("for a case of 'A Y'", func(t *testing.T) {
		want := 4
		got := RoundScoreV2("A Y")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'A X'", func(t *testing.T) {
		want := 3
		got := RoundScoreV2("A X")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'A Z'", func(t *testing.T) {
		want := 8
		got := RoundScoreV2("A Z")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'B X'", func(t *testing.T) {
		want := 1
		got := RoundScoreV2("B X")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'B Y'", func(t *testing.T) {
		want := 5
		got := RoundScoreV2("B Y")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'B Z'", func(t *testing.T) {
		want := 9
		got := RoundScoreV2("B Z")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'C Z'", func(t *testing.T) {
		want := 7
		got := RoundScoreV2("C Z")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'C X'", func(t *testing.T) {
		want := 2
		got := RoundScoreV2("C X")
		toolstest.CompareWithExample(t, want, got)
	})
	t.Run("for a case of 'C Y'", func(t *testing.T) {
		want := 6
		got := RoundScoreV2("C Y")
		toolstest.CompareWithExample(t, want, got)
	})
}

func testData() []string {
	return tools.ExtractStringsFromString(testInput())
}

func testInput() string {
	return `A Y
	B X
	C Z`
}
