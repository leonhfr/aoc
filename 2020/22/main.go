package main

import (
	_ "embed"
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var decks state

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	s := decks
	for !won(s) {
		s = round(s)
	}
	a := append(s.p1, s.p2...)
	return score(a)
}

func part2() int {
	return 0
}

type state struct {
	p1, p2 []int
}

func round(s state) state {
	c1, c2 := s.p1[0], s.p2[0]
	s.p1, s.p2 = s.p1[1:], s.p2[1:]
	if c1 > c2 {
		s.p1 = append(s.p1, c1)
		s.p1 = append(s.p1, c2)
	} else {
		s.p2 = append(s.p2, c2)
		s.p2 = append(s.p2, c1)
	}
	return s
}

func game(s state) bool {
	// true = p1 win
	mem := [][]int{s.p1, s.p2}
	for !won(s) {
		if memCont(mem[:len(mem)-2], s.p1) || memCont(mem[:len(mem)-2], s.p2) {
			return true
		}
		mem = append(mem, s.p1)
		mem = append(mem, s.p2)
	}
	return len(s.p2) == 0
}

func memCont(mem [][]int, d []int) bool {
	for _, m := range mem {
		if IntEql(m, d) {
			return true
		}
	}
	return false
}

func won(s state) bool {
	return len(s.p1) == 0 || len(s.p2) == 0
}

func score(a []int) (r int) {
	for i := 1; i <= len(a); i++ {
		r += i * a[len(a)-i]
	}
	return
}

func init() {
	blocks := strings.Split(input, "\n\n")
	decks.p1 = deck(blocks[0])
	decks.p2 = deck(blocks[1])
}

func deck(str string) (a []int) {
	lines := sh.Lines(str)
	for _, line := range lines[1:] {
		a = append(a, sh.ToInt(line))
	}
	return
}

func IntEql(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
