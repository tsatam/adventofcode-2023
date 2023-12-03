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
type GearRatios [][]int

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	schematic := readInput(input)
	partNumbers := getPartNumbers(schematic)
	return sum(partNumbers)
}

func handlePart2(input string) int {
	schematic := readInput(input)
	gearRatios := getGearRatios(schematic)
	return sum(gearRatios)
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

func sum(m [][]int) int {
	return fp.Sum(fp.Map(m, fp.Sum))
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}

func getGearRatios(s Schematic) GearRatios {
	gearRatios := make(GearRatios, len(s))
	for y := range gearRatios {
		gearRatios[y] = make([]int, len(s[y]))
	}

	for y, row := range s {
		for x, r := range row {
			if r == '*' {
				setGearRatio(s, Point{X: x, Y: y}, gearRatios)
			}
		}
	}
	return gearRatios
}

func setGearRatio(s Schematic, gear Point, gearRatios GearRatios) {
	partNumbersFound := map[Point]int{}

	adjacent := []int{-1, 0, 1}
	for _, dy := range adjacent {
		for _, dx := range adjacent {
			y := gear.Y + dy
			x := gear.X + dx

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

			partNumbersFound[Point{X: x, Y: y}] = partNumber
		}
	}

	if len(partNumbersFound) == 2 {
		result := 1
		for _, n := range partNumbersFound {
			result *= n
		}
		gearRatios[gear.Y][gear.X] = result
	}
}
