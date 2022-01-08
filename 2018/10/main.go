package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
	sp "github.com/leonhfr/aoc/shared/setpoint"
)

//go:embed input
var input string

var data map[vector]vector

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	sec := minBox()
	img := draw(sec)
	fmt.Println(img.String())
	return sec
}

func part2() int {
	return minBox()
}

type vector struct {
	x, y int
}

func minBox() (sec int) {
	for box := math.MaxInt; ; sec++ {
		nmin, nmax := vector{math.MaxInt, math.MaxInt}, vector{math.MinInt, math.MinInt}
		for p, v := range data {
			nmin.x = sh.Min(nmin.x, p.x+sec*v.x)
			nmin.y = sh.Min(nmin.y, p.y+sec*v.y)
			nmax.x = sh.Max(nmax.x, p.x+sec*v.x)
			nmax.y = sh.Max(nmax.y, p.y+sec*v.y)
		}
		nbox := nmax.x + nmax.y - nmin.x - nmin.y
		if nbox > box {
			break
		}
		box = nbox
	}
	return sec - 1
}

func draw(sec int) *sp.SetPoint {
	img := sp.New()
	for p, v := range data {
		img.Add(sp.Point{
			X: p.x + sec*v.x,
			Y: p.y + sec*v.y,
		})
	}
	return img
}

func init() {
	regex := regexp.MustCompile(`^position=<(-?\d+),(-?\d+)>velocity=<(-?\d+),(-?\d+)>$`)
	data = make(map[vector]vector)
	lines := sh.Lines(input)
	for _, line := range lines {
		clean := strings.ReplaceAll(line, " ", "")
		if regex.MatchString(clean) {
			fields := regex.FindStringSubmatch(clean)
			x, y := sh.ToInt(fields[1]), sh.ToInt(fields[2])
			vx, vy := sh.ToInt(fields[3]), sh.ToInt(fields[4])
			data[vector{x, y}] = vector{vx, vy}
		}
	}
}
