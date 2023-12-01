package main

import (
	"fmt"
	"testing"
)

func TestHandlePart1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

	want := 142
	got := handlePart1(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestHandlePart2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

	want := 281
	got := handlePart2(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCalibrationValueForLine(t *testing.T) {
	for _, tt := range []struct {
		line string
		want int
	}{
		{line: "", want: 0},
		{line: "1abc2", want: 12},
		{line: "pqr3stu8vwx", want: 38},
		{line: "a1b2c3d4e5f", want: 15},
		{line: "treb7uchet", want: 77},
	} {
		t.Run(fmt.Sprintf("%s -> %d", tt.line, tt.want), func(t *testing.T) {
			got := calibrationValueForLine(tt.line)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCalibrationValueForLinePart2(t *testing.T) {
	for _, tt := range []struct {
		line string
		want int
	}{
		{line: "", want: 0},
		{line: "two1nine", want: 29},
		{line: "eightwothree", want: 83},
		{line: "abcone2threexyz", want: 13},
		{line: "xtwone3four", want: 24},
		{line: "4nineeightseven2", want: 42},
		{line: "zoneight234", want: 14},
		{line: "7pqrstsixteen", want: 76},
	} {
		t.Run(fmt.Sprintf("%s -> %d", tt.line, tt.want), func(t *testing.T) {
			got := calibrationValueForLinePart2(tt.line)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
