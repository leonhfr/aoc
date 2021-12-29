package main

import (
	_ "embed"
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

const (
	r_empty  = '.'
	r_right  = '>'
	r_bottom = 'v'
)

func main() {
	fmt.Printf("Part 1: %v\n", part1())
}

func part1() int {
	m := new(input)
	for i := 1; ; i++ {
		c := 0
		c += m.next(move{r_right, r_empty, vector{1, 0}})
		c += m.next(move{r_bottom, r_empty, vector{0, 1}})
		if c == 0 {
			return i
		}
	}
}

type matrix [][]rune

type vector struct {
	i, j int
}

type move struct {
	m, e rune
	v    vector
}

type transaction struct {
	from vector
	to   vector
}

func new(input string) matrix {
	lines := sh.Lines(input)
	m := make(matrix, len(lines))
	for y, line := range lines {
		row := make([]rune, len(line))
		for x, r := range line {
			row[x] = r
		}
		m[y] = row
	}
	return m
}

func (m *matrix) next(mv move) int {
	var txs []transaction
	matrix := *m
	for y, row := range matrix {
		for x, v := range row {
			next := vector{x + mv.v.i, y + mv.v.j}
			if next.i >= len(row) {
				next.i = 0
			}
			if next.j >= len(matrix) {
				next.j = 0
			}
			if v == mv.m && matrix[next.j][next.i] == mv.e {
				txs = append(txs, transaction{vector{x, y}, next})
			}
		}
	}
	for _, tx := range txs {
		matrix[tx.from.j][tx.from.i], matrix[tx.to.j][tx.to.i] = matrix[tx.to.j][tx.to.i], matrix[tx.from.j][tx.from.i]
	}
	return len(txs)
}

func (m *matrix) String() string {
	rows := make([]string, len(*m))
	for i, row := range *m {
		rows[i] = string(row)
	}
	return strings.Join(rows, "\n")
}
