package main

import (
	_ "embed"
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var g graph

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	p := explore(p1)
	return len(p)
}

func part2() int {
	p := explore(p2)
	return len(p)
}

type graph map[string][]string

type paths [][]string

func explore(authorized func([]string, string) bool) (p paths) {
	c, queue := []string{}, [][]string{{"start"}}
	for len(queue) > 0 {
		c, queue = queue[0], queue[1:]
		for _, n := range g[c[len(c)-1]] {
			if n == "end" {
				p = append(p, append(c, n))
				continue
			}

			big := strings.ToUpper(n) == n
			if big || authorized(c, n) {
				tmp := append([]string(nil), c...)
				queue = append(queue, append(tmp, n))
			}
		}
	}
	return
}

func p1(path []string, node string) bool {
	for _, n := range path {
		if n == node {
			return false
		}
	}
	return true
}

func p2(path []string, node string) bool {
	visited := map[string]int{node: 1}
	twice := false
	for _, n := range path {
		if strings.ToUpper(n) == n {
			continue
		}

		visited[n]++

		if visited[n] == 2 && twice {
			return false
		}

		if visited[n] == 2 {
			twice = true
		}

		if visited[n] > 2 {
			return false
		}
	}
	return true
}

func init() {
	g = make(graph)
	for _, line := range sh.Lines(input) {
		nodes := strings.Split(line, "-")
		if nodes[1] != "start" {
			g[nodes[0]] = append(g[nodes[0]], nodes[1])
		}
		if nodes[0] != "start" {
			g[nodes[1]] = append(g[nodes[1]], nodes[0])
		}
	}
	delete(g, "end")
}
