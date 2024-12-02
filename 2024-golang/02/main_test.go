package main

import "testing"

func TestSolve1(t *testing.T) {
	input := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
	want := 2
	got := solve1(input)
	if got != want {
		t.Errorf("solve1() = %v, want %v", got, want)
	}
}

func TestSolve2(t *testing.T) {
	input := []string{
		// Add your test input here
	}
	want := 0 // Expected result
	got := solve2(input)
	if got != want {
		t.Errorf("solve2() = %v, want %v", got, want)
	}
}

func TestSafeRow(t *testing.T) {
	tests := []struct {
		name string
		row  string
		want bool
	}{
		{
			name: "safe row - increasing",
			row:  "1 2 3 4 5",
			want: true,
		},
		{
			name: "safe row - decreasing",
			row:  "7 6 4 2 1",
			want: true,
		},
		{
			name: "safe row - all increasing 1-3",
			row:  "1 3 6 7 9",
			want: true,
		},
		{
			name: "unsafe row - duplicate",
			row:  "1 2 3 2 5",
			want: false,
		},
		{
			name: "unsafe row - duplicate at end",
			row:  "1 2 3 4 2",
			want: false,
		},
		{
			name: "unsafe row - too big increase",
			row:  "1 2 7 8 9",
			want: false,
		},
		{
			name: "unsafe row - too big decrease",
			row:  "9 7 6 2 1",
			want: false,
		},
		{
			name: "unsafe row - direction change",
			row:  "1 3 2 4 5",
			want: false,
		},
		{
			name: "unsafe row - plateau",
			row:  "8 6 4 4 1",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := safeRow(tt.row)
			if got != tt.want {
				t.Errorf("safeRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
