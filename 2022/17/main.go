package main

import (
	_ "embed"
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
	set "github.com/leonhfr/aoc/shared/setpoint"
)

//go:embed input
var input string

var jets string

type pattern uint16

type direction rune

const (
	WIDTH = 7

	HLINE pattern = 0b0000_0000_0000_1111
	CROSS pattern = 0b0000_0010_0111_0010
	ANGLE pattern = 0b0000_0100_0100_0111
	VLINE pattern = 0b0001_0001_0001_0001
	BLOCK pattern = 0b0000_0000_0011_0011

	RIGHT direction = '>'
	LEFT  direction = '<'
)

var patterns = [5]pattern{HLINE, CROSS, ANGLE, VLINE, BLOCK}

type rock struct {
	shape pattern
	point set.Point
}

func (r rock) collides(v set.Point, chamber *set.SetPoint) bool {
	if r.point.Y+v.Y < 0 {
		return true
	}

	for i := 0; i < 16; i++ {
		if (1<<i)&r.shape == 0 {
			continue
		}

		next := set.Point{X: r.point.X + v.X + i%4, Y: r.point.Y + v.Y + i/4}
		if next.X < 0 || next.X >= WIDTH {
			return true
		}

		if chamber.Has(next) {
			return true
		}
	}

	return false
}

func (r rock) addTo(chamber *set.SetPoint) {
	for i := 0; i < 16; i++ {
		if (1<<i)&r.shape > 0 {
			chamber.Add(set.Point{X: r.point.X + i%4, Y: r.point.Y + i/4})
		}
	}
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	chamber := set.New()
	solve(chamber, 2022, jets)
	_, top := chamber.Boundaries()
	return top.Y + 1
}

func part2() int {
	return 0
}

func solve(chamber *set.SetPoint, height int, input string) {
ROCK:
	for r, i := 0, 0; r < height; r, i = r+1, i+1 {
		_, top := chamber.Boundaries()
		falling := rock{
			shape: patterns[r%len(patterns)],
			point: set.Point{X: 2, Y: top.Y + 3},
		}
		if r > 0 {
			falling.point.Y++
		}

		for ; ; i++ {
			var v set.Point
			switch d := direction(input[i%len(input)]); d {
			case RIGHT:
				v.X = 1
			case LEFT:
				v.X = -1
			}

			if !falling.collides(v, chamber) {
				falling.point = falling.point.Add(v)
			}

			v = set.Point{0, -1}
			if !falling.collides(v, chamber) {
				falling.point = falling.point.Add(v)
			} else {
				falling.addTo(chamber)
				continue ROCK
			}
		}
	}
}

func init() {
	jets = sh.Lines(input)[0]
}

var test = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
