package day10

import "testing"

func TestSolve1(t *testing.T) {
	testCases := []struct {
		input []string
		want  int
	}{
		{
			input: []string{
				"..90..9",
				"...1.98",
				"...2..7",
				"6543456",
				"765.987",
				"876....",
				"987....",
			},
			want: 4,
		},
		{
			input: []string{
				"...0...",
				"...1...",
				"...2...",
				"6543456",
				"7.....7",
				"8.....8",
				"9.....9",
			},
			want: 2,
		},
		{
			input: []string{
				"10..9..",
				"2...8..",
				"3...7..",
				"4567654",
				"...8..3",
				"...9..2",
				".....01",
			},
			want: 3,
		},

		{
			input: []string{
				"89010123",
				"78121874",
				"87430965",
				"96549874",
				"45678903",
				"32019012",
				"01329801",
				"10456732",
			},
			want: 36,
		},
	}
	for _, tt := range testCases {
		got := Solve1(tt.input)
		if got != tt.want {
			t.Errorf("Solve1() = %v, want %v", got, tt.want)
		}
	}

}

func TestSolve2(t *testing.T) {
	testCases := []struct {
		input []string
		want  int
	}{
		{
			input: []string{
				"89010123",
				"78121874",
				"87430965",
				"96549874",
				"45678903",
				"32019012",
				"01329801",
				"10456732",
			},
			want: 81,
		},
	}
	for _, tt := range testCases {
		got := Solve2(tt.input)
		if got != tt.want {
			t.Errorf("Solve2() = %v, want %v", got, tt.want)
		}
	}
}
