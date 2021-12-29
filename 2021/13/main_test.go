package main

import (
	_ "embed"
	"testing"
)

// PGHZBFJC
var code = `###...##..#..#.####.###..####...##..##.
#..#.#..#.#..#....#.#..#.#.......#.#..#
#..#.#....####...#..###..###.....#.#...
###..#.##.#..#..#...#..#.#.......#.#...
#....#..#.#..#.#....#..#.#....#..#.#..#
#.....###.#..#.####.###..#.....##...##.`

func TestPart1(t *testing.T) {
	want := 790
	if got := part1(); got != want {
		t.Errorf("Part 1: got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got := part2(); got != code {
		t.Errorf("Part 2: got %v, want %v", got, code)
	}
}
