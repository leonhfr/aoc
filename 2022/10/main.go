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
	INDEX_START = 20
	INDEX_STEP  = 40
	INDEX_END   = 220
)

var lines []string

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: \n%v\n", part2())
}

func part1() int {
	return sumRegister(computeRegister())
}

func part2() string {
	return printScreen(computeScreen(computeRegister()))
}

func computeRegister() []int {
	register := []int{1}
	for _, line := range lines {
		switch x := register[len(register)-1]; {
		case line == "noop":
			register = append(register, x)
		case strings.HasPrefix(line, "addx "):
			op := sh.ToInt(strings.TrimPrefix(line, "addx "))
			register = append(register, x, x+op)
		}
	}
	return register
}

func sumRegister(register []int) int {
	var sum int
	for i := INDEX_START; i <= INDEX_END; i += INDEX_STEP {
		sum += i * register[i-1]
	}
	return sum
}

func computeScreen(register []int) [6][40]bool {
	fmt.Println(register)
	var screen [6][40]bool
	for row := 0; row < 6; row++ {
		for pixel := 0; pixel < 40; pixel++ {
			index := 40*row + pixel
			sprite := register[index]
			if sprite-1 <= pixel && pixel <= sprite+1 {
				screen[row][pixel] = true
			}
		}
	}
	return screen
}

func printScreen(screen [6][40]bool) string {
	lines := []string{"", "", "", "", "", ""}
	for i, row := range screen {
		for _, pixel := range row {
			if pixel {
				lines[i] += "#"
			} else {
				lines[i] += " "
			}
		}
	}
	return strings.Join(lines, "\n")
}

func init() {
	lines = sh.Lines(input)
}
