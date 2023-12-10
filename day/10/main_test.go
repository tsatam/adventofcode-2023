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

func TestHandlePart2(t *testing.T) {
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
			want: 1,
		},
		{
			input: `..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`,
			want: 1,
		},
		{
			input: `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`,
			want: 4,
		},
		{
			input: `..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........
`,
			want: 4,
		},
	} {
		t.Run(fmt.Sprintf("%s -> %d", tt.input, tt.want), func(t *testing.T) {
			got := handlePart2(tt.input)

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

func TestFillOutsideEmpty(t *testing.T) {
	for _, tt := range []struct {
		input [][]int
		want  [][]int
	}{
		{
			input: [][]int{
				{-1, -1, -1, -1, -1},
				{-1, 0, 1, 2, -1},
				{-1, 1, -1, 3, -1},
				{-1, 2, 3, 4, -1},
				{-1, -1, -1, -1, -1},
			},
			want: [][]int{
				{-2, -2, -2, -2, -2},
				{-2, 0, 1, 2, -2},
				{-2, 1, -1, 3, -2},
				{-2, 2, 3, 4, -2},
				{-2, -2, -2, -2, -2},
			},
		},
		{
			input: [][]int{
				{-1, -1, 4, 5, -1},
				{-1, 2, 3, 6, -1},
				{0, 1, -1, 7, 8},
				{1, 4, 5, 6, 7},
				{2, 3, -1, -1, -1},
			},
			want: [][]int{
				{-2, -2, 4, 5, -2},
				{-2, 2, 3, 6, -2},
				{0, 1, -1, 7, 8},
				{1, 4, 5, 6, 7},
				{2, 3, -2, -2, -2},
			},
		},
		{
			input: [][]int{
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, 0, 1, 2, 3, 4, 5, 6, 7, -1},
				{-1, 1, 16, 17, 18, 19, 20, 21, 8, -1},
				{-1, 2, 15, -1, -1, -1, -1, 22, 9, -1},
				{-1, 3, 14, -1, -1, -1, -1, 21, 10, -1},
				{-1, 4, 13, 12, 11, 18, 19, 20, 11, -1},
				{-1, 5, -1, -1, 10, 17, -1, -1, 12, -1},
				{-1, 6, 7, 8, 9, 16, 15, 14, 13, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			},
			want: [][]int{
				{-2, -2, -2, -2, -2, -2, -2, -2, -2, -2},
				{-2, 0, 1, 2, 3, 4, 5, 6, 7, -2},
				{-2, 1, 16, 17, 18, 19, 20, 21, 8, -2},
				{-2, 2, 15, -2, -2, -2, -2, 22, 9, -2},
				{-2, 3, 14, -2, -2, -2, -2, 21, 10, -2},
				{-2, 4, 13, 12, 11, 18, 19, 20, 11, -2},
				{-2, 5, -1, -1, 10, 17, -1, -1, 12, -2},
				{-2, 6, 7, 8, 9, 16, 15, 14, 13, -2},
				{-2, -2, -2, -2, -2, -2, -2, -2, -2, -2},
			},
		},
	} {
		t.Run(fmt.Sprintf("%v -> %v", tt.input, tt.want), func(t *testing.T) {
			got := fillOutsideEmpty(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
