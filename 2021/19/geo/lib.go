package geo

import (
	"fmt"

	sp "github.com/leonhfr/aoc/shared/setpoint3d"
)

type RebasedPoints map[sp.Point]*sp.SetPoint

type Orientation struct {
	i, j, k int
	a, b, c rune
}

func Rebase(pts1, pts2 *sp.SetPoint) RebasedPoints {
	rebased := make(RebasedPoints)
	for _, p1 := range pts1.Points() {
		s := sp.New()
		for _, p2 := range pts2.Points() {
			s.Add(sp.SubtractPts(p1, p2))
		}
		rebased[p1] = s
	}
	return rebased
}

func Rotations(points *sp.SetPoint) map[Orientation]*sp.SetPoint {
	rotations := make(map[Orientation]*sp.SetPoint)
	for _, o := range orientations {
		rotations[o] = rotateAll(points, o)
	}
	return rotations
}

func Rotate(p sp.Point, o Orientation) sp.Point {
	return sp.Point{
		X: o.i * element(p, o.a),
		Y: o.j * element(p, o.b),
		Z: o.k * element(p, o.c),
	}
}

func rotateAll(points *sp.SetPoint, o Orientation) *sp.SetPoint {
	s := sp.New()
	for _, p := range points.Points() {
		s.Add(Rotate(p, o))
	}
	return s
}

func element(p sp.Point, r rune) int {
	switch r {
	case 'x':
		return p.X
	case 'y':
		return p.Y
	case 'z':
		return p.Z
	default:
		panic(fmt.Sprintf("expected r to be one of x,y,z, got %v", r))
	}
}

var orientations = []Orientation{
	{1, 1, 1, 'x', 'y', 'z'},
	{1, 1, 1, 'y', 'z', 'x'},
	{1, 1, 1, 'z', 'x', 'y'},
	{1, 1, -1, 'z', 'y', 'x'},
	{1, 1, -1, 'y', 'x', 'z'},
	{1, 1, -1, 'x', 'z', 'y'},
	{1, -1, -1, 'x', 'y', 'z'},
	{1, -1, -1, 'y', 'z', 'x'},
	{1, -1, -1, 'z', 'x', 'y'},
	{1, -1, 1, 'z', 'y', 'x'},
	{1, -1, 1, 'y', 'x', 'z'},
	{1, -1, 1, 'x', 'z', 'y'},
	{-1, 1, -1, 'x', 'y', 'z'},
	{-1, 1, -1, 'y', 'z', 'x'},
	{-1, 1, -1, 'z', 'x', 'y'},
	{-1, 1, 1, 'z', 'y', 'x'},
	{-1, 1, 1, 'y', 'x', 'z'},
	{-1, 1, 1, 'x', 'z', 'y'},
	{-1, -1, 1, 'x', 'y', 'z'},
	{-1, -1, 1, 'y', 'z', 'x'},
	{-1, -1, 1, 'z', 'x', 'y'},
	{-1, -1, -1, 'z', 'y', 'x'},
	{-1, -1, -1, 'y', 'x', 'z'},
	{-1, -1, -1, 'x', 'z', 'y'},
}
