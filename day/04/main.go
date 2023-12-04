package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/tsatam/adventofcode-2023/common/fp"
)

var (
	//go:embed input
	input string
)

type Card struct {
	id      int
	winning []int
	numbers []int
}

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	cards := readInput(input)
	scores := fp.Map(cards, cardScore)
	return fp.Sum(scores)
}

func handlePart2(input string) int {
	cards := readInput(input)

	countProcessed := 0
	toProcessCount := make([]int, len(cards))

	for i, _ := range cards {
		toProcessCount[i] = 1
	}

	for i := 0; i < len(cards); i++ {
		card := cards[i]
		count := toProcessCount[i]

		countProcessed += count

		matches := cardMatches(card)
		for j := 1; j <= matches; j++ {
			toProcessCount[i+j] += count
		}
	}

	return countProcessed
}

func readInput(input string) []Card {
	split := strings.Split(strings.TrimSpace(input), "\n")

	return fp.Map(split, readCard)
}

func readCard(line string) Card {
	splitId := strings.Split(line, ":")
	id := 0
	fmt.Sscanf(splitId[0], "Card %d", &id)

	splitNumbersFromWinning := strings.Split(splitId[1], "|")

	splitWinning := strings.Split(strings.TrimSpace(splitNumbersFromWinning[0]), " ")
	splitWinning = fp.Filter(splitWinning, isNotEmpty)
	winning := fp.Map(splitWinning, parseInt)

	splitNumbers := strings.Split(strings.TrimSpace(splitNumbersFromWinning[1]), " ")
	splitNumbers = fp.Filter(splitNumbers, isNotEmpty)
	numbers := fp.Map(splitNumbers, parseInt)

	return Card{
		id:      id,
		winning: winning,
		numbers: numbers,
	}
}

func cardMatches(card Card) int {
	matches := 0
	winning := make(map[int]struct{}, len(card.winning))
	for _, w := range card.winning {
		winning[w] = struct{}{}
	}
	for _, n := range card.numbers {
		if _, ok := winning[n]; ok {
			matches++
		}
	}
	return matches

}

func cardScore(card Card) int {
	matches := cardMatches(card)
	if matches == 0 {
		return 0
	}
	score := 1
	for i := 1; i < matches; i++ {
		score *= 2
	}
	return score
}

func parseInt(raw string) int {
	n, err := strconv.ParseInt(raw, 10, 0)
	if err != nil {
		panic(err)
	}
	return int(n)
}

func isNotEmpty(s string) bool {
	return s != ""
}
