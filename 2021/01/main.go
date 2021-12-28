package main

import (
	_ "embed"
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var measurements []int

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func init() {
	lines := sh.Lines(input)
	measurements = sh.ToInts(lines)
}

func part1() int {
	var count int
	for prev, i := measurements[0], 1; i < len(measurements); prev, i = measurements[i], i+1 {
		if prev < measurements[i] {
			count++
		}
	}
	return count
}

func part2() int {
	var count int
	for i := 1; i+2 < len(measurements); i = i + 1 {
		previous := measurements[i-1] + measurements[i] + measurements[i+1]
		current := measurements[i] + measurements[i+1] + measurements[i+2]
		if current > previous {
			count++
		}
	}
	return count
}
