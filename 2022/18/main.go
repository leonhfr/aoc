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

type (
	point  struct{ x, y, z int }
	tuple  [2]int
	linMap map[tuple][]int
	bounds struct{ min, max int }
)

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	points := parse(input)
	maps := linearMaps(points)

	var count int
	for _, m := range maps {
		for _, s := range m {
			count += 2

			if len(s) == 1 {
				continue
			}

			sort.Ints(s)
			for i := 1; i < len(s); i++ {
				if s[i-1]+1 != s[i] {
					count += 2
				}
			}
		}
	}
	return count
}

func part2() int {
	points := parse(input)
	maps := linearMaps(points)
	min, max, set := makeSet(points)
	delta := max - min
	funcs := [3]func(v int, t tuple) point{
		func(v int, t tuple) point { return point{v, t[0], t[1]} },
		func(v int, t tuple) point { return point{t[0], v, t[1]} },
		func(v int, t tuple) point { return point{t[0], t[1], v} },
	}

	var count int
	for fi, m := range maps {
		for t, s := range m {
			count += 2

			if len(s) == 1 {
				continue
			}

			sort.Ints(s)
			for i := 1; i < len(s); i++ {
				if s[i-1]+1 != s[i] {
					p := funcs[fi](s[i-1]+1, t)
					if !isTrapped(p, set, delta) {
						count += 2
					}
				}
			}
		}
	}
	return count
}

func makeSet(points []point) (int, int, map[point]struct{}) {
	min, max := math.MaxInt, math.MinInt
	set := make(map[point]struct{})
	for _, p := range points {
		min = sh.Min(min, p.x, p.y, p.z)
		max = sh.Max(max, p.x, p.y, p.z)
		set[p] = struct{}{}
	}
	return min, max, set
}

func isTrapped(p point, set map[point]struct{}, delta int) bool {
	for _, v := range [6]point{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},
	} {
		if !touches(p, v, set, delta) {
			return false
		}
	}
	return true
}

func touches(p, v point, set map[point]struct{}, delta int) bool {
	for d := 1; d <= delta; d++ {
		n := point{p.x + d*v.x, p.y + d*v.y, p.z + d*v.z}
		if _, ok := set[n]; ok {
			return true
		}
	}
	return false
}

func linearMaps(points []point) [3]linMap {
	xm := make(linMap) // [y,z]
	ym := make(linMap) // [x,z]
	zm := make(linMap) // [x,y]
	for _, p := range points {
		xm[tuple{p.y, p.z}] = append(xm[tuple{p.y, p.z}], p.x)
		ym[tuple{p.x, p.z}] = append(ym[tuple{p.x, p.z}], p.y)
		zm[tuple{p.x, p.y}] = append(zm[tuple{p.x, p.y}], p.z)
	}
	return [3]linMap{xm, ym, zm}
}

func parse(input string) []point {
	var points []point
	for _, line := range sh.Lines(input) {
		n := sh.IntList(line)
		points = append(points, point{n[0], n[1], n[2]})
	}
	return points
}
