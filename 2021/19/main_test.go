package main

import (
	_ "embed"
	"testing"
)

func TestPart1(t *testing.T) {
	want := 318
	if got, _ := part1(); got != want {
		t.Errorf("Part 1: got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 12166
	_, scanners := locations()
	if got := part2(scanners); got != want {
		t.Errorf("Part 2: got %v, want %v", got, want)
	}
}
