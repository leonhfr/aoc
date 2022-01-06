package main

import (
	_ "embed"
	"fmt"

	mat "github.com/leonhfr/aoc/shared/matrix"
)

//go:embed input
var input string

var octopuses mat.Matrix

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() (flashes int) {
	dup := octopuses.Duplicate()
	for i := 0; i < 100; i++ {
		flashes += step(dup)
	}
	return
}

func part2() (s int) {
	dup := octopuses.Duplicate()
	for i := 1; ; i++ {
		step(dup)
		if synchronized(dup) {
			return i
		}
	}
}

func init() {
	octopuses = mat.IntMatrix(input)
}

func step(m mat.Matrix) (flashes int) {
	increment(m)
	for i := flash(m); i > 0; i = flash(m) {
		flashes += i
	}
	reset(m)
	return
}

func increment(m mat.Matrix) {
	for _, c := range m.Coordinates() {
		m.Inc(c.I, c.J, 1)
	}
}

func flash(m mat.Matrix) (flashes int) {
	for _, c := range m.Coordinates() {
		if m.Get(c.I, c.J) > 9 {
			spread(m, c.I, c.J)
			m.Set(c.I, c.J, -1)
			flashes++
		}
	}
	return
}

func spread(m mat.Matrix, i, j int) {
	vectors := []struct{ i, j int }{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}

	for _, v := range vectors {
		vi, vj := i+v.i, j+v.j
		if m.Inside(vi, vj) && m.Get(vi, vj) >= 0 {
			m.Inc(vi, vj, 1)
		}
	}
}

func reset(m mat.Matrix) {
	for _, c := range m.Coordinates() {
		if m.Get(c.I, c.J) == -1 {
			m.Set(c.I, c.J, 0)
		}
	}
}

func synchronized(m mat.Matrix) bool {
	for _, c := range m.Coordinates() {
		if m.Get(c.I, c.J) != 0 {
			return false
		}
	}
	return true
}
