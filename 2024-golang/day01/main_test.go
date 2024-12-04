package day01

import "testing"

func TestSolve1(t *testing.T) {
	input := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	want := 11
	got := Solve1(input)
	if got != want {
		t.Errorf("solve1() = %v, want %v", got, want)
	}
}

func TestSolve2(t *testing.T) {
	input := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	want := 31
	got := Solve2(input)
	if got != want {
		t.Errorf("solve2() = %v, want %v", got, want)
	}
}
