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

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	rocks := getRocks(input)
	_, b := rocks.Boundaries()
	yMax := b.Y
	var count int

	x, y := 500, 0
	for {
		if rocks.Has(set.Point{X: x, Y: y}) {
			// restart sand at origin
			x, y = 500, 0
		}

		if y > yMax {
			return count
		}

		switch {
		case !rocks.Has(set.Point{X: x, Y: y + 1}) && y <= yMax:
			// drop down
			y += 1
		case !rocks.Has(set.Point{X: x - 1, Y: y + 1}) && y <= yMax:
			// drop left down
			x -= 1
			y += 1
		case !rocks.Has(set.Point{X: x + 1, Y: y + 1}) && y <= yMax:
			// drop right down
			x += 1
			y += 1
		default:
			count++
			rocks.Add(set.Point{X: x, Y: y})
		}
	}
}

func part2() int {
	rocks := getRocks(input)
	_, b := rocks.Boundaries()
	yMax := b.Y

	for y := 0; y <= yMax; y++ {
		for x := 500 - yMax; x <= 500+yMax; x++ {
			if rocks.Has(set.Point{X: x - 1, Y: y}) && rocks.Has(set.Point{X: x, Y: y}) && rocks.Has(set.Point{X: x + 1, Y: y}) {
				rocks.Add(set.Point{X: x, Y: y + 1})
			}
		}
	}

	side := yMax + 2
	return side*side - rocks.Len()
}

func getRocks(input string) *set.SetPoint {
	rocks := set.New()
	for _, line := range sh.Lines(input) {
		points := strings.Split(line, " -> ")

		for i := 1; i < len(points); i++ {
			a := strings.Split(points[i-1], ",")
			b := strings.Split(points[i], ",")

			x1, y1 := sh.ToInt(a[0]), sh.ToInt(a[1])
			x2, y2 := sh.ToInt(b[0]), sh.ToInt(b[1])

			for x := sh.Min(x1, x2); x <= sh.Max(x1, x2); x++ {
				for y := sh.Min(y1, y2); y <= sh.Max(y1, y2); y++ {
					rocks.Add(set.Point{X: x, Y: y})
				}
			}
		}
	}
	return rocks
}
