package main

import (
	_ "embed"
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var inputs []int
var boards []*board

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func init() {
	lines := sh.Lines(input)
	inputs = sh.IntList(lines[0])
	boards = parseBoardsLines(lines[1:])
}

func part1() int {
	for _, input := range inputs {
		for _, b := range boards {
			b.mark(input)
			if b.hasWon() {
				return b.unmarkedSum() * input
			}
		}
	}
	return 0
}

func part2() int {
	for _, input := range inputs {
		var next []*board

		for _, b := range boards {
			b.mark(input)
			if !b.hasWon() {
				next = append(next, b)
			}
		}

		if len(boards) == 1 && boards[0].hasWon() {
			return boards[0].unmarkedSum() * input
		}

		boards = next
	}
	return 0
}

// Board and methods

type board struct {
	rows []int
	mask []bool
}

func (b *board) mark(number int) {
	for i, n := range b.rows {
		if n == number {
			b.mask[i] = true
		}
	}
}

func (b *board) hasWon() bool {
	for i := 0; i < 5; i++ {
		// Rows
		if b.mask[5*i] && b.mask[5*i+1] && b.mask[5*i+2] && b.mask[5*i+3] && b.mask[5*i+4] {
			return true
		}
		// Columns
		if b.mask[i] && b.mask[i+5] && b.mask[i+10] && b.mask[i+15] && b.mask[i+20] {
			return true
		}
	}
	return false
}

func (b *board) unmarkedSum() (sum int) {
	for i, marked := range b.mask {
		if !marked {
			sum += b.rows[i]
		}
	}
	return
}

// Parsing

func parseBoardLine(line string) []int {
	line = strings.Trim(line, " ")
	line = strings.ReplaceAll(line, "  ", " ")
	row := strings.Split(line, " ")
	return sh.ToInts(row)
}

func parseBoardsLines(lines []string) (boards []*board) {
	var buffer []int
	for _, line := range lines {
		if line == "" {
			continue
		}

		numbers := parseBoardLine(line)
		buffer = append(buffer, numbers...)
		if len(buffer) == 25 {
			boards = append(boards, &board{rows: buffer, mask: make([]bool, 25)})
			buffer = nil
		}
	}
	return
}
