package main

import (
	_ "embed"
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var orbits omap

const (
	center = "COM"
	santa  = "SAN"
	you    = "YOU"
)

func main() {
	fmt.Printf("Part 1: %v\n", part1(orbits))
	fmt.Printf("Part 2: %v\n", part2(orbits))
}

func init() {
	lines := sh.Lines(input)
	orbits = new(lines)
}

func part1(o omap) int {
	return o.count()
}

func part2(o omap) int {
	l, r := o.shortest(santa, you)
	return len(l) + len(r)
}

type omap map[string]string

func new(lines []string) omap {
	o := make(map[string]string)
	for _, line := range lines {
		s := strings.Split(line, ")")
		o[s[1]] = s[0]
	}
	return o
}

func (o omap) count() int {
	c := len(o)
	for _, v := range o {
		for p := o[v]; v != center; p = o[v] {
			v, c = p, c+1
		}
	}
	return c
}

func (o omap) shortest(a, b string) ([]string, []string) {
	pa, pb := o.path(a), o.path(b)
	for i := 0; i < len(pa); i++ {
		for j := 0; j < len(pb); j++ {
			if pa[i] == pb[j] {
				return pa[:i], pb[:j]
			}
		}
	}
	return nil, nil
}

func (o omap) path(k string) []string {
	var paths []string
	for p, ok := o[k]; ok; p, ok = o[k] {
		k, paths = p, append(paths, p)
	}
	return paths
}
