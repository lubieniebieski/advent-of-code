package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
)

const cuttingPoint = 100000

func PartOne(stringsArray []string) (result int) {
	sizes := []int{0}
	for i := 2; i < len(stringsArray); i++ {
		var cmd = stringsArray[i]
		if cmd == "$ cd .." {
			dirSize := sizes[len(sizes)-1]
			if dirSize < cuttingPoint {
				result += dirSize
			}
			sizes = sizes[:len(sizes)-1]
			sizes[len(sizes)-1] += dirSize

		} else if cmd[0:4] == "$ cd" {
			sizes = append(sizes, 0)
			i += 1
		} else {
			split := strings.Split(cmd, " ")
			var size int
			if split[0] != "dir" {
				size, _ = strconv.Atoi(split[0])
			}
			sizes[len(sizes)-1] += size
		}
	}

	return result
}

func PartTwo(stringsArray []string) (result int) {
	return 0
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
