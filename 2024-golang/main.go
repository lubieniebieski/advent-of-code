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
	web := flag.Bool("web", false, "Run in web server mode")
	flag.Parse()

	if *web {
		utils.StartServer(8080)
	} else {
		results := utils.RunAllSolutions()
		for _, result := range results {
			utils.DisplayResults(result.Day, result.Part1, result.Part2)
		}
	}
}
