package main

import (
	_ "embed"
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var masses []int

func main() {
	fmt.Printf("Part 1: %v\n", part1(masses))
	fmt.Printf("Part 2: %v\n", part2(masses))
}

func init() {
	lines := sh.Lines(input)
	masses = sh.ToInts(lines)
}

func part1(masses []int) int {
	sum := 0
	for _, m := range masses {
		sum += fuel(m)
	}
	return sum
}

func part2(masses []int) int {
	sum := 0
	for _, mass := range masses {
		for m := fuel(mass); m > 0; m = fuel(m) {
			sum += m
		}
	}
	return sum
}

func fuel(m int) int {
	return m/3 - 2
}
