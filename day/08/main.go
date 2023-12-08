package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/tsatam/adventofcode-2023/common/fp"
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
	fmt.Printf("Part 2: %d\n", handlePart2(input))
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

func handlePart2(input string) int {
	networkMap := readInput(input)

	currentNodes := []string{}
	for name := range networkMap.nodes {
		if name[2] == 'A' {
			currentNodes = append(currentNodes, name)
		}
	}

	type Cycle struct {
		// c + nx => Z for all n
		c, x int
	}

	allCycles := fp.MapParallel(currentNodes, func(currentNode string) Cycle {
		first, second := -1, -1
		for i := 0; true; i++ {
			instruction := networkMap.instructions[i%len(networkMap.instructions)]
			nodePath := networkMap.nodes[currentNode]
			switch instruction {
			case 'L':
				currentNode = nodePath.left
			case 'R':
				currentNode = nodePath.right
			}

			if currentNode[2] == 'Z' {
				if first == -1 {
					first = i
					continue
				}
				second = i
				break
			}
		}

		return Cycle{
			c: first,
			x: second - first,
		}
	})

	largestCycle := slices.MaxFunc(allCycles, func(a, b Cycle) int {
		return cmp.Compare(a.x, b.x)
	})

	for i := largestCycle.c; i < math.MaxInt; i += largestCycle.x {
		if fp.AllMatch(allCycles, func(it Cycle) bool {
			return (i-it.c)%it.x == 0
		}) {
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
