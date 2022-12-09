package main

import (
	_ "embed"
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var (
	lines []string

	patterns = [3][3]int{
		{3, 6, 0},
		{0, 3, 6},
		{6, 0, 3},
	}
	patterns2 = [3][3]int{
		{2, 0, 1},
		{0, 1, 2},
		{1, 2, 0},
	}
)

func getShapeScore(shape byte) int {
	return int(shape-'X') + 1
}

func getOutcomeScore(op, pl byte) int {
	return patterns[op-'A'][pl-'X']
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	var score int
	for _, line := range lines {
		op, pl := line[0], line[2]
		score += getShapeScore(pl)
		score += getOutcomeScore(op, pl)
	}
	return score
}

func part2() int {
	var score int
	for _, line := range lines {
		op, st := line[0], line[2]
		score += 3 * int(st-'X')
		pl := patterns2[op-'A'][st-'X']
		score += getShapeScore('X' + byte(pl))
	}
	return score
}

func init() {
	lines = sh.Lines(input)
}
