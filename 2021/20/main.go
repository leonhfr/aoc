package main

import (
	_ "embed"
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
	sp "github.com/leonhfr/aoc/shared/setpoint"
)

//go:embed input
var input string

var algorithm string
var image *sp.SetPoint

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	for i := 0; i < 2; i++ {
		image = nextImage(algorithm, image, i%2 == 1)
	}
	return image.Len()
}

func part2() int {
	for i := 0; i < 48; i++ {
		image = nextImage(algorithm, image, i%2 == 1)
	}
	return image.Len()
}

func nextImage(algorithm string, img *sp.SetPoint, outside bool) *sp.SetPoint {
	next := sp.New()
	addBorder(img, outside)
	pmin, pmax := img.Boundaries()
	for x := pmin.X - 1; x <= pmax.X+1; x++ {
		for y := pmin.Y - 1; y <= pmax.Y+1; y++ {
			p := sp.Point{X: x, Y: y}
			if enhancement(algorithm, img, p, outside) {
				next.Add(p)
			}
		}
	}
	return next
}

func addBorder(img *sp.SetPoint, value bool) {
	if !value {
		return
	}
	pmin, pmax := img.Boundaries()
	for x := pmin.X - 1; x <= pmax.X+1; x++ {
		img.Add(sp.Point{X: x, Y: pmin.Y - 1})
		img.Add(sp.Point{X: x, Y: pmax.Y + 1})
	}
	for y := pmin.Y - 1; y <= pmax.Y+1; y++ {
		img.Add(sp.Point{X: pmin.X - 1, Y: y})
		img.Add(sp.Point{X: pmax.X + 1, Y: y})
	}
}

func enhancement(algorithm string, img *sp.SetPoint, p sp.Point, outside bool) bool {
	vectors := []struct{ i, j int }{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{0, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	index := 0
	for _, v := range vectors {
		index = index << 1
		n := sp.Point{X: p.X + v.i, Y: p.Y + v.j}
		if img.Has(n) || (outside && !img.Inside(n)) {
			index++
		}
	}

	return algorithm[index] == '#'
}

func init() {
	lines := sh.Lines(input)
	algorithm = lines[0]
	image = sp.ParseSharp(strings.Join(lines[2:], "\n"))
}
