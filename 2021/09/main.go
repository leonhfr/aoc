package main

import (
	_ "embed"
	"fmt"
	"sort"

	mat "github.com/leonhfr/aoc/shared/matrix"
)

//go:embed input
var input string

var heightmap mat.Matrix

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

func lowestPoints(hm mat.Matrix) (coords []mat.Coordinates) {
	for i := 1; i <= hm.M(); i++ {
		for j := 1; j <= hm.N(); j++ {
			v := hm.Get(i, j)
			up := i == 1 || v <= hm.Get(i-1, j)
			down := i == hm.M() || v < hm.Get(i+1, j)
			left := j == 1 || v < hm.Get(i, j-1)
			right := j == hm.N() || v < hm.Get(i, j+1)
			if up && down && left && right {
				coords = append(coords, mat.Coordinates{I: i, J: j})
			}
		}
	}
	return
}

type basin []mat.Coordinates

func getBasins(hm mat.Matrix, lp []mat.Coordinates) (basins []basin) {
	for _, p := range lp {
		basins = append(basins, getBasin(hm, p))
	}
	return
}

func getBasin(hm mat.Matrix, lp mat.Coordinates) (b basin) {
	queue := []mat.Coordinates{lp}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		if hm.Get(c.I, c.J) == 9 || b.contains(c) {
			continue
		}
		b = append(b, c)

		// Neighbors
		possibles := []mat.Coordinates{
			{I: c.I - 1, J: c.J},
			{I: c.I + 1, J: c.J},
			{I: c.I, J: c.J - 1},
			{I: c.I, J: c.J + 1},
		}
		for _, p := range possibles {
			if hm.Inside(p.I, p.J) {
				queue = append(queue, p)
			}
		}
	}
	return
}

func (b basin) contains(p mat.Coordinates) bool {
	for _, v := range b {
		if v.I == p.I && v.J == p.J {
			return true
		}
	}
	return false
}

func riskLevel(hm mat.Matrix, coords []mat.Coordinates) (risk int) {
	for _, c := range coords {
		risk += hm.Get(c.I, c.J) + 1
	}
	return
}

func init() {
	heightmap = mat.IntMatrix(input)
}
