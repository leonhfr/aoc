package main

import (
	_ "embed"
	"fmt"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var numbers []int

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	return recitation(numbers, 2020)
}

func part2() int {
	return recitation(numbers, 30000000)
}

func recitation(starting []int, n int) (spoken int) {
	memory := make(map[int][]int)
	for t := 1; t <= n; t++ {
		if t <= len(numbers) {
			a := starting[t-1]
			spoken = a
			memory[a] = append(memory[a], t)
			continue
		}

		if mem := memory[spoken]; len(mem) == 1 {
			spoken = 0
			memory[0] = append(memory[0], t)
		} else {
			diff := mem[len(mem)-1] - mem[len(mem)-2]
			spoken = diff
			memory[diff] = append(memory[diff], t)
		}
	}
	return
}

func init() {
	numbers = sh.IntList(input)
}
