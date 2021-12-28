package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var heightmap sh.Matrix

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	lp := lowestPoints(heightmap)
	return riskLevel(heightmap, lp)
}

func part2() int {
	lp := lowestPoints(heightmap)
	basins := getBasins(heightmap, lp)
	var sizes []int
	for _, b := range basins {
		sizes = append(sizes, len(b))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes[0] * sizes[1] * sizes[2]
}

func lowestPoints(hm sh.Matrix) (coords []sh.Coordinates) {
	for i := 0; i < hm.M(); i++ {
		for j := 0; j < hm.N(); j++ {
			v := hm[i][j]
			up := i == 0 || v < hm.Get(i-1, j)
			down := i == hm.M()-1 || v < hm.Get(i+1, j)
			left := j == 0 || v < hm.Get(i, j-1)
			right := j == hm.N()-1 || v < hm.Get(i, j+1)
			if up && down && left && right {
				coords = append(coords, sh.Coordinates{I: i, J: j})
			}
		}
	}
	return
}

type basin []sh.Coordinates

func getBasins(hm sh.Matrix, lp []sh.Coordinates) (basins []basin) {
	for _, p := range lp {
		basins = append(basins, getBasin(hm, p))
	}
	return
}

func getBasin(hm sh.Matrix, lp sh.Coordinates) (b basin) {
	queue := []sh.Coordinates{lp}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		if hm[c.I][c.J] == 9 || b.contains(c) {
			continue
		}
		b = append(b, c)

		// Neighbors
		possibles := []sh.Coordinates{
			{I: c.I - 1, J: c.J},
			{I: c.I + 1, J: c.J},
			{I: c.I, J: c.J - 1},
			{I: c.I, J: c.J + 1},
		}
		for _, p := range possibles {
			if p.I >= 0 && p.I < hm.M() && p.J >= 0 && p.J < hm.N() {
				queue = append(queue, p)
			}
		}
	}
	return
}

func (b basin) contains(p sh.Coordinates) bool {
	for _, v := range b {
		if v.I == p.I && v.J == p.J {
			return true
		}
	}
	return false
}

func riskLevel(hm sh.Matrix, coords []sh.Coordinates) (risk int) {
	for _, c := range coords {
		risk += hm.Get(c.I, c.J) + 1
	}
	return
}

func init() {
	lines := sh.Lines(input)
	heightmap = sh.NewMatrix(len(lines), len(lines[0]))
	for y, line := range lines {
		heightmap[y] = sh.ToInts(strings.Split(line, ""))
	}
}
