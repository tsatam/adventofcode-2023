package common

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
	Hand string
	Bid  int
}
