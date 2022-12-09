package main

import (
	_ "embed"
	"fmt"
	"regexp"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var moves [][3]int

type crates [9][]rune

func newCrates() crates {
	return [9][]rune{
		{'R', 'N', 'P', 'G'},
		{'T', 'J', 'B', 'L', 'C', 'S', 'V', 'H'},
		{'T', 'D', 'B', 'M', 'N', 'L'},
		{'R', 'V', 'P', 'S', 'B'},
		{'G', 'C', 'Q', 'S', 'W', 'M', 'V', 'H'},
		{'W', 'Q', 'S', 'C', 'D', 'B', 'J'},
		{'F', 'Q', 'L'},
		{'W', 'M', 'H', 'T', 'D', 'L', 'F', 'V'},
		{'L', 'P', 'B', 'V', 'M', 'J', 'F'},
	}
}

func (c *crates) popN(from, n int) []rune {
	var runes []rune
	for i := 0; i < n; i++ {
		r := c[from-1][len(c[from-1])-1]
		runes = append([]rune{r}, runes...)
		c[from-1] = c[from-1][:len(c[from-1])-1]
	}
	return runes
}

func (c *crates) pushN(to int, runes []rune) {
	c[to-1] = append(c[to-1], runes...)
}

func (c *crates) top() string {
	result := ""
	for _, stack := range c {
		result += string(stack[len(stack)-1])
	}
	return result
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() string {
	c := newCrates()
	for _, move := range moves {
		for i := 0; i < move[0]; i++ {
			r := c.popN(move[1], 1)
			c.pushN(move[2], r)
		}
	}
	return c.top()
}

func part2() string {
	c := newCrates()
	for _, move := range moves {
		r := c.popN(move[1], move[0])
		c.pushN(move[2], r)
	}
	return c.top()
}

func init() {
	regex := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	for _, line := range sh.Lines(input) {
		if regex.MatchString(line) {
			fields := regex.FindStringSubmatch(line)
			moves = append(moves, [3]int{
				sh.ToInt(fields[1]),
				sh.ToInt(fields[2]),
				sh.ToInt(fields[3]),
			})
		}
	}
}
