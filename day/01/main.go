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
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	lines := readInput(input)
	calibrationValues := fp.Map(lines, calibrationValueForLine)
	sum := fp.Reduce(calibrationValues, 0, func(curr, next int) int { return curr + next })
	return sum
}

func handlePart2(input string) int {
	lines := readInput(input)
	calibrationValues := fp.Map(lines, calibrationValueForLinePart2)
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

func calibrationValueForLinePart2(line string) int {
	if len(line) == 0 {
		return 0
	}

	firstIdx := len(line)
	firstVal := 0
	lastIdx := -1
	lastVal := 0

	firstIdxLiteral := strings.IndexFunc(line, unicode.IsNumber)
	if firstIdxLiteral >= 0 {
		firstIdx = firstIdxLiteral
		firstValTemp, err := strconv.ParseInt(string(line[firstIdxLiteral]), 10, 0)
		if err != nil {
			panic(err)
		}
		firstVal = int(firstValTemp)
	}
	lastIdxLiteral := strings.LastIndexFunc(line, unicode.IsNumber)
	if lastIdxLiteral >= 0 {
		lastIdx = lastIdxLiteral
		lastValTemp, err := strconv.ParseInt(string(line[lastIdxLiteral]), 10, 0)
		if err != nil {
			panic(err)
		}
		lastVal = int(lastValTemp)
	}
	for i, thesearentdigits := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
		firstIdxNotDigit := strings.Index(line, thesearentdigits)
		if firstIdxNotDigit >= 0 && firstIdxNotDigit < firstIdx {
			firstIdx = firstIdxNotDigit
			firstVal = i + 1
		}

		lastIdxNotDigit := strings.LastIndex(line, thesearentdigits)
		if lastIdxNotDigit >= 0 && lastIdxNotDigit > lastIdx {
			lastIdx = lastIdxNotDigit
			lastVal = i + 1
		}
	}

	return 10*firstVal + lastVal
}
