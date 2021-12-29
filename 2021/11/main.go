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
	for i := 0; i < m.M(); i++ {
		for j := 0; j < m.N(); j++ {
			m[i][j]++
		}
	}
}

func flash(m mat.Matrix) (flahes int) {
	for i := 0; i < m.M(); i++ {
		for j := 0; j < m.N(); j++ {
			if m[i][j] > 9 {
				spread(m, i, j)
				m[i][j] = -1
				flahes++
			}
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
		if 0 <= vi && vi < m.M() && 0 <= vj && vj < m.N() {
			if m[vi][vj] >= 0 {
				m[vi][vj]++
			}
		}
	}
}

func reset(m mat.Matrix) {
	for i := 0; i < m.M(); i++ {
		for j := 0; j < m.N(); j++ {
			if m[i][j] == -1 {
				m[i][j] = 0
			}
		}
	}
}

func synchronized(m mat.Matrix) bool {
	for i := 0; i < m.M(); i++ {
		for j := 0; j < m.N(); j++ {
			if m[i][j] != 0 {
				return false
			}
		}
	}
	return true
}
