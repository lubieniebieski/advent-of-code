package day04

import "testing"

func TestSolve1(t *testing.T) {
	input := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	want := 18 // Expected result
	got := Solve1(input)
	if got != want {
		t.Errorf("solve1() = %v, want %v", got, want)
	}
}

func TestSolve2(t *testing.T) {
	input := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	want := 9 // Expected result
	got := Solve2(input)
	if got != want {
		t.Errorf("solve2() = %v, want %v", got, want)
	}
}
