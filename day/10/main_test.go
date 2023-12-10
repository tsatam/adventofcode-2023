package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHandlePart1(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  int
	}{
		{
			input: `.....
.S-7.
.|.|.
.L-J.
.....
`,
			want: 4,
		},
		{
			input: `..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`,
			want: 8,
		},
	} {
		t.Run(fmt.Sprintf("%s -> %d", tt.input, tt.want), func(t *testing.T) {
			got := handlePart1(tt.input)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestReadInput(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  [][]rune
	}{
		{
			input: `.....
.S-7.
.|.|.
.L-J.
.....
`,
			want: [][]rune{
				{'.', '.', '.', '.', '.'},
				{'.', 'S', '-', '7', '.'},
				{'.', '|', '.', '|', '.'},
				{'.', 'L', '-', 'J', '.'},
				{'.', '.', '.', '.', '.'},
			},
		},
		{
			input: `..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`,
			want: [][]rune{
				{'.', '.', 'F', '7', '.'},
				{'.', 'F', 'J', '|', '.'},
				{'S', 'J', '.', 'L', '7'},
				{'|', 'F', '-', '-', 'J'},
				{'L', 'J', '.', '.', '.'},
			},
		},
	} {
		t.Run(fmt.Sprintf("%s -> %v", tt.input, tt.want), func(t *testing.T) {
			got := readInput(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindDistances(t *testing.T) {
	for _, tt := range []struct {
		input [][]rune
		want  [][]int
	}{
		{
			input: [][]rune{
				{'.', '.', '.', '.', '.'},
				{'.', 'S', '-', '7', '.'},
				{'.', '|', '.', '|', '.'},
				{'.', 'L', '-', 'J', '.'},
				{'.', '.', '.', '.', '.'},
			},
			want: [][]int{
				{-1, -1, -1, -1, -1},
				{-1, 0, 1, 2, -1},
				{-1, 1, -1, 3, -1},
				{-1, 2, 3, 4, -1},
				{-1, -1, -1, -1, -1},
			},
		},
		{
			input: [][]rune{
				{'.', '.', 'F', '7', '.'},
				{'.', 'F', 'J', '|', '.'},
				{'S', 'J', '.', 'L', '7'},
				{'|', 'F', '-', '-', 'J'},
				{'L', 'J', '.', '.', '.'},
			},
			want: [][]int{
				{-1, -1, 4, 5, -1},
				{-1, 2, 3, 6, -1},
				{0, 1, -1, 7, 8},
				{1, 4, 5, 6, 7},
				{2, 3, -1, -1, -1},
			},
		},
	} {
		t.Run(fmt.Sprintf("%v -> %v", tt.input, tt.want), func(t *testing.T) {
			got := findDistances(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
