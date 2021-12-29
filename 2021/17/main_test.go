package main

import (
	_ "embed"
	"testing"
)

func TestPart1(t *testing.T) {
	want := 7503
	if got, _, _ := part1(); got != want {
		t.Errorf("Part 1: got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 3229
	if got := part2(16, 18); got != want {
		t.Errorf("Part 2: got %v, want %v", got, want)
	}
}