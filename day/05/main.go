package main

import (
	_ "embed"
	"fmt"
	"math"
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
	fmt.Printf("Part 2: %d\n", handlePart2(input))
}

func handlePart1(input string) int {
	almanac := readInput(input)
	seeds := seedsToLocations(almanac)
	return slices.Min(seeds)
}

func handlePart2(input string) int {
	almanac := readInput(input)

	for location := 0; location < math.MaxInt; location++ { // keep going until we either find a valid location or melt my computer
		eventualSeed := location

		for i := len(almanac.maps) - 1; i >= 0; i-- {
			for _, m := range almanac.maps[i] {
				if eventualSeed >= m.dest && eventualSeed < m.dest+m.rng {
					diff := eventualSeed - m.dest
					eventualSeed = m.source + diff
					break
				}
			}
		}

		if seedInSeedRange(eventualSeed, almanac) {
			return location
		}
	}
	// my computer is on fire, maybe the elves will find the water to put it out
	return math.MaxInt
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
	return fp.MapParallel(almanac.seeds, func(seed int) int {
		location := seed

		for _, m := range almanac.maps {
			for _, c := range m {
				if c.source <= location && location < c.source+c.rng {
					diff := location - c.source
					location = c.dest + diff
					break
				}
			}
		}

		return location
	})
}

func seedInSeedRange(seed int, a Almanac) bool {
	for i := 0; i < len(a.seeds); i += 2 {
		if a.seeds[i] <= seed && seed < a.seeds[i]+a.seeds[i+1] {
			return true
		}
	}

	return false
}

func parseInt(raw string) int {
	n, err := strconv.ParseInt(raw, 10, 0)
	if err != nil {
		panic(err)
	}
	return int(n)
}
