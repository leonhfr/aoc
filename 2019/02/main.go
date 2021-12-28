package main

import (
	_ "embed"
	"fmt"

	ic "github.com/leonhfr/aoc/2019/intcode"
	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var intcode ic.Intcode

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func init() {
	intcode = sh.IntList(input)
}

func part1() int {
	code := ic.New(intcode, 12, 2)
	code.Process()
	return code[0]
}

func part2() int {
	target := 19690720
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			code := ic.New(intcode, noun, verb)
			code.Process()
			if code[0] == target {
				return 100*noun + verb
			}
		}
	}
	return 0
}
