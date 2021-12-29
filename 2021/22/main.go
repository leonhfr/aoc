package main

import (
	_ "embed"
	"fmt"
	"regexp"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var instructions1 []instruction
var instructions2 []instruction

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	return execute(instructions1).volume()
}

func part2() int {
	return execute(instructions2).volume()
}

func execute(instructions []instruction) reactor {
	var r reactor
	for _, i := range instructions {
		var next reactor
		for _, c := range r {
			intersection := c.c.intersection(i.c)
			if intersection.volume() > 0 {
				next = append(next, instruction{!c.on, intersection})
			}
		}
		if i.on {
			r = append(r, i)
		}
		r = append(r, next...)
	}
	return r
}

type cuboid struct {
	xmin, ymin, zmin int
	xmax, ymax, zmax int
}

func (c cuboid) volume() int {
	return sh.Abs(c.xmax-c.xmin) * sh.Abs(c.ymax-c.ymin) * sh.Abs(c.zmax-c.zmin)
}

func (a cuboid) intersection(b cuboid) cuboid {
	xmin, xmax := sh.Max(a.xmin, b.xmin), sh.Min(a.xmax, b.xmax)
	ymin, ymax := sh.Max(a.ymin, b.ymin), sh.Min(a.ymax, b.ymax)
	zmin, zmax := sh.Max(a.zmin, b.zmin), sh.Min(a.zmax, b.zmax)
	if xmax <= xmin || ymax <= ymin || zmax <= zmin {
		return cuboid{0, 0, 0, 0, 0, 0}
	}
	return cuboid{xmin, ymin, zmin, xmax, ymax, zmax}
}

type instruction struct {
	on bool
	c  cuboid
}

func new(line string) instruction {
	regex := regexp.MustCompile(`^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)
	if !regex.MatchString(line) {
		panic("no match")
	}
	fields := regex.FindStringSubmatch(line)
	return instruction{
		fields[1] == "on",
		cuboid{
			sh.ToInt(fields[2]),
			sh.ToInt(fields[4]),
			sh.ToInt(fields[6]),
			sh.ToInt(fields[3]) + 1,
			sh.ToInt(fields[5]) + 1,
			sh.ToInt(fields[7]) + 1,
		}}
}

type reactor []instruction

func (r reactor) volume() (v int) {
	for _, i := range r {
		if i.on {
			v += i.c.volume()
		} else {
			v -= i.c.volume()
		}
	}
	return
}

func filter(c cuboid) bool {
	x := -50 <= c.xmin && c.xmax <= 51
	y := -50 <= c.ymin && c.ymax <= 51
	z := -50 <= c.zmin && c.zmax <= 51
	return x && y && z
}

func init() {
	lines := sh.Lines(input)
	for _, line := range lines {
		instruction := new(line)
		if filter(instruction.c) {
			instructions1 = append(instructions1, instruction)
		}
		instructions2 = append(instructions2, instruction)
	}
}
