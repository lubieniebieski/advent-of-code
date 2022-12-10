package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
)

func NthCharToInt(n int, str string) int {
	value, _ := strconv.Atoi(string(str[n]))
	return value
}

func ToIntMatrix(input []string) [][]int {
	max := len(input)
	matrix := CreateIntMatrix(max)
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			matrix[i][j] = NthCharToInt(j, input[i])
		}
	}
	return matrix
}

func CreateIntMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}
	return matrix
}

func RotateIntMatrix(arr [][]int) [][]int {
	n := len(arr)
	rotated := make([][]int, n)
	for i := range rotated {
		rotated[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[i][j] = arr[n-j-1][i]
		}
	}
	return rotated
}

func RotateBoolMatrix(arr [][]bool) [][]bool {
	n := len(arr)
	rotated := make([][]bool, n)
	for i := range rotated {
		rotated[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[i][j] = arr[n-j-1][i]
		}
	}

	return rotated
}

func CalculateTreeScore(y, x int, matrix [][]int) int {
	max := len(matrix)
	right := 0
	value := matrix[y][x]
	for j := x + 1; j < max; j++ {
		right += 1
		if matrix[y][j] >= value {
			break
		}
	}
	if right == 0 {
		return 0
	}

	left := 0
	for j := x - 1; j >= 0; j-- {
		left += 1
		if matrix[y][j] >= value {
			break
		}
	}
	if left == 0 {
		return 0
	}
	down := 0
	for i := y + 1; i < max; i++ {
		down += 1
		if matrix[i][x] >= value {
			break
		}
	}
	if down == 0 {
		return 0
	}
	up := 0
	for i := y - 1; i >= 0; i-- {
		up += 1
		if matrix[i][x] >= value {
			break
		}

	}
	return right * left * up * down
}

func PrepareDataForResult(data *[]string) (matrix [][]int, maxes [][]bool) {
	max := len(*data)
	maxes = tools.CreateBoolMatrix(max)
	matrix = ToIntMatrix(*data)
	var rotations = 1
	for rotations <= 4 {
		for i := 0; i < max; i++ {
			maxes[0][i] = true
		}
		for i := 1; i < max; i++ {

			var maxHeight = -1
			for j := 0; j < max; j++ {
				if matrix[i][j] > maxHeight {
					maxHeight = matrix[i][j]
					maxes[i][j] = true
				}
			}
		}

		matrix = RotateIntMatrix(matrix)
		maxes = RotateBoolMatrix(maxes)
		rotations += 1
	}
	return matrix, maxes
}

func PartOne(data []string) (result int) {
	_, maxes := PrepareDataForResult(&data)

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			if maxes[i][j] {
				result += 1
			}
		}
	}
	return result
}

func PartTwo(data []string) (result int) {

	matrix, maxes := PrepareDataForResult(&data)

	scores := []int{}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			if maxes[i][j] {
				scores = append(scores, CalculateTreeScore(i, j, matrix))
			}
		}
	}
	_, max := tools.FindMinAndMax(scores)
	return max
}

func parsedData() []string {
	rawData, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return tools.ExtractStringsFromString(string(rawData))
}

func main() {
	data := parsedData()

	fmt.Printf("Part One: %v \n", PartOne(data))
	fmt.Printf("Part Two: %v \n", PartTwo(data))
}
