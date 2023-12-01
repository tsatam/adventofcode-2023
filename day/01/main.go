package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"

	"github.com/tsatam/adventofcode-2023/common/fp"
)

var (
	//go:embed input
	input     string
	notdigits []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
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

	return 10*int(line[firstIdx]-'0') + int(line[lastIdx]-'0')
}

func calibrationValueForLinePart2(line string) int {
	if len(line) == 0 {
		return 0
	}

	return 10*firstDigit(line) + lastDigit(line)
}

func firstDigit(line string) int {
	for i := 0; i < len(line); i++ {
		r := rune(line[i])
		if unicode.IsNumber(r) {
			return int(r - '0')
		}
		for digitIdx, phrase := range notdigits {
			if strings.Index(line, phrase) == i {
				return digitIdx + 1
			}
		}
	}
	return 0
}

func lastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		r := rune(line[i])
		if unicode.IsNumber(r) {
			return int(r - '0')
		}
		for digitIdx, phrase := range notdigits {
			idxOfPhrase := strings.LastIndex(line, phrase)
			if idxOfPhrase+len(phrase)-1 == i {
				return digitIdx + 1
			}
		}
	}
	return 0
}
