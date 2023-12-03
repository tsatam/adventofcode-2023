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
	input string
)

type Point struct {
	X, Y int
}

// coords retrieved via y,x
type Schematic [][]rune
type PartNumbers [][]int

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
}

func handlePart1(input string) int {
	schematic := readInput(input)
	partNumbers := getPartNumbers(schematic)
	return sumPartNumbers(partNumbers)
}

func readInput(input string) Schematic {
	split := strings.Split(strings.TrimSpace(input), "\n")

	schematic := make(Schematic, len(split))

	for y, row := range split {
		schematic[y] = []rune(row)
	}

	return schematic
}

func getPartNumbers(s Schematic) PartNumbers {
	partNumbers := make(PartNumbers, len(s))
	for y := range partNumbers {
		partNumbers[y] = make([]int, len(s[y]))
	}

	for y, row := range s {
		for x, r := range row {
			if isSymbol(r) {
				setPartNumbersAroundSymbol(s, Point{X: x, Y: y}, partNumbers)
			}
		}
	}
	return partNumbers
}

func setPartNumbersAroundSymbol(s Schematic, p Point, partNumbers PartNumbers) {
	adjacent := []int{-1, 0, 1}
	for _, dy := range adjacent {
		for _, dx := range adjacent {
			y := p.Y + dy
			x := p.X + dx

			outOfBounds := y < 0 || y >= len(s) || x < 0 || x >= len(s[y])
			if outOfBounds || !unicode.IsDigit(s[y][x]) {
				continue
			}

			partNumber := 0

			for ; x < len(s[y]) && unicode.IsDigit(s[y][x]); x++ {
			}
			x--

			base := 1

			for ; x >= 0 && unicode.IsDigit(s[y][x]); x-- {
				partNumber += base * int(s[y][x]-'0')
				base *= 10
			}
			x++

			partNumbers[y][x] = partNumber
		}
	}
}

func sumPartNumbers(p PartNumbers) int {
	return fp.Sum(fp.Map(p, fp.Sum))
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}
