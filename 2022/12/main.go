package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"math"

	sh "github.com/leonhfr/aoc/shared"
	mat "github.com/leonhfr/aoc/shared/matrix"
)

//go:embed input
var input string

var (
	heightMap  mat.Matrix
	start, end mat.Coordinates
	vectors    = []vector{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
)

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	dm := dijkstraMin(heightMap)
	return dm.Get(start.I, start.J)
}

func part2() int {
	dm := dijkstraMin(heightMap)
	steps := math.MaxInt
	for _, coords := range heightMap.Coordinates() {
		if heightMap.Get(coords.I, coords.J) == 0 { // a
			if s := dm.Get(coords.I, coords.J); s < steps {
				steps = s
			}
		}
	}
	return steps
}

type (
	vector struct{ i, j int }
	level  struct {
		coords mat.Coordinates
		level  int
	}
	levelHeap []level
)

func (h levelHeap) Len() int           { return len(h) }
func (h levelHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h levelHeap) Less(i, j int) bool { return h[i].level > h[j].level }
func (h *levelHeap) Push(x any)        { *h = append(*h, x.(level)) }
func (h *levelHeap) Pop() any {
	tmp, n := *h, len(*h)
	x := tmp[n-1]
	*h = tmp[0 : n-1]
	return x
}

func dijkstraMin(matrix mat.Matrix) mat.Matrix {
	dm := mat.NewMatrixWithDefault(matrix.M(), matrix.N(), math.MaxInt)
	dm.Set(end.I, end.J, 0)
	h := &levelHeap{}
	heap.Init(h)
	heap.Push(h, level{mat.Coordinates{I: end.I, J: end.J}, 0})

	for h.Len() != 0 {
		l := heap.Pop(h).(level)
		for _, v := range vectors {
			i := l.coords.I + v.i
			j := l.coords.J + v.j
			if !matrix.Inside(i, j) {
				continue
			}

			// cannot climb more than 1
			if matrix.Get(l.coords.I, l.coords.J)-matrix.Get(i, j) > 1 {
				continue
			}

			next := dm.Get(l.coords.I, l.coords.J) + 1
			if dm.Get(i, j) > next {
				dm.Set(i, j, next)
				heap.Push(h, level{mat.Coordinates{I: i, J: j}, next})
			}
		}
	}

	return dm
}

func init() {
	lines := sh.Lines(input)
	heightMap = mat.NewMatrix(len(lines), len(lines[0]))

	for i, line := range lines {
		for j, r := range line {
			switch r {
			case 'S':
				heightMap[i][j] = int('a' - 'a')
				start.I = i + 1
				start.J = j + 1
			case 'E':
				heightMap[i][j] = int('z' - 'a')
				end.I = i + 1
				end.J = j + 1
			default:
				heightMap[i][j] = int(r - 'a')
			}
		}
	}
}
