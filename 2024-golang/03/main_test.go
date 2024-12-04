package main

import "testing"

func TestSolve1(t *testing.T) {
	input := []string{
		"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
	}
	want := 161
	got := solve1(input)
	if got != want {
		t.Errorf("solve1() = %v, want %v", got, want)
	}
}

func TestSolve2(t *testing.T) {
	input := []string{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	}
	want := 48
	got := solve2(input)
	if got != want {
		t.Errorf("solve2() = %v, want %v", got, want)
	}
}
