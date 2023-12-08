package main

import (
	"fmt"
	"os"
	"time"

	"github.com/lubieniebieski/advent-of-code/2023-golang/tools"
)

func PartOne(input string) (result int) {
	order, nodes := PrepareNodes(input)
	startNode := nodes["AAA"]
	return FindZZZ(order, startNode, result)
}

func PartTwo(input string) (result int) {
	return result
}

func parsedData() string {
	rawData, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(rawData)
}

func main() {
	data := parsedData()
	start := time.Now()
	fmt.Printf("Part One: %d \n", PartOne(data))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Printf("Part Two: %d \n", PartTwo(data))
	fmt.Println(time.Since(start))
}

// Main code

type Node struct {
	ID    string
	Left  *Node
	Right *Node
}

func NodeFromString(input string) Node {
	return Node{
		ID: input[0:3],
		Left: &Node{
			ID: input[7:10],
		},
		Right: &Node{
			ID: input[12:15],
		},
	}
}

func PrepareNodes(input string) (string, map[string]*Node) {
	strArray := tools.ExtractStringsFromString(input)
	order := strArray[0]
	nodes := make(map[string]*Node)
	for _, line := range strArray[2:] {
		node := NodeFromString(line)
		nodes[node.ID] = &node
	}
	for _, node := range nodes {
		node.Left = nodes[node.Left.ID]
		node.Right = nodes[node.Right.ID]
	}
	return order, nodes
}

func FindZZZ(order string, currentNode *Node, result int) int {
	i := 0
	for i < len(order) {
		if currentNode.ID == "ZZZ" {
			return result
		}
		if string(order[i]) == "L" {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
		if i == len(order)-1 {
			i = 0
		} else {
			i++
		}
		result++
	}
	return result
}
