package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/tsatam/adventofcode-2023/common/fp"
	"golang.org/x/exp/slices"
)

var (
	//go:embed input
	input string
)

type Almanac struct {
	seeds []int
	maps  [][]AlmanacMap
}

type AlmanacMap struct {
	dest, source, rng int
}

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
}

func handlePart1(input string) int {
	almanac := readInput(input)
	seeds := seedsToLocations(almanac)
	return slices.Min(seeds)
}

func readInput(input string) Almanac {
	sections := strings.Split(input, "\n\n")
	seedsLine := strings.Split(sections[0], " ")[1:]
	seeds := fp.Map(seedsLine, parseInt)

	maps := make([][]AlmanacMap, len(sections)-1)

	for i, rawMap := range sections[1:] {
		lines := strings.Split(strings.TrimSpace(rawMap), "\n")
		maps[i] = make([]AlmanacMap, len(lines)-1)
		for j, rawLine := range lines[1:] {
			var dest, source, rng int
			if _, err := fmt.Sscanf(rawLine, "%d %d %d", &dest, &source, &rng); err != nil {
				panic(err)
			}
			maps[i][j] = AlmanacMap{dest, source, rng}
		}
	}

	return Almanac{
		seeds: seeds,
		maps:  maps,
	}
}

func seedsToLocations(almanac Almanac) []int {
	return fp.Map(almanac.seeds, func(seed int) int {
		location := seed

		for _, m := range almanac.maps {
			for _, c := range m {
				if location >= c.source && location < c.source+c.rng {
					diff := location - c.source
					location = c.dest + diff
					break
				}
			}
		}

		return location
	})
}

func parseInt(raw string) int {
	n, err := strconv.ParseInt(raw, 10, 0)
	if err != nil {
		panic(err)
	}
	return int(n)
}
