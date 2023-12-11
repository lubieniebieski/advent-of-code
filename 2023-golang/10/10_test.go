package main

import (
	"reflect"
	"slices"
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
		`.....
		.S-7.
		.|.|.
		.L-J.
		.....`,
		4,
	},
	{
		`..F7.
		.FJ|.
		SJ.L7
		|F--J
		LJ...`,
		8,
	},
	{
		`...........
		.S-------7.
		.|F-----7|.
		.||.....||.
		.||.....||.
		.|L-7.F-J|.
		.|..|.|..|.
		.L--J.L--J.
		...........`,
		4,
	},
}

func TestPartOne(t *testing.T) {
	t.Run("Works for provided example", func(t *testing.T) {
		want := testCases[1].want
		got := PartOne(testCases[1].input)
		toolstest.CompareWithExample(t, want, got)
	})

	t.Run("Works for provided example #2", func(t *testing.T) {
		want := testCases[2].want
		got := PartOne(testCases[2].input)
		toolstest.CompareWithExample(t, want, got)
	})
}

func TestPartTwo(t *testing.T) {
	t.Run("Works for provided example", func(t *testing.T) {
		want := testCases[3].want
		got := PartTwo(testCases[3].input)
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

func TestParsePipePoints(t *testing.T) {
	input := `...
	|S|`
	want := [][]PipePoint{
		{
			{0, 0, "."},
			{1, 0, "."},
			{2, 0, "."},
		},
		{
			{0, 1, "|"},
			{1, 1, "S"},
			{2, 1, "|"},
		},
	}
	got := ParsePipePoints(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParsePipePoints() = %v, want %v", got, want)
	}
}

func TestPipePoint_IsStart(t *testing.T) {
	tests := []struct {
		name string
		p    PipePoint
		want bool
	}{
		{
			"Returns true if id is S",
			PipePoint{0, 0, "S"},
			true,
		},
		{
			"Returns false if id is not S",
			PipePoint{0, 0, "L"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsStart(); got != tt.want {
				t.Errorf("PipePoint.IsStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPipePoint_WhereDoIStart(t *testing.T) {
	t.Run("Simple example", func(t *testing.T) {
		// F7
		// LS
		pipeMap := [][]PipePoint{
			{
				{0, 0, "F"},
				{1, 0, "7"},
			},
			{
				{0, 1, "L"},
				{1, 1, "S"},
			},
		}
		startPoint := pipeMap[1][1]
		want := []PipePoint{
			pipeMap[0][1],
			pipeMap[1][0],
		}
		got := startPoint.WhereDoIStart(pipeMap)
		if len(got) != len(want) {
			t.Errorf("Length is wrong, PipePoint.WhereDoIStart() = %v, want %v", got, want)
		}
		for i := range got {
			if !slices.Contains(want, got[i]) {
				t.Errorf("PipePoint.WhereDoIStart() = %v, want %v", got, want)
			}
		}
	})
	t.Run("more complicated example", func(t *testing.T) {
		pipePoints := ParsePipePoints(testCases[2].input)
		start := pipePoints[2][0]
		want := []PipePoint{
			pipePoints[2][1],
			pipePoints[3][0],
		}
		got := start.WhereDoIStart(pipePoints)
		if len(got) != len(want) {
			t.Errorf("Length is wrong, PipePoint.WhereDoIStart() = %v, want %v", got, want)
		}
		for i := range got {
			if !slices.Contains(want, got[i]) {
				t.Errorf("PipePoint.WhereDoIStart() = %v, want %v", got, want)
			}
		}
	})
}

func TestPipePoint_MoveFromPoint(t *testing.T) {
	pipeMap := [][]PipePoint{
		{
			{0, 0, "F"},
			{1, 0, "J"},
		},
		{
			{0, 1, "L"},
			{1, 1, "S"},
		},
	}
	point := pipeMap[1][0]
	want := pipeMap[0][0]
	previousPoint := pipeMap[1][1]
	_, got := point.MoveFromPoint(previousPoint, pipeMap)
	if got != want {
		t.Errorf("PipePoint.MoveFromPoint() = %v, want %v", got, want)
	}
}
