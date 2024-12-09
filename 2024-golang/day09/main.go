package day09 // Change this for each new day

import (
	"strconv"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

func ConvertLayoutToBlocks(input string) string {
	newString := ""
	id := 0
	for i := 0; i < len(input); i += 2 {
		blockLength, _ := utils.GetIntFromString(string(input[i]))
		freeSpace := 0
		if i < len(input)-1 {
			freeSpace, _ = utils.GetIntFromString(string(input[i+1]))
		}
		strId := strconv.Itoa(id)
		newString = newString + strings.Repeat(strId, blockLength)
		id++
		newString = newString + strings.Repeat(".", freeSpace)
	}
	return newString
}

func CalculateChecksum(input string) int {
	checksum := 0
	freeSpace := strings.Count(input, ".")
	freeSpaceIndex := len(input) - freeSpace
	for i := 0; i < freeSpaceIndex; i++ {
		integer, _ := utils.GetIntFromString(string(input[i]))
		checksum += i * integer
	}
	return checksum
}

func SortLayout(input string) string {
	freeSpace := strings.Count(input, ".")
	freeSpaceIndex := len(input) - freeSpace
	for i := len(input) - 1; i >= freeSpaceIndex; i-- {
		if input[i] != '.' {
			ind := strings.Index(input, ".")
			newString := input[:ind] + string(input[i]) + input[ind+1:i] + "." + input[i+1:]
			return SortLayout(newString)
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
	// TODO: Implement solution for part 2
	return 0
}
