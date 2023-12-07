package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHandlePart1(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`
	want := 6440
	got := handlePart1(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestReadInput(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`
	want := []Round{
		{"32T3K", 765},
		{"T55J5", 684},
		{"KK677", 28},
		{"KTJJT", 220},
		{"QQQJA", 483},
	}

	got := readInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSortRounds(t *testing.T) {
	input := []Round{
		{"32T3K", 765},
		{"T55J5", 684},
		{"KK677", 28},
		{"KTJJT", 220},
		{"QQQJA", 483},
	}
	want := []Round{
		{"32T3K", 765},
		{"KTJJT", 220},
		{"KK677", 28},
		{"T55J5", 684},
		{"QQQJA", 483},
	}
	got := sortRounds(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetTypeRank(t *testing.T) {
	for _, tt := range []struct {
		hand string
		want HandType
	}{
		{"32T3K", OnePair},
		{"T55J5", ThreeOfAKind},
		{"KK677", TwoPair},
		{"KTJJT", TwoPair},
		{"QQQJA", ThreeOfAKind},

		{"AAAAA", FiveOfAKind},
		{"AA8AA", FourOfAKind},
		{"23332", FullHouse},
		{"TTT98", ThreeOfAKind},
		{"23432", TwoPair},
		{"A23A4", OnePair},
		{"23456", HighCard},
	} {
		t.Run(fmt.Sprintf("%s -> %d", tt.hand, tt.want), func(t *testing.T) {
			got := getTypeRank(tt.hand)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestIsFiveOfAKind(t *testing.T) {
	for _, tt := range []struct {
		hand string
		want bool
	}{
		{"32T3K", false},
		{"T55J5", false},
		{"KK677", false},
		{"KTJJT", false},
		{"QQQJA", false},

		{"AAAAA", true},
		{"AA8AA", false},
		{"23332", false},
		{"TTT98", false},
		{"23432", false},
		{"A23A4", false},
		{"23456", false},
	} {
		t.Run(fmt.Sprintf("%s -> %t", tt.hand, tt.want), func(t *testing.T) {
			got := isFiveOfAKind(tt.hand)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}

func TestIsFourOfAKind(t *testing.T) {
	for _, tt := range []struct {
		hand string
		want bool
	}{
		{"32T3K", false},
		{"T55J5", false},
		{"KK677", false},
		{"KTJJT", false},
		{"QQQJA", false},

		{"AAAAA", false},
		{"AA8AA", true},
		{"23332", false},
		{"TTT98", false},
		{"23432", false},
		{"A23A4", false},
		{"23456", false},
	} {
		t.Run(fmt.Sprintf("%s -> %t", tt.hand, tt.want), func(t *testing.T) {
			got := isFourOfAKind(tt.hand)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}

func TestIsFullHouse(t *testing.T) {
	for _, tt := range []struct {
		hand string
		want bool
	}{
		{"32T3K", false},
		{"T55J5", false},
		{"KK677", false},
		{"KTJJT", false},
		{"QQQJA", false},

		{"AAAAA", false},
		{"AA8AA", false},
		{"23332", true},
		{"TTT98", false},
		{"23432", false},
		{"A23A4", false},
		{"23456", false},
	} {
		t.Run(fmt.Sprintf("%s -> %t", tt.hand, tt.want), func(t *testing.T) {
			got := isFullHouse(tt.hand)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}

func TestIsThreeOfAKind(t *testing.T) {
	for _, tt := range []struct {
		hand string
		want bool
	}{
		{"32T3K", false},
		{"T55J5", true},
		{"KK677", false},
		{"KTJJT", false},
		{"QQQJA", true},

		{"AAAAA", false},
		{"AA8AA", false},
		{"23332", false},
		{"TTT98", true},
		{"23432", false},
		{"A23A4", false},
		{"23456", false},
	} {
		t.Run(fmt.Sprintf("%s -> %t", tt.hand, tt.want), func(t *testing.T) {
			got := isThreeOfAKind(tt.hand)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}

func TestIsTwoPair(t *testing.T) {
	for _, tt := range []struct {
		hand string
		want bool
	}{
		{"32T3K", false},
		{"T55J5", false},
		{"KK677", true},
		{"KTJJT", true},
		{"QQQJA", false},

		{"AAAAA", false},
		{"AA8AA", false},
		{"23332", false},
		{"TTT98", false},
		{"23432", true},
		{"A23A4", false},
		{"23456", false},
	} {
		t.Run(fmt.Sprintf("%s -> %t", tt.hand, tt.want), func(t *testing.T) {
			got := isTwoPair(tt.hand)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}

func TestIsOnePair(t *testing.T) {
	for _, tt := range []struct {
		hand string
		want bool
	}{
		{"32T3K", true},
		{"T55J5", false},
		{"KK677", false},
		{"KTJJT", false},
		{"QQQJA", false},

		{"AAAAA", false},
		{"AA8AA", false},
		{"23332", false},
		{"TTT98", false},
		{"23432", false},
		{"A23A4", true},
		{"23456", false},
	} {
		t.Run(fmt.Sprintf("%s -> %t", tt.hand, tt.want), func(t *testing.T) {
			got := isOnePair(tt.hand)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}
