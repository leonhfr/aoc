package main

import (
	_ "embed"
	"testing"
)

func TestPart1(t *testing.T) {
	want := "2-0-0=1-0=2====20=-2"
	if got := part1(); got != want {
		t.Errorf("Part 1: got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 0
	if got := part2(); got != want {
		t.Errorf("Part 2: got %v, want %v", got, want)
	}
}
