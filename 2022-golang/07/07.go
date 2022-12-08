package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lubieniebieski/advent-of-code/2022-golang/tools"
)

const cuttingPoint = 100000

type File struct {
	parentName string
	name       string
	size       int
}

func createFile(parent, name string, size int) File {
	return File{parentName: parent, name: name, size: size}
}

func RecursiveSum(parent string, list map[string]File) int {
	var total int
	for _, v := range list {
		if v.parentName == parent {
			if v.size == 0 {
				total = RecursiveSum(v.name, list)
			}
			total += v.size
		}
	}
	return total
}

func PartOne(stringsArray []string) (result int) {
	allFiles := make(map[string]File)
	allFiles["/"] = File{name: "/", parentName: "", size: 0}
	parentName := "/"
	for i := 2; i < len(stringsArray); i++ {
		var cmd = stringsArray[i]
		if cmd == "$ cd .." {
			parentName = allFiles[parentName].parentName
		} else if cmd[0:4] == "$ cd" {
			name := strings.Split(cmd, " ")[2]
			allFiles[name] = createFile(parentName, name, 0)
			parentName = name
			i += 1
		} else {
			split := strings.Split(cmd, " ")
			var size int
			var name = split[1]
			if split[0] != "dir" {
				size, _ = strconv.Atoi(split[0])
			}
			allFiles[name] = createFile(parentName, name, size)
		}
	}
	var sums = make(map[string]int)

	tmp := make(map[string]File)
	for k, v := range allFiles {
		if v.size >= cuttingPoint {

		} else {
			tmp[k] = v
		}

	}
	for _, v := range tmp {
		if v.size <= cuttingPoint {
			sums[v.name] = RecursiveSum(v.name, tmp)
		}

	}
	var total int
	for _, sum := range sums {
		if sum <= cuttingPoint {
			total += sum
		}
	}
	return total
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
