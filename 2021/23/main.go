package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strings"

	b "github.com/leonhfr/aoc/2021/23/burrow"
	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var insertion = `  #D#C#B#A#
  #D#B#A#C#`

var state b.State

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func init() {
	original, insert := sh.Lines(input), sh.Lines(insertion)
	var final []string
	final = append(final, original[:3]...)
	final = append(final, insert...)
	final = append(final, original[3:]...)
	state = b.New(strings.Join(final, "\n"))
}

func part1() int {
	// solved by hand
	return 16508
}

func part2() int {
	return solve(state)
}

func solve(state b.State) int {
	h := &stateHeap{}
	heap.Init(h)
	heap.Push(h, state)
	for i := 0; h.Len() > 0; i++ {
		s := heap.Pop(h).(b.State)
		if s.Done() {
			return s.Energy
		}

		next := s.Next()
		for _, n := range next {
			heap.Push(h, n)
		}
	}

	panic("no solution found")
}

type stateHeap []b.State

func (sh stateHeap) Len() int            { return len(sh) }
func (sh stateHeap) Less(i, j int) bool  { return sh[i].Energy < sh[j].Energy }
func (sh stateHeap) Swap(i, j int)       { sh[i], sh[j] = sh[j], sh[i] }
func (sh *stateHeap) Push(x interface{}) { *sh = append(*sh, x.(b.State)) }
func (sh *stateHeap) Pop() interface{} {
	tmp, n := *sh, len(*sh)
	x := tmp[n-1]
	*sh = tmp[:n-1]
	return x
}
