package main

import (
	_ "embed"
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	max, _ := solve(input)
	return max
}

func part2() int {
	_, min := solve(input)
	return min
}

// From: https://www.reddit.com/r/adventofcode/comments/rnejv5/comment/hpsobsi
func solve(input string) (int, int) {
	blocks := strings.Split(strings.TrimRight(input, "\n"), "inp w\n")[1:]
	max, min := make([]int, len(blocks)), make([]int, len(blocks))
	var stack [][]int
	for i, block := range blocks {
		instructions := strings.Split(block, "\n")
		if instructions[3] == "div z 1" {
			// add y <val>
			s := strings.Split(instructions[14], " ")
			addY := sh.ToInt(s[len(s)-1])
			stack = append(stack, []int{i, addY})
		} else if instructions[3] == "div z 26" {
			addY := stack[len(stack)-1]
			j, x := addY[0], addY[1]
			stack = stack[:len(stack)-1]
			// add x <val>
			s := strings.Split(instructions[4], " ")
			addX := sh.ToInt(s[len(s)-1])
			diff := x + addX
			if diff < 0 {
				i, j, diff = j, i, -diff
			}
			max[i] = 9
			max[j] = 9 - diff
			min[i] = 1 + diff
			min[j] = 1
		}
	}
	return sh.ToInt(ints(max)), sh.ToInt(ints(min))
}

func ints(input []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(input)), ""), "[]")
}
