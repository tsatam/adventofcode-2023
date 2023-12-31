package main

import (
	"reflect"
	"testing"
)

func TestHandlePart1(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

	want := 4361
	got := handlePart1(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestHandlePart2(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

	want := 467835
	got := handlePart2(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestReadInput(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`
	want := Schematic{
		{'4', '6', '7', '.', '.', '1', '1', '4', '.', '.'},
		{'.', '.', '.', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '3', '5', '.', '.', '6', '3', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
		{'6', '1', '7', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '+', '.', '5', '8', '.'},
		{'.', '.', '5', '9', '2', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '7', '5', '5', '.'},
		{'.', '.', '.', '$', '.', '*', '.', '.', '.', '.'},
		{'.', '6', '6', '4', '.', '5', '9', '8', '.', '.'},
	}

	got := readInput(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got \n%v, want \n%v", got, want)
	}
}

func TestGetPartNumbers(t *testing.T) {
	input := Schematic{
		{'4', '6', '7', '.', '.', '1', '1', '4', '.', '.'},
		{'.', '.', '.', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '3', '5', '.', '.', '6', '3', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
		{'6', '1', '7', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '+', '.', '5', '8', '.'},
		{'.', '.', '5', '9', '2', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '7', '5', '5', '.'},
		{'.', '.', '.', '$', '.', '*', '.', '.', '.', '.'},
		{'.', '6', '6', '4', '.', '5', '9', '8', '.', '.'},
	}

	want := PartNumbers{
		{467, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 35, 0, 0, 0, 633, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{617, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 592, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 755, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 664, 0, 0, 0, 598, 0, 0, 0, 0},
	}

	got := getPartNumbers(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumPartNumbers(t *testing.T) {
	input := PartNumbers{
		{467, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 35, 0, 0, 0, 633, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{617, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 592, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 755, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 664, 0, 0, 0, 598, 0, 0, 0, 0},
	}

	want := 4361

	got := sum(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestGetGearRatios(t *testing.T) {
	input := Schematic{
		{'4', '6', '7', '.', '.', '1', '1', '4', '.', '.'},
		{'.', '.', '.', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '3', '5', '.', '.', '6', '3', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
		{'6', '1', '7', '*', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '+', '.', '5', '8', '.'},
		{'.', '.', '5', '9', '2', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '7', '5', '5', '.'},
		{'.', '.', '.', '$', '.', '*', '.', '.', '.', '.'},
		{'.', '6', '6', '4', '.', '5', '9', '8', '.', '.'},
	}

	want := GearRatios{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 16345, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 451490, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	got := getGearRatios(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
