package main

import (
	_ "embed"
	"fmt"
	"math"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

type vector struct {
	x, y, z, w int
}

type space struct {
	vectors  map[vector]struct{}
	min, max vector
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	cubes := parse()
	for i := 0; i < 6; i++ {
		cubes = cycle(cubes, false)
	}
	return len(cubes.vectors)
}

func part2() int {
	cubes := parse()
	for i := 0; i < 6; i++ {
		cubes = cycle(cubes, true)
	}
	return len(cubes.vectors)
}

func cycle(p *space, fourth bool) *space {
	n := new()
	for x := p.min.x - 1; x <= p.max.x+1; x++ {
		for y := p.min.y - 1; y <= p.max.y+1; y++ {
			for z := p.min.z - 1; z <= p.max.z+1; z++ {
				if fourth {
					for w := p.min.w - 1; w <= p.max.w+1; w++ {
						v := vector{x, y, z, w}
						if neighboors(p, v, []int{-1, 0, 1}) {
							n.vectors[v] = struct{}{}
						}
					}
				} else {
					v := vector{x, y, z, 0}
					if neighboors(p, v, []int{0}) {
						n.vectors[v] = struct{}{}
					}
				}
			}
		}
	}
	n.boundaries()
	return n
}

func neighboors(cubes *space, v vector, wvs []int) bool {
	c, vs := 0, [3]int{-1, 0, 1}
	for _, x := range vs {
		for _, y := range vs {
			for _, z := range vs {
				for _, w := range wvs {
					n := vector{v.x + x, v.y + y, v.z + z, v.w + w}
					if _, ok := cubes.vectors[n]; v != n && ok {
						c++
					}
				}
			}
		}
	}

	if _, ok := cubes.vectors[v]; ok {
		return c == 2 || c == 3
	}
	return c == 3
}

func new() *space {
	return &space{
		vectors: make(map[vector]struct{}),
		min:     vector{math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt},
		max:     vector{math.MinInt, math.MinInt, math.MinInt, math.MinInt},
	}
}

func (cubes *space) boundaries() {
	for v := range cubes.vectors {
		cubes.min.x = sh.Min(cubes.min.x, v.x)
		cubes.min.y = sh.Min(cubes.min.y, v.y)
		cubes.min.z = sh.Min(cubes.min.z, v.z)
		cubes.min.w = sh.Min(cubes.min.w, v.w)
		cubes.max.x = sh.Max(cubes.max.x, v.x)
		cubes.max.y = sh.Max(cubes.max.y, v.y)
		cubes.max.z = sh.Max(cubes.max.z, v.z)
		cubes.max.w = sh.Max(cubes.max.w, v.w)
	}
}

func parse() *space {
	cubes := new()
	for y, line := range sh.Lines(input) {
		for x, r := range line {
			if r == '#' {
				cubes.vectors[vector{x, y, 0, 0}] = struct{}{}
			}
		}
	}
	cubes.boundaries()
	return cubes
}
