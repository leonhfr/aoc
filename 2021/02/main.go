package main

import (
	_ "embed"
	"fmt"
	"regexp"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var instructions []instruction

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func init() {
	lines := sh.Lines(input)
	for _, line := range lines {
		instructions = append(instructions, new(line))
	}
}

func part1() int {
	var h, v int
	for _, i := range instructions {
		h += i.h
		v += i.v
	}
	return h * v
}

func part2() int {
	var aim, h, v int
	for _, i := range instructions {
		aim += i.v
		h += i.h
		v += aim * i.h
	}
	return h * v
}

type instruction struct {
	h, v int
}

func new(str string) instruction {
	regex := regexp.MustCompile(`^(down|up|forward) (\d)+$`)
	var i instruction

	if !regex.MatchString(str) {
		return i
	}

	fields := regex.FindStringSubmatch(str)
	delta := sh.ToInt(fields[2])
	switch fields[1] {
	case "down":
		i.v = delta
	case "up":
		i.v = -delta
	case "forward":
		i.h = delta
	}
	return i
}
