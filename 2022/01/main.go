package main

import (
	_ "embed"
	"fmt"
	"sort"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var (
	lines    []string
	calories []int
)

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	return sh.Max(calories...)
}

func part2() int {
	sort.Ints(calories)
	l := len(calories)
	return calories[l-1] + calories[l-2] + calories[l-3]
}

func sumCalories() []int {
	calories := []int{0}
	for _, line := range lines {
		if line == "" {
			calories = append(calories, 0)
			continue
		}

		calories[len(calories)-1] += sh.ToInt(line)
	}
	return calories
}

func init() {
	lines = sh.Lines(input)
	calories = sumCalories()
}
