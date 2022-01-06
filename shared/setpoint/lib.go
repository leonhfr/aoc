package setpoint

import (
	"math"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

type Point struct {
	X, Y int
}

type SetPoint struct {
	points     map[Point]struct{}
	pmin, pmax Point
}

func New() *SetPoint {
	return &SetPoint{
		points: make(map[Point]struct{}),
	}
}

func (sp *SetPoint) Add(p Point) {
	sp.points[p] = struct{}{}
	sp.boundaries(p)
}

func (sp *SetPoint) Del(p Point) {
	delete(sp.points, p)
	sp.pmin, sp.pmax = Point{math.MaxInt, math.MaxInt}, Point{math.MinInt, math.MinInt}
	for p := range sp.points {
		sp.boundaries(p)
	}
}

func (sp *SetPoint) Has(p Point) bool {
	_, ok := sp.points[p]
	return ok
}

func (sp *SetPoint) Len() int {
	return len(sp.points)
}

func (sp *SetPoint) Boundaries() (Point, Point) {
	return sp.pmin, sp.pmax
}

func (sp *SetPoint) Inside(p Point) bool {
	return sp.pmin.X <= p.X && p.X <= sp.pmax.X && sp.pmin.Y <= p.Y && p.Y <= sp.pmax.Y
}

func (sp *SetPoint) Points() []Point {
	var points []Point
	for p := range sp.points {
		points = append(points, p)
	}
	return points
}

func ManhattanDistance(p1, p2 Point) int {
	return sh.Abs(p1.X-p2.X) + sh.Abs(p1.Y-p2.Y)
}

func (sp *SetPoint) boundaries(p Point) {
	sp.pmin.X = sh.Min(sp.pmin.X, p.X)
	sp.pmin.Y = sh.Min(sp.pmin.Y, p.Y)
	sp.pmax.X = sh.Max(sp.pmax.X, p.X)
	sp.pmax.Y = sh.Max(sp.pmax.Y, p.Y)
}

func (sp *SetPoint) String() string {
	m := sh.Abs(sp.pmax.Y-sp.pmin.Y) + 1
	n := sh.Abs(sp.pmax.X-sp.pmin.X) + 1
	matrix := make([]string, m)
	for y := sp.pmin.Y; y <= sp.pmax.Y; y++ {
		row := make([]string, n)
		for x := sp.pmin.X; x <= sp.pmax.X; x++ {
			if _, ok := sp.points[Point{x, y}]; ok {
				row[x-sp.pmin.X] = "#"
			} else {
				row[x-sp.pmin.X] = "."
			}
		}
		matrix[y-sp.pmin.Y] = strings.Join(row, "")
	}
	return strings.Join(matrix, "\n")
}

func ParseSharp(str string) *SetPoint {
	sp := New()
	for y, line := range sh.Lines(str) {
		for x, r := range line {
			if r == '#' {
				sp.Add(Point{x, y})
			}
		}
	}
	return sp
}
