package day08

import "testing"

func TestSolve1(t *testing.T) {
	input := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}
	want := 14 // Expected result
	got := Solve1(input)
	if got != want {
		t.Errorf("solve1() = %v, want %v", got, want)
	}
}

func TestSolve2(t *testing.T) {
	input := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}

	want := 34 // Expected result
	got := Solve2(input)
	if got != want {
		t.Errorf("solve2() = %v, want %v", got, want)
	}
}
