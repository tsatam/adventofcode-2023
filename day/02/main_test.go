package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHandlePart1(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

	want := 8
	got := handlePart1(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestHandlePart2(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

	want := 2286
	got := handlePart2(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestReadLine(t *testing.T) {
	for _, tt := range []struct {
		line string
		want Game
	}{
		{
			line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			want: Game{
				id: 1,
				rounds: []Round{
					{blue: 3, red: 4},
					{red: 1, green: 2, blue: 6},
					{green: 2},
				},
			},
		},
		{
			line: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			want: Game{
				id: 2,
				rounds: []Round{
					{blue: 1, green: 2},
					{green: 3, blue: 4, red: 1},
					{green: 1, blue: 1},
				},
			},
		},
		{
			line: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			want: Game{
				id: 3,
				rounds: []Round{
					{green: 8, blue: 6, red: 20},
					{blue: 5, red: 4, green: 13},
					{green: 5, red: 1},
				},
			},
		},
		{
			line: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			want: Game{
				id: 4,
				rounds: []Round{
					{green: 1, red: 3, blue: 6},
					{green: 3, red: 6},
					{green: 3, blue: 15, red: 14},
				},
			},
		},
		{
			line: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			want: Game{
				id: 5,
				rounds: []Round{
					{red: 6, blue: 1, green: 3},
					{blue: 2, red: 1, green: 2},
				},
			},
		},
	} {
		t.Run(fmt.Sprintf("[%s] -> %v", tt.line, tt.want), func(t *testing.T) {
			got := readLine(tt.line)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsGamePossible(t *testing.T) {
	for _, tt := range []struct {
		game Game
		want bool
	}{
		{
			game: Game{
				id: 1,
				rounds: []Round{
					{blue: 3, red: 4},
					{red: 1, green: 2, blue: 6},
					{green: 2},
				},
			},
			want: true,
		},
		{
			game: Game{
				id: 2,
				rounds: []Round{
					{blue: 1, green: 2},
					{green: 3, blue: 4, red: 1},
					{green: 1, blue: 1},
				},
			},
			want: true,
		},
		{
			game: Game{
				id: 3,
				rounds: []Round{
					{green: 8, blue: 6, red: 20},
					{blue: 5, red: 4, green: 13},
					{green: 5, red: 1},
				},
			},
			want: false,
		},
		{
			game: Game{
				id: 4,
				rounds: []Round{
					{green: 1, red: 3, blue: 6},
					{green: 3, red: 6},
					{green: 3, blue: 15, red: 14},
				},
			},
			want: false,
		},
		{
			game: Game{
				id: 5,
				rounds: []Round{
					{red: 6, blue: 1, green: 3},
					{blue: 2, red: 1, green: 2},
				},
			},
			want: true,
		},
	} {
		t.Run(fmt.Sprintf("Game %d -> %v", tt.game.id, tt.want), func(t *testing.T) {
			got := isGamePossible(tt.game)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGamePower(t *testing.T) {
	for _, tt := range []struct {
		game Game
		want int
	}{
		{
			game: Game{
				id: 1,
				rounds: []Round{
					{blue: 3, red: 4},
					{red: 1, green: 2, blue: 6},
					{green: 2},
				},
			},
			want: 48,
		},
		{
			game: Game{
				id: 2,
				rounds: []Round{
					{blue: 1, green: 2},
					{green: 3, blue: 4, red: 1},
					{green: 1, blue: 1},
				},
			},
			want: 12,
		},
		{
			game: Game{
				id: 3,
				rounds: []Round{
					{green: 8, blue: 6, red: 20},
					{blue: 5, red: 4, green: 13},
					{green: 5, red: 1},
				},
			},
			want: 1560,
		},
		{
			game: Game{
				id: 4,
				rounds: []Round{
					{green: 1, red: 3, blue: 6},
					{green: 3, red: 6},
					{green: 3, blue: 15, red: 14},
				},
			},
			want: 630,
		},
		{
			game: Game{
				id: 5,
				rounds: []Round{
					{red: 6, blue: 1, green: 3},
					{blue: 2, red: 1, green: 2},
				},
			},
			want: 36,
		},
	} {
		t.Run(fmt.Sprintf("Game %d -> %v", tt.game.id, tt.want), func(t *testing.T) {
			got := gamePower(tt.game)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
