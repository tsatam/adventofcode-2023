package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"github.com/tsatam/adventofcode-2023/common/fp"
)

var (
	//go:embed input
	input    string
	cardRank = map[rune]int{
		'2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8,
		'T': 9, 'J': 10, 'Q': 11, 'K': 12, 'A': 13,
	}
)

type HandType int

const (
	FiveOfAKind  HandType = 7
	FourOfAKind  HandType = 6
	FullHouse    HandType = 5
	ThreeOfAKind HandType = 4
	TwoPair      HandType = 3
	OnePair      HandType = 2
	HighCard     HandType = 1
)

type Round struct {
	hand string
	bid  int
}

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
}

func handlePart1(input string) int {
	rounds := readInput(input)
	rounds = sortRounds(rounds)
	result := 0
	for i, round := range rounds {
		result += round.bid * (i + 1)
	}
	return result
}

func readInput(input string) []Round {
	split := strings.Split(strings.TrimSpace(input), "\n")
	return fp.Map(split, func(line string) Round {
		var hand string
		var bid int
		if _, err := fmt.Sscanf(line, "%s %d", &hand, &bid); err != nil {
			panic(err)
		}
		return Round{hand: hand, bid: bid}
	})
}

func sortRounds(rounds []Round) []Round {
	sorted := slices.Clone(rounds)

	slices.SortFunc(sorted, func(a, b Round) int {
		typeRankA, typeRankB := getTypeRank(a.hand), getTypeRank(b.hand)
		if typeRankA != typeRankB {
			return int(typeRankA - typeRankB)
		}

		for i := 0; i < 5; i++ {
			cardInPosA, cardInPosB := a.hand[i], b.hand[i]
			cardRankA, cardRankB := cardRank[rune(cardInPosA)], cardRank[rune(cardInPosB)]

			if cardRankA != cardRankB {
				return cardRankA - cardRankB
			}
		}

		return 0
	})

	return sorted
}

func getTypeRank(hand string) HandType {
	switch {
	case isFiveOfAKind(hand):
		return FiveOfAKind
	case isFourOfAKind(hand):
		return FourOfAKind
	case isFullHouse(hand):
		return FullHouse
	case isThreeOfAKind(hand):
		return ThreeOfAKind
	case isTwoPair(hand):
		return TwoPair
	case isOnePair(hand):
		return OnePair
	default:
		return HighCard
	}
}

func isFiveOfAKind(hand string) bool {
	compareTo := rune(hand[0])

	for _, c := range hand {
		if c != compareTo {
			return false
		}
	}

	return true
}

func isFourOfAKind(hand string) bool {
	cardCount := countCards(hand)

	if len(cardCount) != 2 {
		return false
	}
	for _, count := range cardCount {
		if !(count == 1 || count == 4) {
			return false
		}
	}
	return true
}

func isFullHouse(hand string) bool {
	cardCount := countCards(hand)

	if len(cardCount) != 2 {
		return false
	}
	for _, count := range cardCount {
		if !(count == 2 || count == 3) {
			return false
		}
	}
	return true
}

func isThreeOfAKind(hand string) bool {
	cardCount := countCards(hand)

	if len(cardCount) != 3 {
		return false
	}
	for _, count := range cardCount {
		if !(count == 1 || count == 3) {
			return false
		}
	}
	return true
}

func isTwoPair(hand string) bool {
	cardCount := countCards(hand)

	if len(cardCount) != 3 {
		return false
	}
	for _, count := range cardCount {
		if !(count == 1 || count == 2) {
			return false
		}
	}
	return true
}

func isOnePair(hand string) bool {
	cardCount := countCards(hand)

	if len(cardCount) != 4 {
		return false
	}
	for _, count := range cardCount {
		if !(count == 1 || count == 2) {
			return false
		}
	}
	return true
}

func countCards(hand string) map[rune]int {
	cardCount := map[rune]int{}
	for _, card := range hand {
		cardCount[card] = cardCount[card] + 1
	}
	return cardCount
}
