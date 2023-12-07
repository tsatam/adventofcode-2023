package joker

import (
	"slices"

	"github.com/tsatam/adventofcode-2023/day/07/common"
)

var (
	cardRank = map[rune]int{
		'J': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'T': 10, 'Q': 11, 'K': 12, 'A': 13,
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
	cardCount := countCards(hand)

	if cardCount['J'] == 5 {
		return true
	}
	for c, count := range cardCount {
		if c == 'J' {
			continue
		}

		if count+cardCount['J'] == 5 {
			return true
		}
	}

	return false
}

func isFourOfAKind(hand string) bool {
	cardCount := countCards(hand)

	for c, count := range cardCount {
		if c == 'J' {
			continue
		}
		if count+cardCount['J'] == 4 {
			return true
		}
	}
	return false
}

func isFullHouse(hand string) bool {
	cardCount := countCards(hand)

	switch cardCount['J'] {
	case 2, 1:
		return len(cardCount) == 3
	default:
		return len(cardCount) == 2
	}
}

func isThreeOfAKind(hand string) bool {
	cardCount := countCards(hand)

	for c, count := range cardCount {
		if c == 'J' {
			continue
		}
		if count+cardCount['J'] == 3 {
			return true
		}
	}
	return false
}

// any 2 pair made with jokers could instead be a 3 of a kind
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

	for c, count := range cardCount {
		if c == 'J' {
			continue
		}
		if count+cardCount['J'] == 2 {
			return true
		}
	}
	return false
}

func countCards(hand string) map[rune]int {
	cardCount := map[rune]int{}
	for _, card := range hand {
		cardCount[card] = cardCount[card] + 1
	}
	return cardCount
}
