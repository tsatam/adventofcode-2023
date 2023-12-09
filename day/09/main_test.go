package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHandlePart1(t *testing.T) {
	input := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`
	want := 114
	got := handlePart1(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestHandlePart2(t *testing.T) {
	input := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`
	want := 2
	got := handlePart2(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestReadInput(t *testing.T) {
	input := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`
	want := [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}
	got := readInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLineToHistory(t *testing.T) {
	for _, tt := range []struct {
		input []int
		want  [][]int
	}{
		{
			input: []int{0, 3, 6, 9, 12, 15},
			want: [][]int{
				{0, 3, 6, 9, 12, 15},
				{3, 3, 3, 3, 3},
				{0, 0, 0, 0},
			},
		},
		{
			input: []int{1, 3, 6, 10, 15, 21},
			want: [][]int{
				{1, 3, 6, 10, 15, 21},
				{2, 3, 4, 5, 6},
				{1, 1, 1, 1},
				{0, 0, 0},
			},
		},
		{
			input: []int{10, 13, 16, 21, 30, 45},
			want: [][]int{
				{10, 13, 16, 21, 30, 45},
				{3, 3, 5, 9, 15},
				{0, 2, 4, 6},
				{2, 2, 2},
				{0, 0},
			},
		},
	} {
		t.Run(fmt.Sprintf("%v -> %v", tt.input, tt.want), func(t *testing.T) {

			got := lineToHistory(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtrapolateHistory(t *testing.T) {
	for _, tt := range []struct {
		input [][]int
		want  int
	}{
		{
			input: [][]int{
				{0, 3, 6, 9, 12, 15},
				{3, 3, 3, 3, 3},
				{0, 0, 0, 0},
			},
			want: 18,
		},
		{
			input: [][]int{
				{1, 3, 6, 10, 15, 21},
				{2, 3, 4, 5, 6},
				{1, 1, 1, 1},
				{0, 0, 0},
			},
			want: 28,
		},
		{
			input: [][]int{
				{10, 13, 16, 21, 30, 45},
				{3, 3, 5, 9, 15},
				{0, 2, 4, 6},
				{2, 2, 2},
				{0, 0},
			},
			want: 68,
		},
	} {
		t.Run(fmt.Sprintf("%v -> %d", tt.input, tt.want), func(t *testing.T) {
			got := extrapolateHistory(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtrapolateHistoryBackwards(t *testing.T) {
	for _, tt := range []struct {
		input [][]int
		want  int
	}{
		{
			input: [][]int{
				{0, 3, 6, 9, 12, 15},
				{3, 3, 3, 3, 3},
				{0, 0, 0, 0},
			},
			want: -3,
		},
		{
			input: [][]int{
				{1, 3, 6, 10, 15, 21},
				{2, 3, 4, 5, 6},
				{1, 1, 1, 1},
				{0, 0, 0},
			},
			want: 0,
		},
		{
			input: [][]int{
				{10, 13, 16, 21, 30, 45},
				{3, 3, 5, 9, 15},
				{0, 2, 4, 6},
				{2, 2, 2},
				{0, 0},
			},
			want: 5,
		},
	} {
		t.Run(fmt.Sprintf("%v -> %d", tt.input, tt.want), func(t *testing.T) {
			got := extrapolateHistoryBackwards(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
