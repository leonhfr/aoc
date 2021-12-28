package main

import (
	_ "embed"
	"fmt"
	"sort"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var ints []int

const target = 2020

func main() {
	fmt.Printf("Part 1: %v\n", part1(ints))
	fmt.Printf("Part 2: %v\n", part2(ints))
}

func init() {
	lines := sh.Lines(input)
	ints = sh.ToInts(lines)
	sort.Ints(ints)
}

func part1(ints []int) int {
	a, b := twoSum(ints, target)
	return a * b
}

func part2(ints []int) int {
	a, b, c := threeSum(ints, target)
	return a * b * c
}

func twoSum(ints []int, sum int) (int, int) {
	for l, r := 0, len(ints)-1; l < r; {
		switch vl, vr := ints[l], ints[r]; {
		case vl+vr == sum:
			return vl, vr
		case vl+vr < sum:
			l++
		case vl+vr > sum:
			r--
		}
	}
	return 0, 0
}

func threeSum(ints []int, sum int) (int, int, int) {
	for i := 0; i < len(ints)-2; i++ {
		a := ints[i]
		b, c := twoSum(ints[i+1:], sum-a)
		if b != 0 && c != 0 {
			return a, b, c
		}
	}
	return 0, 0, 0
}
