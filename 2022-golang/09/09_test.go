package main

import (
	"testing"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
	"github.com/lubieniebieski/advent-of-code/2022-golang/tools/toolstest"
)

func TestMakeMove(t *testing.T) {
	t.Run("Right direction", func(t *testing.T) {
		matrix := tools.CreateBoolMatrix(5)
		header := Point{x: 0, y: 4}
		tail := Point{x: 0, y: 4}
		step := Step{"R", 4}
		MakeMove(&matrix, &header, &tail, step)

		t.Run("head&tail positions", func(t *testing.T) {
			assertPosition(t, header, Point{x: 4, y: 4})
			assertPosition(t, tail, Point{x: 3, y: 4})
		})

		t.Run("correct true values", func(t *testing.T) {
			expectedTrues := []Point{
				{x: 0, y: 4},
				{x: 1, y: 4},
				{x: 2, y: 4},
				{x: 3, y: 4},
			}
			assertTrueCoordinates(t, matrix, expectedTrues)
		})

		t.Run("Marks the right number of true fields", func(t *testing.T) {
			assertCorrectNumberOfTrue(t, matrix, 4)
		})
	})

	t.Run("Left direction", func(t *testing.T) {
		matrix := tools.CreateBoolMatrix(5)
		header := Point{x: 3, y: 0}
		tail := Point{x: 3, y: 0}
		step := Step{"L", 3}
		MakeMove(&matrix, &header, &tail, step)

		t.Run("head&tail positions", func(t *testing.T) {
			assertPosition(t, header, Point{x: 0, y: 0})
			assertPosition(t, tail, Point{x: 1, y: 0})
		})

		t.Run("correct true values", func(t *testing.T) {
			expectedTrues := []Point{
				{x: 3, y: 0},
				{x: 2, y: 0},
				{x: 1, y: 0},
			}
			assertTrueCoordinates(t, matrix, expectedTrues)
		})

		t.Run("Marks the right number of true fields", func(t *testing.T) {
			assertCorrectNumberOfTrue(t, matrix, 3)
		})
	})

	t.Run("Up direction", func(t *testing.T) {
		matrix := tools.CreateBoolMatrix(5)
		header := Point{x: 0, y: 4}
		tail := Point{x: 0, y: 4}
		step := Step{"U", 4}
		MakeMove(&matrix, &header, &tail, step)
		t.Run("head&tail positions", func(t *testing.T) {
			assertPosition(t, header, Point{x: 0, y: 0})
			assertPosition(t, tail, Point{x: 0, y: 1})
		})

		t.Run("correct true values", func(t *testing.T) {
			expectedTrues := []Point{
				{x: 0, y: 1},
				{x: 0, y: 2},
				{x: 0, y: 3},
				{x: 0, y: 4},
			}
			assertTrueCoordinates(t, matrix, expectedTrues)
		})

		t.Run("Marks the right number of true fields", func(t *testing.T) {
			assertCorrectNumberOfTrue(t, matrix, 4)
		})
	})
	t.Run("Down direction", func(t *testing.T) {
		matrix := tools.CreateBoolMatrix(5)
		header := Point{x: 0, y: 3}
		tail := Point{x: 0, y: 2}
		step := Step{"D", 1}
		MakeMove(&matrix, &header, &tail, step)

		t.Run("head&tail positions", func(t *testing.T) {
			assertPosition(t, header, Point{x: 0, y: 4})
			assertPosition(t, tail, Point{x: 0, y: 3})
		})

		t.Run("correct true values", func(t *testing.T) {
			expectedTrues := []Point{
				{x: 0, y: 2},
				{x: 0, y: 3},
			}
			assertTrueCoordinates(t, matrix, expectedTrues)
		})

		t.Run("Marks the right number of true fields", func(t *testing.T) {
			assertCorrectNumberOfTrue(t, matrix, 2)
		})
	})
	t.Run("Diagonal case #1", func(t *testing.T) {
		matrix := tools.CreateBoolMatrix(5)
		header := Point{x: 2, y: 2}
		tail := Point{x: 1, y: 3}
		step := Step{"U", 1}
		MakeMove(&matrix, &header, &tail, step)
		t.Run("head&tail positions", func(t *testing.T) {
			assertPosition(t, header, Point{x: 2, y: 1})
			assertPosition(t, tail, Point{x: 2, y: 2})
		})

		t.Run("correct true values", func(t *testing.T) {
			expectedTrues := []Point{
				{x: 1, y: 3},
				{x: 2, y: 2},
			}
			PrintGrid(&matrix)
			assertTrueCoordinates(t, matrix, expectedTrues)
		})

		t.Run("Marks the right number of true fields", func(t *testing.T) {
			assertCorrectNumberOfTrue(t, matrix, 2)
		})
	})

	t.Run("Diagonal case #2", func(t *testing.T) {
		matrix := tools.CreateBoolMatrix(5)
		header := Point{x: 2, y: 2}
		tail := Point{x: 1, y: 3}
		step := Step{"R", 1}
		MakeMove(&matrix, &header, &tail, step)

		t.Run("head&tail positions", func(t *testing.T) {
			assertPosition(t, header, Point{x: 3, y: 2})
			assertPosition(t, tail, Point{x: 2, y: 2})
		})

		t.Run("correct true values", func(t *testing.T) {
			expectedTrues := []Point{
				{x: 1, y: 3},
				{x: 2, y: 2},
			}
			assertTrueCoordinates(t, matrix, expectedTrues)
		})

		t.Run("Marks the right number of true fields", func(t *testing.T) {
			assertCorrectNumberOfTrue(t, matrix, 2)
		})
	})

	t.Run("Diagonal case #3", func(t *testing.T) {
		matrix := tools.CreateBoolMatrix(5)
		header := Point{x: 1, y: 0}
		tail := Point{x: 2, y: 0}
		step := Step{"D", 1}
		MakeMove(&matrix, &header, &tail, step)

		t.Run("head&tail positions", func(t *testing.T) {
			assertPosition(t, header, Point{x: 1, y: 1})
			assertPosition(t, tail, Point{x: 2, y: 0})
		})

		t.Run("Marks the right number of true fields", func(t *testing.T) {
			assertCorrectNumberOfTrue(t, matrix, 1)
		})
	})

	t.Run("Diagonal case #4", func(t *testing.T) {
		matrix := tools.CreateBoolMatrix(6)
		header := Point{x: 4, y: 4}
		tail := Point{x: 3, y: 4}
		step := Step{"U", 4}
		MakeMove(&matrix, &header, &tail, step)

		t.Run("head&tail positions", func(t *testing.T) {
			assertPosition(t, header, Point{x: 4, y: 0})
			assertPosition(t, tail, Point{x: 4, y: 1})
		})

		t.Run("correct true values", func(t *testing.T) {
			expectedTrues := []Point{
				{x: 4, y: 1},
				{x: 4, y: 2},
				{x: 4, y: 3},
				{x: 3, y: 4},
			}
			assertTrueCoordinates(t, matrix, expectedTrues)
		})

		t.Run("Marks the right number of true fields", func(t *testing.T) {
			assertCorrectNumberOfTrue(t, matrix, 4)
		})
	})

	t.Run("Diagonal case #5", func(t *testing.T) {
		matrix := tools.CreateBoolMatrix(6)
		header := Point{x: 4, y: 0}
		tail := Point{x: 4, y: 1}
		step := Step{"L", 3}
		MakeMove(&matrix, &header, &tail, step)

		t.Run("head&tail positions", func(t *testing.T) {
			assertPosition(t, header, Point{x: 1, y: 0})
			assertPosition(t, tail, Point{x: 2, y: 0})
		})

		t.Run("correct true values", func(t *testing.T) {
			expectedTrues := []Point{
				{x: 2, y: 0},
				{x: 3, y: 0},
			}
			assertTrueCoordinates(t, matrix, expectedTrues)
		})

		t.Run("Marks the right number of true fields", func(t *testing.T) {
			assertCorrectNumberOfTrue(t, matrix, 3)
		})
	})

	t.Run("Diagonal case #6", func(t *testing.T) {
		matrix := tools.CreateBoolMatrix(6)
		header := Point{x: 1, y: 1}
		tail := Point{x: 2, y: 0}
		step := Step{"R", 4}
		MakeMove(&matrix, &header, &tail, step)

		t.Run("head&tail positions", func(t *testing.T) {
			assertPosition(t, header, Point{x: 5, y: 1})
			assertPosition(t, tail, Point{x: 4, y: 1})
		})

		t.Run("correct true values", func(t *testing.T) {
			expectedTrues := []Point{
				{x: 2, y: 0},
				{x: 3, y: 1},
				{x: 4, y: 1},
			}
			assertTrueCoordinates(t, matrix, expectedTrues)
		})

		t.Run("Marks the right number of true fields", func(t *testing.T) {
			assertCorrectNumberOfTrue(t, matrix, 3)
		})
	})
}

func TestPartOne(t *testing.T) {
	want := 13
	got := PartOne(testData())
	toolstest.CompareWithExample(t, want, got)
}

func assertPosition(t testing.TB, got Point, want Point) {
	t.Helper()
	if got.x != want.x || got.y != want.y {
		t.Errorf("Wrong position: expected %v, got %v", want, got)
	}
}

func assertCorrectNumberOfTrue(t testing.TB, matrix [][]bool, expectedResult int) {
	t.Helper()
	if got := CountTrue(&matrix); got != expectedResult {
		t.Errorf("Expected matrix to add %v *true, got %v", expectedResult, got)
	}
}

func assertTrueCoordinates(t testing.TB, matrix [][]bool, expectedTrues []Point) {
	t.Helper()
	for _, test := range expectedTrues {
		if !matrix[test.y][test.x] {
			t.Errorf("Expected %v coordinates to be marked true", test)
		}
	}
}

func testData() []string {
	return tools.ExtractStringsFromString(testInput())
}

func testInput() string {
	return `R 4
	U 4
	L 3
	D 1
	R 4
	D 1
	L 5
	R 2`
}
