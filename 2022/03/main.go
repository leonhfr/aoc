package main

import (
	_ "embed"
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var lines []string

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	var sum int
	for _, line := range lines {
		m := makeSet(line[:len(line)/2])
		for _, r := range line[len(line)/2:] {
			if _, ok := m[r]; ok {
				sum += getPriority(r)
				break
			}
		}
	}
	return sum
}

func part2() int {
	var sum int
	for i := 0; i < len(lines); i += 3 {
		a := makeSet(lines[i])
		b := makeSet(lines[i+1])
		c := makeSet(lines[i+2])
		d := makeSet(lines[i] + lines[i+1] + lines[i+2])
		for r := range d {
			_, okA := a[r]
			_, okB := b[r]
			_, okC := c[r]
			if okA && okB && okC {
				sum += getPriority(r)
				break
			}
		}
	}
	return sum
}

func init() {
	lines = sh.Lines(input)
}

func makeSet(s string) map[rune]struct{} {
	m := make(map[rune]struct{})
	for _, r := range s {
		m[r] = struct{}{}
	}
	return m
}

func getPriority(r rune) int {
	if 'a' <= r && r <= 'z' {
		return int(r - 'a' + 1)
	}
	return int(r - 'A' + 27)
}
