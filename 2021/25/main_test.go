package main

import (
	_ "embed"
	"testing"
)

func TestPart1(t *testing.T) {
	want := 568
	if got := part1(); got != want {
		t.Errorf("Part 1: got %v, want %v", got, want)
	}
}
