package main

import (
	_ "embed"
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var vents []vent

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	var pts points
	for _, vent := range vents {
		if vent.p1.x == vent.p2.x || vent.p1.y == vent.p2.y {
			pts = append(pts, vent.toPoints()...)
		}
	}
	return pts.overlaps()
}

func part2() int {
	var pts points
	for _, vent := range vents {
		pts = append(pts, vent.toPoints()...)

	}
	return pts.overlaps()
}

type point struct {
	x, y int
}

type points []point

type set map[point]int

func (pts points) overlaps() (count int) {
	s := make(set)
	for _, p := range pts {
		s[p]++
		if s[p] == 2 {
			count++
		}
	}
	return
}

type vent struct {
	p1, p2 point
}

func (v vent) toPoints() (p points) {
	i, j := v.unitVec()
	for t := 0; t < v.nPoints(); t++ {
		p = append(p, point{v.p1.x + t*i, v.p1.y + t*j})
	}
	return
}

func (v vent) unitVec() (i, j int) {
	if v.p1.x != v.p2.x {
		i = (v.p2.x - v.p1.x) / sh.Abs(v.p2.x-v.p1.x)
	}
	if v.p1.y != v.p2.y {
		j = (v.p2.y - v.p1.y) / sh.Abs(v.p2.y-v.p1.y)
	}
	return
}

func (v vent) nPoints() int {
	return sh.Max(sh.Abs(v.p2.x-v.p1.x), sh.Abs(v.p2.y-v.p1.y)) + 1
}

func init() {
	lines := sh.Lines(input)
	for _, line := range lines {
		points := strings.Split(line, " -> ")
		p1 := sh.IntList(points[0])
		p2 := sh.IntList(points[1])
		vents = append(vents, vent{point{p1[0], p1[1]}, point{p2[0], p2[1]}})
	}
}
