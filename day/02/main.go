package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/tsatam/adventofcode-2023/common/fp"
)

var (
	//go:embed input
	input string
)

type Game struct {
	id     int
	rounds []Round
}

type Round struct {
	red, green, blue int
}

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
}

func handlePart1(input string) int {
	games := readInput(input)
	possibleGames := fp.Filter(games, isGamePossible)
	idSums := fp.Reduce(possibleGames, 0, func(curr int, next Game) int {
		return curr + next.id
	})
	return idSums
}

func readInput(input string) []Game {
	rawSplit := strings.Split(strings.TrimSpace(input), "\n")
	return fp.Map(rawSplit, readLine)
}

func readLine(line string) Game {
	split := strings.Split(line, ": ")
	var id int
	if _, err := fmt.Sscanf(split[0], "Game %d", &id); err != nil {
		panic(err)
	}

	rawRounds := strings.Split(split[1], "; ")
	rounds := fp.Map(rawRounds, readRound)

	return Game{
		id:     id,
		rounds: rounds,
	}
}

func readRound(round string) Round {
	r := Round{}
	for _, rawColorCount := range strings.Split(round, ", ") {
		var count int
		var color string
		if _, err := fmt.Sscanf(rawColorCount, "%d %s", &count, &color); err != nil {
			panic(err)
		}

		switch color {
		case "red":
			r.red = count
		case "blue":
			r.blue = count
		case "green":
			r.green = count
		default:
			panic("invalid color " + color)
		}
	}
	return r
}

func isGamePossible(g Game) bool {
	for _, r := range g.rounds {
		if r.red > 12 || r.green > 13 || r.blue > 14 {
			return false
		}
	}
	return true
}
