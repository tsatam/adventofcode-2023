package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHandlePart1(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200
`
	want := 288
	got := handlePart1(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestReadInput(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200
`
	want := []Race{
		{time: 7, distance: 9},
		{time: 15, distance: 40},
		{time: 30, distance: 200},
	}

	got := readInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSolveRace(t *testing.T) {
	for _, tt := range []struct {
		input Race
		want  int
	}{
		{input: Race{time: 7, distance: 9}, want: 4},
		{input: Race{time: 15, distance: 40}, want: 8},
		{input: Race{time: 30, distance: 200}, want: 9},
	} {
		t.Run(fmt.Sprintf("Race %v -> %d", tt.input, tt.want), func(t *testing.T) {
			got := solveRace(tt.input)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
