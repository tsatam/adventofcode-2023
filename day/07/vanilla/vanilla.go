package vanilla

import (
	"slices"

	"github.com/tsatam/adventofcode-2023/day/07/common"
)

var (
	cardRank = map[rune]int{
		'2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8,
		'T': 9, 'J': 10, 'Q': 11, 'K': 12, 'A': 13,
	}
)

func SortRounds(rounds []common.Round) []common.Round {
	sorted := slices.Clone(rounds)

	slices.SortFunc(sorted, func(a, b common.Round) int {
		typeRankA, typeRankB := getTypeRank(a.Hand), getTypeRank(b.Hand)
		if typeRankA != typeRankB {
			return int(typeRankA - typeRankB)
		}

		for i := 0; i < 5; i++ {
			cardInPosA, cardInPosB := a.Hand[i], b.Hand[i]
			cardRankA, cardRankB := cardRank[rune(cardInPosA)], cardRank[rune(cardInPosB)]

			if cardRankA != cardRankB {
				return cardRankA - cardRankB
			}
		}

		return 0
	})

	return sorted
}

func getTypeRank(hand string) common.HandType {
	switch {
	case isFiveOfAKind(hand):
		return common.FiveOfAKind
	case isFourOfAKind(hand):
		return common.FourOfAKind
	case isFullHouse(hand):
		return common.FullHouse
	case isThreeOfAKind(hand):
		return common.ThreeOfAKind
	case isTwoPair(hand):
		return common.TwoPair
	case isOnePair(hand):
		return common.OnePair
	default:
		return common.HighCard
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
		if count == 4 {
			return true
		}
	}
	return false
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
