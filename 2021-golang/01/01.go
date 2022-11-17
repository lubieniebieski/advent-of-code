package main

import (
	"fmt"

	"github.com/lubieniebieski/advent-of-code/2022/tools"
)

func PartOne(nums []int) (int, error) {
  increasedCounter := 0
  for i := 1; i < len(nums); i++ {
    if nums[i] > nums[i-1] {
      increasedCounter++
    }
  }
  return increasedCounter, nil
}

func main() {
  nums := tools.ParseIntegersFromFile("input.txt")
  answer, err := PartOne(nums)

  if err != nil {
    panic(err)
  }

  fmt.Printf("Part One: %d \n", answer)
}
