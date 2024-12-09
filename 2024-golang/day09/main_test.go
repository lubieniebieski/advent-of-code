package day09

import (
	"strings"
	"testing"
)

// Helper function to convert readable string to slice
func stringToSlice(s string) []string {
	return strings.Split(s, "")
}

func TestConvertLayoutToBlocks(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"12345", stringToSlice("0..111....22222")},
		{"2333133121414131402", stringToSlice("00...111...2...333.44.5555.6666.777.888899")},
	}

	for _, c := range cases {
		got := ConvertLayoutToBlocks(c.in)
		if strings.Join(got, "") != strings.Join(c.want, "") {
			t.Errorf("ConvertLayoutToBlocks(%q) == %q, want %q", c.in, strings.Join(got, ""), strings.Join(c.want, ""))
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
		got := SortLayout(stringToSlice(c.in))
		if strings.Join(got, "") != c.want {
			t.Errorf("SortLayout(%q) == %q, want %q", c.in, strings.Join(got, ""), c.want)
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
		got := CalculateChecksum(stringToSlice(c.in))
		if got != c.want {
			t.Errorf("CalculateChecksum(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestSolve1(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"2333133121414131402", 1928},
	}

	for _, c := range cases {
		got := Solve1([]string{c.in})
		if got != c.want {
			t.Errorf("Solve1(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

func TestSolve2(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"2333133121414131402", 2858}, // Update with actual test case
	}

	for _, c := range cases {
		got := Solve2([]string{c.in})
		if got != c.want {
			t.Errorf("Solve2(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
