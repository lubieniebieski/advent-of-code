package day09 // Change this for each new day

import (
	"slices"
	"strconv"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

type Block struct {
	id        int
	size      int
	freeSpace int
}

func ConvertLayoutToBlocks(input string) []Block {
	var blocks []Block
	id := 0
	for i := 0; i < len(input); i += 2 {
		blockLength, _ := utils.GetIntFromString(string(input[i]))
		freeSpace := 0
		if i < len(input)-1 {
			freeSpace, _ = utils.GetIntFromString(string(input[i+1]))
		}
		blocks = append(blocks, Block{
			id:        id,
			size:      blockLength,
			freeSpace: freeSpace,
		})
		id++
	}
	return blocks
}

func SortBlocks(blocks []Block) []Block {
	// iterating over ids
	for i := len(blocks) - 1; i >= 0; i-- {
		blockToMoveIdx := slices.IndexFunc(blocks, func(b Block) bool {
			return b.id == i
		})
		for j := 0; j < blockToMoveIdx; j++ {
			// find block of ID=i

			blockToMove := blocks[blockToMoveIdx]
			if blocks[j].freeSpace >= blockToMove.size {
				blocks[blockToMoveIdx-1].freeSpace += blockToMove.size + blockToMove.freeSpace
				blockToMove.freeSpace = blocks[j].freeSpace - blockToMove.size
				blocks[j].freeSpace = 0

				blocks = slices.Delete(blocks, blockToMoveIdx, blockToMoveIdx+1)
				blocks = slices.Insert(blocks, j+1, blockToMove)
				break
			}
		}
	}
	return blocks
}

func BlocksToString(blocks []Block) []string {
	var result []string
	for _, block := range blocks {
		for i := 0; i < block.size; i++ {
			result = append(result, strconv.Itoa(block.id))
		}
		for i := 0; i < block.freeSpace; i++ {
			result = append(result, ".")
		}
	}
	return result
}

func CalculateChecksum(input []string) int {
	checksum := 0
	for i := 0; i < len(input); i++ {
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
	blocks := ConvertLayoutToBlocks(input[0])
	layout := BlocksToString(blocks)
	sortedLayout := SortLayout(layout)
	return CalculateChecksum(sortedLayout)
}

func Solve2(input []string) int {
	blocks := ConvertLayoutToBlocks(input[0])
	sorted := SortBlocks(blocks)

	checksum := 0
	i := 0
	for _, block := range sorted {
		for j := 0; j < block.size; j++ {
			checksum += i * block.id
			i++
		}
		if block.freeSpace > 0 {
			i += block.freeSpace
		}
	}

	return checksum
}
