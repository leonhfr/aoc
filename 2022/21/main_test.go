package main

import (
	_ "embed"
	"testing"
)

func TestPart1(t *testing.T) {
	want := 56490240862410
	if got := part1(); got != want {
		t.Errorf("Part 1: got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 3403989691757
	if got := part2(); got != want {
		t.Errorf("Part 2: got %v, want %v", got, want)
	}
}
