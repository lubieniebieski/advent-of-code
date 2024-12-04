package main

import (
	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

func main() {
	utils.Run(DAY, utils.Solution{
		Part1: Solve1,
		Part2: Solve2,
	})
}
