package main

import (
	_ "embed"
	"testing"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed sample1
var sample1 string

//go:embed sample2
var sample2 string

//go:embed sample3
var sample3 string

//go:embed sample4
var sample4 string

//go:embed sample5
var sample5 string

func TestPart1(t *testing.T) {
	want := 253
	if got := part1(new(sh.Lines(input))); got != want {
		t.Errorf("Part 1: got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 815
	if got := part2(new(sh.Lines(input))); got != want {
		t.Errorf("Part 2: got %v, want %v", got, want)
	}
}

func TestMonitoring(t *testing.T) {
	type expected struct {
		max  int
		best point
	}

	tests := []struct {
		name string
		args string
		want expected
	}{
		{"Sample 1", sample1, expected{8, point{3, 4}}},
		{"Sample 2", sample2, expected{33, point{5, 8}}},
		{"Sample 3", sample3, expected{35, point{1, 2}}},
		{"Sample 4", sample4, expected{41, point{6, 3}}},
		{"Sample 5", sample5, expected{210, point{11, 13}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asteroids := new(sh.Lines(tt.args))
			if got1, got2 := asteroids.monitoring(); got1 != tt.want.max || got2.x != tt.want.best.x || got2.y != tt.want.best.y {
				t.Errorf("Got max=%v,p=%v, want max=%v,p=%v", got1, got2, tt.want.max, tt.want.best)
			}
		})
	}
}
