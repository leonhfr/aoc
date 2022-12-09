package main

import (
	_ "embed"
	"fmt"

	"github.com/leonhfr/aoc/shared/matrix"
	"github.com/leonhfr/aoc/shared/setpoint"
)

//go:embed input
var input string

var forest matrix.Matrix

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	set := setpoint.New()

	for i := 1; i <= forest.M(); i++ {
		left, right := -1, -1
		for j := 1; j <= forest.N(); j++ {
			if p := forest.Get(i, j); p > left {
				set.Add(setpoint.Point{X: j, Y: i})
				left = p
			}

			if p := forest.Get(i, forest.N()-j+1); p > right {
				set.Add(setpoint.Point{X: forest.N() - j + 1, Y: i})
				right = p
			}
		}
	}

	for j := 1; j <= forest.N(); j++ {
		top, bottom := -1, -1
		for i := 1; i <= forest.M(); i++ {
			if p := forest.Get(i, j); p > top {
				set.Add(setpoint.Point{X: j, Y: i})
				top = p
			}

			if p := forest.Get(forest.M()-i+1, j); p > bottom {
				set.Add(setpoint.Point{X: j, Y: forest.M() - i + 1})
				bottom = p
			}
		}
	}

	return set.Len()
}

func part2() int {
	var score int
	for _, tree := range forest.Coordinates() {
		if tree.I == 1 || tree.I == forest.M() ||
			tree.J == 1 || tree.J == forest.N() {
			continue
		}

		var up, down, left, right int
		p := forest.Get(tree.I, tree.J)

		// up
		for i := 1; i < tree.I; i++ {
			up++
			if t := forest.Get(tree.I-i, tree.J); t >= p {
				break
			}
		}

		// down
		for i := 1; i < forest.M()-tree.I+1; i++ {
			down++
			if t := forest.Get(tree.I+i, tree.J); t >= p {
				break
			}
		}

		// left
		for j := 1; j < tree.J; j++ {
			left++
			if t := forest.Get(tree.I, tree.J-j); t >= p {
				break
			}
		}

		// right
		for j := 1; j < forest.N()-tree.J+1; j++ {
			right++
			if t := forest.Get(tree.I, tree.J+j); t >= p {
				break
			}
		}

		if s := up * down * left * right; score < s {
			score = s
		}
	}
	return score
}

func init() {
	forest = matrix.IntMatrix(input)
}
