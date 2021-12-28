package main

import (
	_ "embed"
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var pop population

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	return pop.fishes(80)
}

func part2() int {
	return pop.fishes(256)
}

type population map[int]int

func (p population) fishes(days int) int {
	pop := dup(p)
	for i := 0; i < days; i++ {
		pop.shift()
	}
	return pop.count()
}

func (p population) shift() {
	new := p[0]
	for k := 0; k < 8; k++ {
		p[k] = p[k+1]
	}
	p[6] += new
	p[8] = new
}

func (p population) count() (c int) {
	for _, v := range p {
		c += v
	}
	return
}

func dup(p population) population {
	d := make(population)
	for k, v := range p {
		d[k] = v
	}
	return d
}

func init() {
	pop = make(population)
	lines := sh.Lines(input)
	for _, i := range sh.IntList(lines[0]) {
		pop[i]++
	}
}
