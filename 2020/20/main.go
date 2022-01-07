package main

import (
	_ "embed"
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
	sp "github.com/leonhfr/aoc/shared/setpoint"
)

//go:embed sample
var input string

var tiles map[int]tile
var image map[sp.Point]tile

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	rearrange()
	r := 1
	for id, t := range tiles {
		if len(t.neighbors) == 2 {
			r *= id
		}
	}
	fmt.Println(tiles)
	return r
}

func part2() int {
	return 0
}

type side int

const (
	top side = iota
	right
	bottom
	left
	rev_top
	rev_right
	rev_bottom
	rev_left
)

type tile struct {
	img       string
	edges     map[side]string
	neighbors map[int]side
}

func rearrange() {
	for id1, t1 := range tiles {
		for id2, t2 := range tiles {
			if id1 == id2 {
				continue
			}

			for s1, e1 := range t1.edges {
				for s2, e2 := range t2.edges {
					if e1 == e2 {
						tiles[id2].neighbors[id1] = s2
						tiles[id1].neighbors[id2] = s1
					}
				}
			}
		}
	}
}

func init() {
	tiles = make(map[int]tile)
	image = make(map[sp.Point]tile)
	blocks := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	for _, block := range blocks {
		lines := sh.Lines(block)
		id := sh.ToInt(strings.Trim(strings.Split(lines[0], " ")[1], ":"))
		tiles[id] = tile{
			strings.Join(lines[1:], "\n"),
			edges(lines[1:]),
			make(map[int]side),
		}
	}
}

func edges(lines []string) map[side]string {
	l, r := "", ""
	for _, line := range lines {
		l += line[:1]
		r += line[len(line)-1:]
	}
	return map[side]string{
		top:        lines[0],
		right:      r,
		bottom:     lines[len(lines)-1],
		left:       l,
		rev_top:    sh.ReverseStr(lines[0]),
		rev_right:  sh.ReverseStr(r),
		rev_bottom: sh.ReverseStr(lines[len(lines)-1]),
		rev_left:   sh.ReverseStr(l),
	}
}
