package main

import (
	_ "embed"
	"fmt"
	"regexp"

	sh "github.com/leonhfr/aoc/shared"
	sp "github.com/leonhfr/aoc/shared/setpoint"
)

//go:embed input
var input string

var matrix *sp.SetPoint
var instructions []instruction

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: \n%v\n", part2())
}

func part1() int {
	fold(matrix, instructions[0])
	return matrix.Len()
}

func part2() string {
	for _, i := range instructions {
		fold(matrix, i)
	}
	return matrix.String()
}

type instruction struct {
	x bool
	n int
}

func fold(matrix *sp.SetPoint, i instruction) {
	for _, p := range matrix.Points() {
		if i.x && p.X > i.n {
			// fold left
			matrix.Add(sp.Point{X: 2*i.n - p.X, Y: p.Y})
			matrix.Del(p)
		}

		if !i.x && p.Y > i.n {
			// fold up
			matrix.Add(sp.Point{X: p.X, Y: 2*i.n - p.Y})
			matrix.Del(p)
		}
	}
}

func init() {
	pointRegex := regexp.MustCompile(`^(\d+),(\d+)$`)
	instructionRegex := regexp.MustCompile(`^fold along (x|y)=(\d+)$`)

	matrix = sp.New()

	for _, line := range sh.Lines(input) {
		if pointRegex.MatchString(line) {
			fields := pointRegex.FindStringSubmatch(line)
			x, y := sh.ToInt(fields[1]), sh.ToInt(fields[2])
			matrix.Add(sp.Point{X: x, Y: y})
		}

		if instructionRegex.MatchString(line) {
			fields := instructionRegex.FindStringSubmatch(line)
			x := fields[1] == "x"
			n := sh.ToInt(fields[2])
			instructions = append(instructions, instruction{x, n})
		}
	}
}
