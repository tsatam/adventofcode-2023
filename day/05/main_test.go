package main

import (
	"reflect"
	"testing"
)

func TestHandlePart1(t *testing.T) {
	input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`
	want := 35
	got := handlePart1(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestHandlePart2(t *testing.T) {
	input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`
	want := 46
	got := handlePart2(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestReadInput(t *testing.T) {
	input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`
	want := Almanac{
		seeds: []int{79, 14, 55, 13},
		maps: [][]AlmanacMap{
			{
				{50, 98, 2},
				{52, 50, 48},
			},
			{
				{0, 15, 37},
				{37, 52, 2},
				{39, 0, 15},
			},
			{
				{49, 53, 8},
				{0, 11, 42},
				{42, 0, 7},
				{57, 7, 4},
			},
			{
				{88, 18, 7},
				{18, 25, 70},
			},
			{
				{45, 77, 23},
				{81, 45, 19},
				{68, 64, 13},
			},
			{
				{0, 69, 1},
				{1, 0, 69},
			},
			{
				{60, 56, 37},
				{56, 93, 4},
			},
		},
	}

	got := readInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSeedsToLocations(t *testing.T) {
	input := Almanac{
		seeds: []int{79, 14, 55, 13},
		maps: [][]AlmanacMap{
			{
				{50, 98, 2},
				{52, 50, 48},
			},
			{
				{0, 15, 37},
				{37, 52, 2},
				{39, 0, 15},
			},
			{
				{49, 53, 8},
				{0, 11, 42},
				{42, 0, 7},
				{57, 7, 4},
			},
			{
				{88, 18, 7},
				{18, 25, 70},
			},
			{
				{45, 77, 23},
				{81, 45, 19},
				{68, 64, 13},
			},
			{
				{0, 69, 1},
				{1, 0, 69},
			},
			{
				{60, 56, 37},
				{56, 93, 4},
			},
		},
	}
	want := []int{82, 43, 86, 35}
	got := seedsToLocations(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
