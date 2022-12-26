package main

import (
	_ "embed"
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
	set "github.com/leonhfr/aoc/shared/setpoint"
)

//go:embed input
var input string

type direction uint8

const (
	N direction = iota
	E
	S
	W
)

var (
	vectors = [4]set.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	origin  = set.Point{0, -1}
)

type blizzard struct {
	origin set.Point
	dir    direction
}

var lines []string

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	blizzards, width, height := parse(input)
	dest := set.Point{width - 1, height}
	return solve(blizzards, origin, dest, width, height, 0)
}

func part2() int {
	blizzards, width, height := parse(input)
	dest := set.Point{width - 1, height}
	total := solve(blizzards, origin, dest, width, height, 0)
	total = solve(blizzards, dest, origin, width, height, total)
	total = solve(blizzards, origin, dest, width, height, total)
	return total
}

func solve(blizzards []blizzard, origin, dest set.Point, width, height, time int) int {
	positions := set.New()
	positions.Add(origin)

	for t := time; ; t++ {
		forbidden := snapshot(blizzards, width, height, t)
		positions = moves(positions, forbidden, origin, dest, width, height)
		if positions.Has(dest) {
			return t
		}
	}
}

func moves(positions, forbidden *set.SetPoint, origin, dest set.Point, width, height int) *set.SetPoint {
	next := set.New()

	for _, pos := range positions.Points() {
		for _, v := range []set.Point{{}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			n := set.Point{pos.X + v.X, pos.Y + v.Y}
			inside := 0 <= n.X && n.X < width && 0 <= n.Y && n.Y < height
			if (inside || n == origin || n == dest) && !forbidden.Has(n) {
				next.Add(n)
			}
		}
	}
	return next
}

func snapshot(blizzards []blizzard, width, height, time int) *set.SetPoint {
	pos := set.New()
	for _, blizzard := range blizzards {
		v := vectors[blizzard.dir]
		pos.Add(set.Point{
			(blizzard.origin.X + time*(v.X+width)) % width,
			(blizzard.origin.Y + time*(v.Y+height)) % height,
		})
	}
	return pos
}

// blizzards, width, height
func parse(input string) ([]blizzard, int, int) {
	var blizzards []blizzard
	var width, height int
	for y, line := range sh.Lines(input) {
		if y == 0 || strings.HasPrefix(line, "##") {
			continue
		}

		width = len(line) - 2
		height++

		for x, r := range line {
			switch r {
			case '^':
				blizzards = append(blizzards, blizzard{set.Point{x - 1, y - 1}, N})
			case '>':
				blizzards = append(blizzards, blizzard{set.Point{x - 1, y - 1}, E})
			case 'v':
				blizzards = append(blizzards, blizzard{set.Point{x - 1, y - 1}, S})
			case '<':
				blizzards = append(blizzards, blizzard{set.Point{x - 1, y - 1}, W})
			}
		}
	}
	return blizzards, width, height
}

var test = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`
