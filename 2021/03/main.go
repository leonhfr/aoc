package main

import (
	_ "embed"
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var binaries []string

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func init() {
	binaries = sh.Lines(input)
}

func part1() int {
	var e, g string
	for i := 0; i < len(binaries[0]); i = i + 1 {
		most, least := mostAndLeastCommonRuneAt(binaries, i)
		g += string(most)
		e += string(least)
	}
	return sh.BinToInt(e) * sh.BinToInt(g)
}

func part2() int {
	o2 := filter(binaries, 0, true)
	co2 := filter(binaries, 0, false)
	return sh.BinToInt(o2) * sh.BinToInt(co2)
}

func mostAndLeastCommonRuneAt(binaries []string, at int) (rune, rune) {
	var c0, c1 int
	for _, binary := range binaries {
		switch binary[at] {
		case '0':
			c0++
		case '1':
			c1++
		}
	}

	if c0 > c1 {
		return '0', '1'
	}

	return '1', '0'
}

func filter(binaries []string, at int, mostCommon bool) string {
	if len(binaries) == 1 {
		return binaries[0]
	}

	most, least := mostAndLeastCommonRuneAt(binaries, at)
	comparator := least
	if mostCommon {
		comparator = most
	}

	filtered := make([]string, 0)
	for _, binary := range binaries {
		if binary[at] == byte(comparator) {
			filtered = append(filtered, binary)
		}
	}

	return filter(filtered, at+1, mostCommon)
}
