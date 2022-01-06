package main

import (
	_ "embed"
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
	sp "github.com/leonhfr/aoc/shared/setpoint"
)

//go:embed input
var input string

var instructions []instruction

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	var s state
	for _, i := range instructions {
		s.move1(i)
	}
	return sp.ManhattanDistance(sp.Point{}, s.pos)
}

func part2() int {
	s := state{sp.Point{X: 0, Y: 0}, sp.Point{X: -10, Y: 1}, east}
	for _, i := range instructions {
		s.move2(i)
	}
	return sp.ManhattanDistance(sp.Point{}, s.pos)
}

type instruction struct {
	action rune
	value  int
}

func new(str string) instruction {
	return instruction{
		rune(str[0]),
		sh.ToInt(str[1:]),
	}
}

type orientation int

const (
	east orientation = iota
	north
	west
	south
)

type state struct {
	pos, wp sp.Point
	facing  orientation
}

func (s *state) move1(i instruction) {
	switch i.action {
	case 'N':
		s.pos.Y += i.value
	case 'S':
		s.pos.Y -= i.value
	case 'W':
		s.pos.X += i.value
	case 'E':
		s.pos.X -= i.value
	case 'L':
		s.facing = (s.facing + orientation(i.value/90)) % 4
	case 'R':
		s.facing = (s.facing - orientation(i.value/90) + 4) % 4
	case 'F':
		switch s.facing {
		case east:
			s.move1(instruction{'E', i.value})
		case north:
			s.move1(instruction{'N', i.value})
		case south:
			s.move1(instruction{'S', i.value})
		case west:
			s.move1(instruction{'W', i.value})
		}
	}
}

func (s *state) move2(i instruction) {
	switch i.action {
	case 'N':
		s.wp.Y += i.value
	case 'S':
		s.wp.Y -= i.value
	case 'W':
		s.wp.X += i.value
	case 'E':
		s.wp.X -= i.value
	case 'L':
		for k := 0; k < 4-i.value/90; k++ {
			s.wp = rotate90(s.wp)
		}
	case 'R':
		for k := 0; k < i.value/90; k++ {
			s.wp = rotate90(s.wp)
		}
	case 'F':
		s.pos.X += s.wp.X * i.value
		s.pos.Y += s.wp.Y * i.value
	}
}

func rotate90(p sp.Point) sp.Point {
	return sp.Point{
		X: -p.Y,
		Y: p.X,
	}
}

func init() {
	for _, line := range sh.Lines(input) {
		instructions = append(instructions, new(line))
	}
}
