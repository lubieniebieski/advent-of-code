package main

import (
	"flag"

	"github.com/lubieniebieski/advent-of-code/2024-golang/day01"
	"github.com/lubieniebieski/advent-of-code/2024-golang/day02"
	"github.com/lubieniebieski/advent-of-code/2024-golang/day03"
	"github.com/lubieniebieski/advent-of-code/2024-golang/day04"
	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

func init() {
	utils.RegisterSolution(1, utils.Solution{Part1: day01.Solve1, Part2: day01.Solve2})
	utils.RegisterSolution(2, utils.Solution{Part1: day02.Solve1, Part2: day02.Solve2})
	utils.RegisterSolution(3, utils.Solution{Part1: day03.Solve1, Part2: day03.Solve2})
	utils.RegisterSolution(4, utils.Solution{Part1: day04.Solve1, Part2: day04.Solve2})

}

func main() {
	port := flag.Int("port", 8080, "Port to run the server on")
	flag.Parse()

	utils.StartServer(*port)
}
