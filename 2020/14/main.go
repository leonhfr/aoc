package main

import (
	_ "embed"
	"fmt"
	"regexp"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var instructions []instruction
var mem map[int]int

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	execute()
	return sum()
}

func part2() int {
	return 0
}

type instruction struct {
	mask         bool
	addrs, value int
	ones, zeroes int
}

func execute() {
	ones, zeroes := 0, 0
	for _, i := range instructions {
		if i.mask {
			ones, zeroes = i.ones, i.zeroes
		} else {
			v := i.value
			v = v | ones
			v = v & zeroes
			mem[i.addrs] = v
		}
	}
}

func sum() (s int) {
	for _, v := range mem {
		s += v
	}
	return
}

func init() {
	mem = make(map[int]int)
	reg1 := regexp.MustCompile(`^mask = (.+)$`)
	reg2 := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
	lines := sh.Lines(input)
	for _, line := range lines {
		if reg1.MatchString(line) {
			fields := reg1.FindStringSubmatch(line)
			ones, zeroes := 0, 0
			for _, r := range fields[1] {
				ones, zeroes = ones<<1, zeroes<<1
				switch r {
				case '1':
					zeroes = zeroes | 1
					ones = ones | 1
				case 'X':
					zeroes = zeroes | 1
				}
			}
			instructions = append(instructions, instruction{
				true, 0, 0, ones, zeroes,
			})
		}

		if reg2.MatchString(line) {
			fields := reg2.FindStringSubmatch(line)
			instructions = append(instructions, instruction{
				false, sh.ToInt(fields[1]), sh.ToInt(fields[2]), 0, 0,
			})
		}
	}
}
