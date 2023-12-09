package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/tsatam/adventofcode-2023/common/fp"
)

var (
	//go:embed input
	input string
)

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
}

func handlePart1(input string) int {
	lines := readInput(input)
	histories := fp.Map(lines, lineToHistory)
	extrapolated := fp.Map(histories, extrapolateHistory)
	return fp.Sum(extrapolated)
}

func readInput(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return fp.Map(lines, func(line string) []int {
		split := strings.Split(line, " ")
		return fp.Map(split, parseInt)
	})
}

func lineToHistory(line []int) [][]int {
	history := make([][]int, 1)
	history[0] = line

	lastRowIsAllZero := false

	for i := 0; !lastRowIsAllZero; i++ {
		lastRow := history[i]
		nextRow := make([]int, len(lastRow)-1)

		for j := range nextRow {
			nextRow[j] = lastRow[j+1] - lastRow[j]
		}

		history = append(history, nextRow)
		lastRowIsAllZero = fp.AllMatch(nextRow, func(it int) bool { return it == 0 })
	}

	return history
}

func extrapolateHistory(history [][]int) int {
	history[len(history)-1] = append(history[len(history)-1], 0)
	for i := len(history) - 2; i >= 0; i-- {
		history[i] = append(history[i], history[i][len(history[i])-1]+history[i+1][len(history[i+1])-1])
	}
	return history[0][len(history[0])-1]
}

func parseInt(raw string) int {
	n, err := strconv.ParseInt(raw, 10, 0)
	if err != nil {
		panic(err)
	}
	return int(n)
}
