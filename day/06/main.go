package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/tsatam/adventofcode-2023/common/fp"
)

var (
	//go:embed input
	input string
)

type Race struct {
	time, distance int
}

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
}

func handlePart1(input string) int {
	races := readInput(input)
	solutions := fp.Map(races, solveRace)
	return fp.Reduce(solutions, 1, func(curr, next int) int { return curr * next })
}

func readInput(input string) []Race {
	lines := strings.Split(input, "\n")

	entries := fp.Map(
		lines,
		func(l string) []int {
			entries := strings.Split(l, " ")
			entries = fp.Filter(entries, func(it string) bool { return it != "" })
			entries = fp.Filter(entries, func(it string) bool { return unicode.IsDigit(rune(it[0])) })
			return fp.Map(entries, parseInt)
		},
	)

	races := make([]Race, len(entries[0]))
	for i, t := range entries[0] {
		d := entries[1][i]
		races[i] = Race{time: t, distance: d}
	}

	return races
}

func solveRace(r Race) int {
	diff := math.Sqrt(float64((r.time * r.time) - (4 * r.distance)))

	min := int(math.Floor((float64(r.time) - diff) / 2))
	max := int(math.Ceil((float64(r.time) + diff) / 2))

	return max - min - 1
}

func parseInt(raw string) int {
	n, err := strconv.ParseInt(raw, 10, 0)
	if err != nil {
		panic(err)
	}
	return int(n)
}

/*
	thinking space

	For Race R, we have Time t and Minimum Distance d

	Want to find all possible Time-To-Hold n given the constraints:

	0 <= n <= t
	distance achieved is n * (t - n) = d, find the two bounds (range/"number of ways" will be the delta between)
	tn - n^2 = d
	-n^2 + tn - d = 0

	n = (t +- sqrt(t^2 - 4d))/2

*/
