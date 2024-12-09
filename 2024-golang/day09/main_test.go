package day09

import "testing"

func TestConvertLayoutToBlocks(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"12345", "0..111....22222"},
		{"2333133121414131402", "00...111...2...333.44.5555.6666.777.888899"},
	}

	for _, c := range cases {
		got := ConvertLayoutToBlocks(c.in)
		if got != c.want {
			t.Errorf("ConvertLayoutToBlocks(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSortLayout(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"0..111....22222", "022111222......"},
		{"00...111...2...333.44.5555.6666.777.888899", "0099811188827773336446555566.............."},
	}

	for _, c := range cases {
		got := SortLayout(c.in)
		if got != c.want {
			t.Errorf("SortLayout(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCalculateChecksum(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"0099811188827773336446555566..............", 1928},
	}

	for _, c := range cases {
		got := CalculateChecksum(c.in)
		if got != c.want {
			t.Errorf("CalculateChecksum(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSolve1(t *testing.T) {
	input := []string{
		"2333133121414131402",
	}
	want := 1928 // Expected result
	got := Solve1(input)
	if got != want {
		t.Errorf("solve1() = %v, want %v", got, want)
	}
}

func TestSolve2(t *testing.T) {
	input := []string{
		// Add your test input here
	}
	want := 0 // Expected result
	got := Solve2(input)
	if got != want {
		t.Errorf("solve2() = %v, want %v", got, want)
	}
}
