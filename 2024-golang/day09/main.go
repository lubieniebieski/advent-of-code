package day09 // Change this for each new day

import (
	"slices"
	"strconv"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

func ConvertLayoutToBlocks(input string) []string {
	var newSlice []string
	id := 0
	for i := 0; i < len(input); i += 2 {
		blockLength, _ := utils.GetIntFromString(string(input[i]))
		freeSpace := 0
		if i < len(input)-1 {
			freeSpace, _ = utils.GetIntFromString(string(input[i+1]))
		}
		strId := strconv.Itoa(id)
		for j := 0; j < blockLength; j++ {
			newSlice = append(newSlice, strId)
		}
		id++
		for j := 0; j < freeSpace; j++ {
			newSlice = append(newSlice, ".")
		}
	}
	return newSlice
}

func CalculateChecksum(input []string) int {
	checksum := 0
	freeSpace := strings.Count(strings.Join(input, ""), ".")
	freeSpaceIndex := len(input) - freeSpace
	for i := 0; i < freeSpaceIndex; i++ {
		integer, _ := utils.GetIntFromString(input[i])
		checksum += i * integer
	}
	return checksum
}

func SortLayout(input []string) []string {
	freeSpace := strings.Count(strings.Join(input, ""), ".")
	freeSpaceIndex := len(input) - freeSpace

	for i := len(input) - 1; i >= freeSpaceIndex; i-- {
		if input[i] != "." {
			j := slices.Index(input, ".")
			input[i], input[j] = input[j], input[i]
		}
	}
	return input
}

func Solve1(input []string) int {
	layout := ConvertLayoutToBlocks(input[0])
	sortedLayout := SortLayout(layout)
	return CalculateChecksum(sortedLayout)
}

func Solve2(input []string) int {
	return 0
}
