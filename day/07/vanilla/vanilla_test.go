package vanilla

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/tsatam/adventofcode-2023/day/07/common"
)

func TestSortRounds(t *testing.T) {
	input := []common.Round{
		{Hand: "32T3K", Bid: 765},
		{Hand: "T55J5", Bid: 684},
		{Hand: "KK677", Bid: 28},
		{Hand: "KTJJT", Bid: 220},
		{Hand: "QQQJA", Bid: 483},
	}
	want := []common.Round{
		{Hand: "32T3K", Bid: 765},
		{Hand: "KTJJT", Bid: 220},
		{Hand: "KK677", Bid: 28},
		{Hand: "T55J5", Bid: 684},
		{Hand: "QQQJA", Bid: 483},
	}
	got := SortRounds(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetTypeRank(t *testing.T) {
	for _, tt := range []struct {
		hand string
		want common.HandType
	}{
		{"32T3K", common.OnePair},
		{"T55J5", common.ThreeOfAKind},
		{"KK677", common.TwoPair},
		{"KTJJT", common.TwoPair},
		{"QQQJA", common.ThreeOfAKind},

		{"AAAAA", common.FiveOfAKind},
		{"AA8AA", common.FourOfAKind},
		{"23332", common.FullHouse},
		{"TTT98", common.ThreeOfAKind},
		{"23432", common.TwoPair},
		{"A23A4", common.OnePair},
		{"23456", common.HighCard},
	} {
		t.Run(fmt.Sprintf("%s -> %d", tt.hand, tt.want), func(t *testing.T) {
			got := getTypeRank(tt.hand)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
