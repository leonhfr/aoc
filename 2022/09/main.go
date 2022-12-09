package main

import (
	_ "embed"
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

type (
	vector struct{ i, j int }
	move   struct {
		vector
		dist int
	}
	point struct{ x, y int }
)

var moves []move

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	head, tail := point{}, point{}
	visited := map[point]struct{}{
		{}: {},
	}

	for _, move := range moves {
		if move.i != 0 {
			for i := 0; i < move.dist; i++ {
				x, y := head.x, head.y
				head.x += move.i
				if sh.Abs(tail.x-head.x) > 1 {
					tail.x, tail.y = x, y
					visited[tail] = struct{}{}
				}
			}
		}

		if move.j != 0 {
			for j := 0; j < move.dist; j++ {
				x, y := head.x, head.y
				head.y += move.j
				if sh.Abs(tail.y-head.y) > 1 {
					tail.x, tail.y = x, y
					visited[tail] = struct{}{}
				}
			}
		}
	}

	return len(visited)
}

func part2() int {
	rope := [10]point{}
	visited := map[point]struct{}{
		{}: {},
	}

	for _, move := range moves {
		for k := 0; k < move.dist; k++ {
			rope[0].x += move.i
			rope[0].y += move.j
			for p := 1; p <= 9; p++ {
				x1, y1 := rope[p-1].x, rope[p-1].y
				x2, y2 := rope[p].x, rope[p].y
				dx, dy := x1-x2, y1-y2
				switch {
				case dx == 2 && dy == 2:
					rope[p] = point{x1 - 1, y1 - 1}
				case dx == 2 && dy == -2:
					rope[p] = point{x1 - 1, y1 + 1}
				case dx == -2 && dy == 2:
					rope[p] = point{x1 + 1, y1 - 1}
				case dx == -2 && dy == -2:
					rope[p] = point{x1 + 1, y1 + 1}
				case dx == 2:
					rope[p] = point{x1 - 1, y1}
				case dx == -2:
					rope[p] = point{x1 + 1, y1}
				case dy == 2:
					rope[p] = point{x1, y1 - 1}
				case dy == -2:
					rope[p] = point{x1, y1 + 1}
				}
			}
			visited[rope[9]] = struct{}{}
		}
	}

	return len(visited)
}

func init() {
	for _, line := range sh.Lines(input) {
		fields := strings.Fields(line)
		unit := unitVector(fields[0])
		dist := sh.ToInt(fields[1])
		moves = append(moves, move{
			vector: unit,
			dist:   dist,
		})
	}
}

func unitVector(dir string) vector {
	switch dir {
	case "U":
		return vector{0, 1}
	case "R":
		return vector{1, 0}
	case "D":
		return vector{0, -1}
	default:
		return vector{-1, 0}
	}
}
