package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHandlePart1(t *testing.T) {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`
	want := 13
	got := handlePart1(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestHandlePart2(t *testing.T) {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`
	want := 30
	got := handlePart2(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestReadInput(t *testing.T) {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`
	want := []Card{
		{id: 1, winning: []int{41, 48, 83, 86, 17}, numbers: []int{83, 86, 6, 31, 17, 9, 48, 53}},
		{id: 2, winning: []int{13, 32, 20, 16, 61}, numbers: []int{61, 30, 68, 82, 17, 32, 24, 19}},
		{id: 3, winning: []int{1, 21, 53, 59, 44}, numbers: []int{69, 82, 63, 72, 16, 21, 14, 1}},
		{id: 4, winning: []int{41, 92, 73, 84, 69}, numbers: []int{59, 84, 76, 51, 58, 5, 54, 83}},
		{id: 5, winning: []int{87, 83, 26, 28, 32}, numbers: []int{88, 30, 70, 12, 93, 22, 82, 36}},
		{id: 6, winning: []int{31, 18, 13, 56, 72}, numbers: []int{74, 77, 10, 23, 35, 67, 36, 11}},
	}

	got := readInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCardScore(t *testing.T) {
	for _, tt := range []struct {
		card Card
		want int
	}{
		{card: Card{id: 1, winning: []int{41, 48, 83, 86, 17}, numbers: []int{83, 86, 6, 31, 17, 9, 48, 53}}, want: 8},
		{card: Card{id: 2, winning: []int{13, 32, 20, 16, 61}, numbers: []int{61, 30, 68, 82, 17, 32, 24, 19}}, want: 2},
		{card: Card{id: 3, winning: []int{1, 21, 53, 59, 44}, numbers: []int{69, 82, 63, 72, 16, 21, 14, 1}}, want: 2},
		{card: Card{id: 4, winning: []int{41, 92, 73, 84, 69}, numbers: []int{59, 84, 76, 51, 58, 5, 54, 83}}, want: 1},
		{card: Card{id: 5, winning: []int{87, 83, 26, 28, 32}, numbers: []int{88, 30, 70, 12, 93, 22, 82, 36}}, want: 0},
		{card: Card{id: 6, winning: []int{31, 18, 13, 56, 72}, numbers: []int{74, 77, 10, 23, 35, 67, 36, 11}}, want: 0},
	} {
		t.Run(fmt.Sprintf("Card %d has scord %d", tt.card.id, tt.want), func(t *testing.T) {
			got := cardScore(tt.card)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCardMatches(t *testing.T) {
	for _, tt := range []struct {
		card Card
		want int
	}{
		{card: Card{id: 1, winning: []int{41, 48, 83, 86, 17}, numbers: []int{83, 86, 6, 31, 17, 9, 48, 53}}, want: 4},
		{card: Card{id: 2, winning: []int{13, 32, 20, 16, 61}, numbers: []int{61, 30, 68, 82, 17, 32, 24, 19}}, want: 2},
		{card: Card{id: 3, winning: []int{1, 21, 53, 59, 44}, numbers: []int{69, 82, 63, 72, 16, 21, 14, 1}}, want: 2},
		{card: Card{id: 4, winning: []int{41, 92, 73, 84, 69}, numbers: []int{59, 84, 76, 51, 58, 5, 54, 83}}, want: 1},
		{card: Card{id: 5, winning: []int{87, 83, 26, 28, 32}, numbers: []int{88, 30, 70, 12, 93, 22, 82, 36}}, want: 0},
		{card: Card{id: 6, winning: []int{31, 18, 13, 56, 72}, numbers: []int{74, 77, 10, 23, 35, 67, 36, 11}}, want: 0},
	} {
		t.Run(fmt.Sprintf("Card %d has scord %d", tt.card.id, tt.want), func(t *testing.T) {
			got := cardMatches(tt.card)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
