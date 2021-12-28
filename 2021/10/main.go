package main

import (
	_ "embed"
	"fmt"
	"sort"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var dict = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var corruptedPoints = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var incompletePoints = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var lines []string

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() (score int) {
	for _, l := range lines {
		_, r := parseLine(l)
		if r != 0 {
			score += corruptedPoints[r]
		}
	}
	return
}

func part2() int {
	var scores []int
	for _, l := range lines {
		stack, r := parseLine(l)
		if r == 0 {
			scores = append(scores, scoreStack(stack))
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func parseLine(line string) (stack []rune, illegal rune) {
	for _, r := range line {
		if cl, ok := dict[r]; ok {
			stack = append(stack, cl)
		} else {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if r != last {
				return stack, r
			}
		}
	}
	return stack, 0
}

func scoreStack(stack []rune) (score int) {
	for len(stack) > 0 {
		score = 5*score + incompletePoints[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
	}
	return
}

func init() {
	lines = sh.Lines(input)
}
