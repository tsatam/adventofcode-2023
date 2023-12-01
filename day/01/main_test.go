package main

import (
	"fmt"
	"testing"
)

func TestHandle(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	want := 142
	got := handle(input)

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
