package main

import (
	_ "embed"
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var alignments []int

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	var fuel int
	median := sh.Median(alignments)
	for _, a := range alignments {
		fuel += sh.Abs(a - median)
	}
	return fuel
}

func part2() int {
	var fuel int
	mean := sh.Mean(alignments)
	for _, a := range alignments {
		fuel += sh.TriangularNumber(sh.Abs(a - mean))
	}
	return fuel
}

func init() {
	lines := sh.Lines(input)
	alignments = sh.IntList(lines[0])
}
