package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"github.com/tsatam/adventofcode-2023/common/cartesian"
	"github.com/tsatam/adventofcode-2023/common/fp"
)

var (
	//go:embed input
	input string
)

func main() {
	fmt.Printf("Part 1: %d\n", handlePart1(input))
}

func handlePart1(input string) int {
	pipemap := readInput(input)
	distances := findDistances(pipemap)

	return fp.Reduce(distances, 0, func(curr int, next []int) int {
		nextMax := slices.Max(next)
		if nextMax > curr {
			return nextMax
		}
		return curr
	})
}

func readInput(input string) [][]rune {
	rows := strings.Split(strings.TrimSpace(input), "\n")
	return fp.Map(rows, func(row string) []rune { return []rune(row) })
}

func findDistances(pipemap [][]rune) [][]int {
	distances := make([][]int, len(pipemap))

	start := cartesian.Point{X: -1, Y: -1}

	for i, row := range pipemap {
		distances[i] = make([]int, len(row))
		for j, c := range row {
			distances[i][j] = -1
			if c == 'S' {
				start = cartesian.Point{X: j, Y: i}
				distances[i][j] = 0
			}
		}
	}

	type QueueEntry struct {
		p cartesian.Point
		d int
	}

	queue := []QueueEntry{
		{p: start, d: 0},
	}

	isInBounds := func(p cartesian.Point) bool {
		return p.X >= 0 && p.Y >= 0 && p.Y < len(pipemap) && p.X < len(pipemap[p.Y])
	}

	for len(queue) > 0 {
		self := queue[0]
		queue = queue[1:]

		pipe := pipemap[self.p.Y][self.p.X]

		for _, dir := range []cartesian.Direction{cartesian.Up, cartesian.Down, cartesian.Left, cartesian.Right} {
			if !canTraverse(pipe, dir) {
				continue
			}
			next := self.p.Move(dir)
			if !isInBounds(next) {
				continue
			}
			if !canTraverse(pipemap[next.Y][next.X], reverse(dir)) {
				continue
			}

			currMinDistance := distances[next.Y][next.X]
			if currMinDistance == -1 || currMinDistance > self.d+1 {
				distances[next.Y][next.X] = self.d + 1
				queue = append(queue, QueueEntry{p: next, d: self.d + 1})
			}
		}
	}

	return distances
}

func reverse(d cartesian.Direction) cartesian.Direction {
	switch d {
	case cartesian.Up:
		return cartesian.Down
	case cartesian.Left:
		return cartesian.Right
	case cartesian.Down:
		return cartesian.Up
	case cartesian.Right:
		return cartesian.Left
	}
	panic("should not have reached here")
}

func canTraverse(pipe rune, d cartesian.Direction) bool {
	switch pipe {
	case 'S':
		return true
	case '|':
		return d == cartesian.Up || d == cartesian.Down
	case '-':
		return d == cartesian.Left || d == cartesian.Right
	case 'L':
		return d == cartesian.Up || d == cartesian.Right
	case 'J':
		return d == cartesian.Up || d == cartesian.Left
	case '7':
		return d == cartesian.Down || d == cartesian.Left
	case 'F':
		return d == cartesian.Down || d == cartesian.Right
	default:
		return false
	}
}
