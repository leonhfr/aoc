package main

import (
	_ "embed"
	"testing"
)

func TestPart1(t *testing.T) {
	want := 3369286
	if got := part1(masses); got != want {
		t.Errorf("Part 1: got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 5051054
	if got := part2(masses); got != want {
		t.Errorf("Part 2: got %v, want %v", got, want)
	}
}
