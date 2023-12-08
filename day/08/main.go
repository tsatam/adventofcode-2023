package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed input
	input string
)

type NetworkMap struct {
	instructions string
	nodes        map[string]Node
}

type Node struct {
	left, right string
}

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
}

func handlePart1(input string) int {
	networkMap := readInput(input)

	currentNode := "AAA"
	for i := 0; true; i++ {
		instruction := networkMap.instructions[i%len(networkMap.instructions)]
		nodePath := networkMap.nodes[currentNode]
		switch instruction {
		case 'L':
			currentNode = nodePath.left
		case 'R':
			currentNode = nodePath.right
		}

		if currentNode == "ZZZ" {
			return i + 1
		}
	}
	return -1
}

func readInput(input string) NetworkMap {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	nodes := make(map[string]Node, len(lines)-2)

	for _, line := range lines[2:] {
		var name, left, right string
		if _, err := fmt.Sscanf(line, "%3s = (%3s, %3s)", &name, &left, &right); err != nil {
			panic(err)
		}
		nodes[name] = Node{left: left, right: right}
	}

	return NetworkMap{
		instructions: lines[0],
		nodes:        nodes,
	}
}
