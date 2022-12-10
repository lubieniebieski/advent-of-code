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

func CreateBoolMatrix(size int) [][]bool {
	matrix := make([][]bool, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]bool, size)
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

func PartOne(data []string) (result int) {
	max := len(data)
	maxes := CreateBoolMatrix(max)
	matrix := ToIntMatrix(data)
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
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			if maxes[i][j] {
				result += 1
			}
		}
	}
	return result
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
