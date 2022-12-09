package main

import (
	_ "embed"
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var line string

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	return indexBuffer(line, 4)
}

func part2() int {
	return indexBuffer(line, 14)
}

func indexBuffer(s string, size int) int {
	for i := size; i <= len(s); i++ {
		if uniqueChars(s[i-size : i]) {
			return i
		}
	}
	return 0
}

func uniqueChars(s string) bool {
	var seen uint32
	for _, r := range s {
		var index uint32 = 1 << (r - 'a')
		if seen&index > 0 {
			return false
		}
		seen ^= index
	}
	return true
}

func init() {
	line = sh.Lines(input)[0]
}
