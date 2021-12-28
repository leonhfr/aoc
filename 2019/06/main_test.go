package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 249308
	if got := part1(orbits); got != want {
		t.Errorf("Part 1: got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 349
	if got := part2(orbits); got != want {
		t.Errorf("Part 2: got %v, want %v", got, want)
	}
}
