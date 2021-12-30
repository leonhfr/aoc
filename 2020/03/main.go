package main

import (
	_ "embed"
	"fmt"

	sp "github.com/leonhfr/aoc/shared/setpoint"
)

//go:embed input
var input string

var forest *sp.SetPoint

type slope struct{ right, down int }

var slopes = []slope{
	{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2},
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	return count(slopes[1])
}

func part2() int {
	result := 1
	for _, s := range slopes {
		result *= count(s)
	}
	return result
}

func count(s slope) (c int) {
	_, pmax := forest.Boundaries()
	for x, y := 0, 0; y <= pmax.Y; x, y = (x+s.right)%(pmax.X+1), y+s.down {
		if forest.Has(sp.Point{X: x, Y: y}) {
			c++
		}
	}
	return
}

func init() {
	forest = sp.ParseSharp(input)
}
