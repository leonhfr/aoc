package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"github.com/leonhfr/aoc/2021/19/geo"
	sh "github.com/leonhfr/aoc/shared"
	sp "github.com/leonhfr/aoc/shared/setpoint3d"
)

//go:embed input
var input string

var perspectives []perspective

type perspective struct {
	id  int
	set *sp.SetPoint
}

func main() {
	beacons, scanners := part1()
	fmt.Printf("Part 1: %v\n", beacons)
	fmt.Printf("Part 2: %v\n", part2(scanners))
}

func part1() (int, *sp.SetPoint) {
	beacons, scanners := locations()
	return beacons.Len(), scanners
}

func part2(scanners *sp.SetPoint) int {
	var distances []int
	for _, p1 := range scanners.Points() {
		for _, p2 := range scanners.Points() {
			distances = append(distances, sp.ManhattanDistance(p1, p2))
		}
	}
	sort.Ints(distances)
	return distances[len(distances)-1]
}

func locations() (beacons, scanners *sp.SetPoint) {
	beacons, scanners = perspectives[0].set, sp.New()
	scanners.Add(sp.Point{})
	pers, queue := perspectives[0], perspectives[1:]
	for len(queue) > 0 {
		pers, queue = queue[0], queue[1:]

		sol := solve(pers, beacons)
		if sol == nil {
			queue = append(queue, pers)
			continue
		}

		scanner := sp.SubtractPts(sol.p0, sol.p)
		scanners.Add(scanner)
		for _, p := range pers.set.Points() {
			rebasedP := sp.AddPts(geo.Rotate(p, sol.o), scanner)
			beacons.Add(rebasedP)
		}
	}
	return beacons, scanners
}

type solution struct {
	p0, p sp.Point
	o     geo.Orientation
}

func solve(pers perspective, beacons *sp.SetPoint) *solution {
	rebased0 := geo.Rebase(beacons, beacons)
	for orientation, oriented := range geo.Rotations(pers.set) {
		rebased := geo.Rebase(oriented, oriented)
		for p0, set0 := range rebased0 {
			for p, set := range rebased {
				intersection := set0.Intersect(set)
				if intersection.Len() > 11 {
					return &solution{p0, p, orientation}
				}
			}
		}
	}
	return nil
}

func init() {
	chunks := strings.Split(input, "\n\n")
	for id, chunk := range chunks {
		pers := perspective{id, sp.New()}
		for _, line := range strings.Split(chunk, "\n") {
			if strings.Contains(line, "---") || len(line) == 0 {
				continue
			}

			c := sh.ToInts(strings.Split(line, ","))
			pers.set.Add(sp.Point{X: c[0], Y: c[1], Z: c[2]})
		}
		perspectives = append(perspectives, pers)
	}
}
