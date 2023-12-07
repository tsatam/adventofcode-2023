package main

import (
	"reflect"
	"testing"

	. "github.com/tsatam/adventofcode-2023/day/07/common"
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

func TestHandlePart2(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`
	want := 5905
	got := handlePart2(input)

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
		{Hand: "32T3K", Bid: 765},
		{Hand: "T55J5", Bid: 684},
		{Hand: "KK677", Bid: 28},
		{Hand: "KTJJT", Bid: 220},
		{Hand: "QQQJA", Bid: 483},
	}

	got := readInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
