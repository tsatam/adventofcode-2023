package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/tsatam/adventofcode-2023/common/fp"
)

var (
	//go:embed input
	input string
)

func main() {
	fmt.Printf("%d\n", handle(input))
}

func handle(input string) int {
	lines := readInput(input)
	calibrationValues := fp.Map(lines, calibrationValueForLine)
	sum := fp.Reduce(calibrationValues, 0, func(curr, next int) int { return curr + next })
	return sum
}

func readInput(input string) []string {
	return strings.Split(input, "\n")
}

func calibrationValueForLine(line string) int {
	if len(line) == 0 {
		return 0
	}

	firstIdx := strings.IndexFunc(line, unicode.IsNumber)
	lastIdx := strings.LastIndexFunc(line, unicode.IsNumber)

	res, err := strconv.ParseInt(fmt.Sprintf("%c%c", line[firstIdx], line[lastIdx]), 10, 0)
	if err != nil {
		panic(err)
	}

	return int(res)
}
