package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var asteroids amap

func main() {
	fmt.Printf("Part 1: %v\n", part1(asteroids))
	fmt.Printf("Part 2: %v\n", part2(asteroids))
}

func init() {
	lines := sh.Lines(input)
	asteroids = new(lines)
}

func part1(a amap) int {
	max, _ := a.monitoring()
	return max
}

func part2(a amap) int {
	_, p := a.monitoring()
	pts := a.vaporization(p)
	return 100*pts[200].x + pts[200].y
}

type point struct {
	x, y int
}

type set map[point]struct{}

type amap set

func new(lines []string) amap {
	a := make(amap)
	for y, line := range lines {
		for x, p := range line {
			if p == '#' {
				a[point{x, y}] = struct{}{}
			}
		}
	}
	return a
}

func (a amap) monitoring() (max int, best point) {
	for p1 := range a {
		s := make(set)
		for p2 := range a {
			if p1.x != p2.x || p1.y != p2.y {
				dx, dy := p2.x-p1.x, p2.y-p1.y
				dx, dy = dx/sh.Abs(sh.Gcd(dx, dy)), dy/sh.Abs(sh.Gcd(dx, dy))
				s[point{dx, dy}] = struct{}{}
			}
		}
		if len(s) > max {
			max, best = len(s), p1
		}
	}
	return
}

func (a amap) vaporization(best point) []point {
	order, done := []point{best}, set{best: struct{}{}}
	for len(order) != len(a) {
		closest := make(map[point]point)
		for p := range a {
			if _, ok := done[p]; !ok {
				dx, dy := p.x-best.x, p.y-best.y
				dx, dy = dx/sh.Abs(sh.Gcd(dx, dy)), dy/sh.Abs(sh.Gcd(dx, dy))
				if c, ok := closest[point{dx, dy}]; !ok || sh.Abs(p.x-best.x)+sh.Abs(p.y-best.y) < sh.Abs(c.x-best.x)+sh.Abs(c.y-best.y) {
					closest[point{dx, dy}] = p
				}
			}
		}
		var sli []point
		for _, p := range closest {
			sli = append(sli, p)
			done[p] = struct{}{}
		}
		sort.Slice(sli, func(i, j int) bool {
			a := -math.Atan2(float64(sli[i].x-best.x), float64(sli[i].y-best.y))
			b := -math.Atan2(float64(sli[j].x-best.x), float64(sli[j].y-best.y))
			return a < b
		})
		order = append(order, sli...)
	}
	return order
}
