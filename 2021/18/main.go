package main

import (
	_ "embed"
	"fmt"
	"sort"

	sf "github.com/leonhfr/aoc/2021/18/snailfish"
	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var numbers []string

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	return sum(numbers...).Magnitude()
}

func part2() int {
	var magnitudes []int
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j {
				a := sf.New(numbers[i])
				b := sf.New(numbers[j])
				mag := sf.Add(a, b).Magnitude()
				magnitudes = append(magnitudes, mag)
			}
		}
	}
	sort.Ints(magnitudes)
	return magnitudes[len(magnitudes)-1]
}

func sum(numbers ...string) *sf.SnailFish {
	var c *sf.SnailFish
	for _, n := range numbers {
		s := sf.New(n)
		if c != nil {
			c = sf.Add(c, s)
		} else {
			c = s
		}
	}
	return c
}

func init() {
	numbers = sh.Lines(input)
}
