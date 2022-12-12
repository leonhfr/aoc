package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"math"

	mat "github.com/leonhfr/aoc/shared/matrix"
)

//go:embed input
var input string

var matrix mat.Matrix

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	return dijkstraMinCost(matrix)
}

func part2() int {
	extended := extendMatrix(matrix)
	return dijkstraMinCost(extended)
}

func dijkstraMinCost(matrix mat.Matrix) int {
	vectors := []vector{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	dm := mat.NewMatrixWithDefault(matrix.M(), matrix.N(), math.MaxInt)
	dm.Set(1, 1, 0)
	h := &riskHeap{}
	heap.Init(h)
	heap.Push(h, risk{mat.Coordinates{I: 1, J: 1}, 0})

	for h.Len() != 0 {
		r := heap.Pop(h).(risk)
		for _, v := range vectors {
			i := r.c.I + v.i
			j := r.c.J + v.j
			if !matrix.Inside(i, j) {
				continue
			}

			nr := dm.Get(r.c.I, r.c.J) + matrix.Get(i, j)
			if dm.Get(i, j) > nr {
				dm.Set(i, j, nr)
				heap.Push(h, risk{mat.Coordinates{I: i, J: j}, nr})
			}
		}
	}

	return dm.Get(dm.M(), dm.N())
}

func extendMatrix(matrix mat.Matrix) mat.Matrix {
	extended := mat.NewMatrix(5*matrix.M(), 5*matrix.N())
	for _, c := range matrix.Coordinates() {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				level := matrix.Get(c.I, c.J) + i + j
				if level > 9 {
					level = level % 9
				}
				extended.Set(i*matrix.M()+c.I, j*matrix.N()+c.J, level)
			}
		}
	}
	return extended
}

type vector struct {
	i, j int
}

type risk struct {
	c     mat.Coordinates
	level int
}

type riskHeap []risk

func (rh riskHeap) Len() int            { return len(rh) }
func (rh riskHeap) Swap(i, j int)       { rh[i], rh[j] = rh[j], rh[i] }
func (rh riskHeap) Less(i, j int) bool  { return rh[i].level < rh[j].level }
func (rh *riskHeap) Push(x interface{}) { *rh = append(*rh, x.(risk)) }
func (rh *riskHeap) Pop() interface{} {
	tmp, n := *rh, len(*rh)
	x := tmp[n-1]
	*rh = tmp[0 : n-1]
	return x
}

func init() {
	matrix = mat.IntMatrix(input)
}
