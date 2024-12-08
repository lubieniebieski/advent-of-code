package main

import (
	"flag"
	"strconv"

	"github.com/lubieniebieski/advent-of-code/2024-golang/day01"
	"github.com/lubieniebieski/advent-of-code/2024-golang/day02"
	"github.com/lubieniebieski/advent-of-code/2024-golang/day03"
	"github.com/lubieniebieski/advent-of-code/2024-golang/day04"
	"github.com/lubieniebieski/advent-of-code/2024-golang/day05"
	"github.com/lubieniebieski/advent-of-code/2024-golang/day06"
	"github.com/lubieniebieski/advent-of-code/2024-golang/day07"
	"github.com/lubieniebieski/advent-of-code/2024-golang/day08"

	"github.com/lubieniebieski/advent-of-code/2024-golang/utils"
)

func init() {
	utils.RegisterSolution(1, utils.Solution{Part1: day01.Solve1, Part2: day01.Solve2})
	utils.RegisterSolution(2, utils.Solution{Part1: day02.Solve1, Part2: day02.Solve2})
	utils.RegisterSolution(3, utils.Solution{Part1: day03.Solve1, Part2: day03.Solve2})
	utils.RegisterSolution(4, utils.Solution{Part1: day04.Solve1, Part2: day04.Solve2})
	utils.RegisterSolution(5, utils.Solution{Part1: day05.Solve1, Part2: day05.Solve2})
	utils.RegisterSolution(6, utils.Solution{Part1: day06.Solve1, Part2: day06.Solve2})
	utils.RegisterSolution(7, utils.Solution{Part1: day07.Solve1, Part2: day07.Solve2})
	utils.RegisterSolution(8, utils.Solution{Part1: day08.Solve1, Part2: day08.Solve2})

}

func main() {
	web := flag.Bool("web", false, "Run in web server mode")
	flag.Parse()

	if *web {
		utils.StartServer(8080)
		return
	}

	// If there's a non-flag argument, treat it as the day number
	if flag.NArg() > 0 {
		if day, err := strconv.Atoi(flag.Arg(0)); err == nil {
			result, _ := utils.RunDay(day)
			utils.DisplayResults(result.Day, result.Part1, result.Part2)
			return
		}
	}

	// If no specific day provided, run all solutions
	results := utils.RunAllSolutions()
	for _, result := range results {
		utils.DisplayResults(result.Day, result.Part1, result.Part2)
	}
}
