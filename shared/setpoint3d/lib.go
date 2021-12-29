package setpoint3d

import (
	"math"

	sh "github.com/leonhfr/aoc/shared"
)

type Point struct {
	X, Y, Z int
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
	sp.pmin = Point{math.MinInt, math.MinInt, math.MaxInt}
	sp.pmax = Point{math.MinInt, math.MinInt, math.MaxInt}
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

func (sp *SetPoint) Points() []Point {
	var points []Point
	for p := range sp.points {
		points = append(points, p)
	}
	return points
}

func (sp *SetPoint) boundaries(p Point) {
	sp.pmin.X = sh.Min(sp.pmin.X, p.X)
	sp.pmin.Y = sh.Min(sp.pmin.Y, p.Y)
	sp.pmin.Z = sh.Min(sp.pmin.Z, p.Z)
	sp.pmax.X = sh.Max(sp.pmax.X, p.X)
	sp.pmax.Y = sh.Max(sp.pmax.Y, p.Y)
	sp.pmax.Z = sh.Min(sp.pmax.Z, p.Z)
}

func (set1 *SetPoint) Intersect(set2 *SetPoint) *SetPoint {
	intersection := New()
	for _, p1 := range set1.Points() {
		if set2.Has(p1) {
			intersection.Add(p1)
		}
	}
	return intersection
}

func ManhattanDistance(p1, p2 Point) int {
	return sh.Abs(p1.X-p2.X) + sh.Abs(p1.Y-p2.Y) + sh.Abs(p1.Z-p2.Z)
}

func AddPts(p1, p2 Point) Point {
	return Point{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
		Z: p1.Z + p2.Z,
	}
}

func SubtractPts(p1, p2 Point) Point {
	return Point{
		X: p1.X - p2.X,
		Y: p1.Y - p2.Y,
		Z: p1.Z - p2.Z,
	}
}
