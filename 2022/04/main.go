package main

import (
	_ "embed"
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string
var pairs [][4]int

var lines []string

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	var n int
	for _, pair := range pairs {
		if contains(pair) {
			n++
		}
	}
	return n
}

func part2() int {
	var n int
	for _, pair := range pairs {
		if overlap(pair) {
			n++
		}
	}
	return n
}

func contains(p [4]int) bool {
	a := p[0] <= p[2] && p[3] <= p[1]
	b := p[2] <= p[0] && p[1] <= p[3]
	return a || b
}

func overlap(p [4]int) bool {
	return p[0] <= p[3] && p[2] <= p[1]
}

func init() {
	lines = sh.Lines(input)
	for _, line := range lines {
		s := strings.FieldsFunc(line, func(r rune) bool { return r == ',' || r == '-' })
		pair := [4]int{sh.ToInt(s[0]), sh.ToInt(s[1]), sh.ToInt(s[2]), sh.ToInt(s[3])}
		pairs = append(pairs, pair)
	}
}
