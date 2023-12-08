package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHandlePart1(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  int
	}{
		{
			input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`,
			want: 2,
		},
		{
			input: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`,
			want: 6,
		},
	} {
		t.Run(fmt.Sprintf("%s -> %d", tt.input, tt.want), func(t *testing.T) {
			got := handlePart1(tt.input)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestHandlePart2(t *testing.T) {
	input := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`
	want := 6
	got := handlePart2(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestReadInput(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  NetworkMap
	}{
		{
			input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`,
			want: NetworkMap{
				instructions: "RL",
				nodes: map[string]Node{
					"AAA": {"BBB", "CCC"},
					"BBB": {"DDD", "EEE"},
					"CCC": {"ZZZ", "GGG"},
					"DDD": {"DDD", "DDD"},
					"EEE": {"EEE", "EEE"},
					"GGG": {"GGG", "GGG"},
					"ZZZ": {"ZZZ", "ZZZ"},
				},
			},
		},
		{
			input: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`,
			want: NetworkMap{
				instructions: "LLR",
				nodes: map[string]Node{
					"AAA": {"BBB", "BBB"},
					"BBB": {"AAA", "ZZZ"},
					"ZZZ": {"ZZZ", "ZZZ"},
				},
			},
		},
	} {
		t.Run(fmt.Sprintf("%s -> %v", tt.input, tt.want), func(t *testing.T) {
			got := readInput(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
