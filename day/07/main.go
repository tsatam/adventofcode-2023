package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/tsatam/adventofcode-2023/common/fp"
	"github.com/tsatam/adventofcode-2023/day/07/common"
	"github.com/tsatam/adventofcode-2023/day/07/joker"
	"github.com/tsatam/adventofcode-2023/day/07/vanilla"
)

var (
	//go:embed input
	input string
)

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	rounds := readInput(input)
	rounds = vanilla.SortRounds(rounds)
	result := 0
	for i, round := range rounds {
		result += round.Bid * (i + 1)
	}
	return result
}

func handlePart2(input string) int {
	rounds := readInput(input)
	rounds = joker.SortRounds(rounds)
	result := 0
	for i, round := range rounds {
		result += round.Bid * (i + 1)
	}
	return result
}

func readInput(input string) []common.Round {
	split := strings.Split(strings.TrimSpace(input), "\n")
	return fp.Map(split, func(line string) common.Round {
		var hand string
		var bid int
		if _, err := fmt.Sscanf(line, "%s %d", &hand, &bid); err != nil {
			panic(err)
		}
		return common.Round{Hand: hand, Bid: bid}
	})
}
