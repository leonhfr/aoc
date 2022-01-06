package main

import (
	_ "embed"
	"fmt"
	"testing"

	sh "github.com/leonhfr/aoc/shared"
	sp "github.com/leonhfr/aoc/shared/setpoint"
)

//go:embed sample
var sample string

func TestPart1(t *testing.T) {
	want := 5425
	if got := part1(); got != want {
		t.Errorf("Part 1: got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 14052
	if got := part2(); got != want {
		t.Errorf("Part 2: got %v, want %v", got, want)
	}
}

func Test_nextImage(t *testing.T) {
	lines := sh.Lines(sample)
	algorithm = lines[0]

	tests := []struct {
		img  string
		out  bool
		want string
	}{
		{
			"#..#.\n#....\n##..#\n..#..\n..###", false,
			".##.##.\n#..#.#.\n##.#..#\n####..#\n.#..##.\n..##..#\n...#.#.",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			img := sp.ParseSharp(tt.img)
			if got := nextImage(algorithm, img, false).String(); got != tt.want {
				t.Errorf("nextImage(_, %v, %v) = %v, want %v", tt.img, tt.out, got, tt.want)
			}
		})
	}
}

func Test_addBorder(t *testing.T) {
	tests := []struct {
		img  string
		out  bool
		want string
	}{
		{"#.\n.#", false, "#.\n.#"},
		{"#.\n.#", true, "####\n##.#\n#.##\n####"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			img := sp.ParseSharp(tt.img)
			addBorder(img, tt.out)
			if got := img.String(); got != tt.want {
				t.Errorf("addBorder(%v, %v) = %v, want %v", tt.img, tt.out, got, tt.want)
			}
		})
	}
}

func Test_enhancement(t *testing.T) {
	lines := sh.Lines(sample)
	algorithm = lines[0]

	tests := []struct {
		img       string
		p         sp.Point
		out, want bool
	}{
		{"...\n#..\n.#.", sp.Point{X: 1, Y: 1}, false, true},
		{"...\n#..\n.#.", sp.Point{X: 1, Y: 1}, true, true},
		{"...\n#..\n.#.", sp.Point{X: 0, Y: 1}, false, true},
		{"...\n#..\n.#.", sp.Point{X: 0, Y: 1}, true, true},
		{"...\n#..\n...", sp.Point{X: 0, Y: 1}, true, false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			img := sp.ParseSharp(tt.img)
			if got := enhancement(algorithm, img, tt.p, tt.out); got != tt.want {
				t.Errorf("enhancement(_, %v, %v, %v) = %v, want %v", tt.img, tt.p, tt.out, got, tt.want)
			}
		})
	}
}
